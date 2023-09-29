// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package rule provides the CLI subcommand for managing rules

package engine

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/xeipuuv/gojsonschema"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/stacklok/mediator/internal/db"
	evalerrors "github.com/stacklok/mediator/internal/engine/errors"
	"github.com/stacklok/mediator/internal/engine/eval"
	"github.com/stacklok/mediator/internal/engine/ingester"
	engif "github.com/stacklok/mediator/internal/engine/interfaces"
	"github.com/stacklok/mediator/internal/engine/remediate"
	"github.com/stacklok/mediator/internal/providers"
	mediatorv1 "github.com/stacklok/mediator/pkg/api/protobuf/go/mediator/v1"
)

// RuleMeta is the metadata for a rule
// TODO: We probably should care about a version
type RuleMeta struct {
	// Name is the name of the rule
	Name string
	// Provider is the ID of the provider that this rule is for
	Provider string
	// Organization is the ID of the organization that this rule is for
	Organization *string
	// Group is the ID of the group that this rule is for
	Group *string
}

// String returns a string representation of the rule meta
func (r *RuleMeta) String() string {
	if r.Group != nil {
		return fmt.Sprintf("%s/group/%s/%s", r.Provider, *r.Group, r.Name)
	}
	return fmt.Sprintf("%s/org/%s/%s", r.Provider, *r.Organization, r.Name)
}

// RuleValidator validates a rule against a schema
type RuleValidator struct {
	// schema is the schema that this rule type must conform to
	schema *gojsonschema.Schema
	// paramSchema is the schema that the parameters for this rule type must conform to
	paramSchema *gojsonschema.Schema
}

// NewRuleValidator creates a new rule validator
func NewRuleValidator(rt *mediatorv1.RuleType) (*RuleValidator, error) {
	// Load schemas
	schemaLoader := gojsonschema.NewGoLoader(rt.Def.RuleSchema)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return nil, fmt.Errorf("cannot create json schema: %w", err)
	}

	var paramSchema *gojsonschema.Schema
	if rt.Def.ParamSchema != nil {
		paramSchemaLoader := gojsonschema.NewGoLoader(rt.Def.ParamSchema)
		paramSchema, err = gojsonschema.NewSchema(paramSchemaLoader)
		if err != nil {
			return nil, fmt.Errorf("cannot create json schema for params: %w", err)
		}
	}

	return &RuleValidator{
		schema:      schema,
		paramSchema: paramSchema,
	}, nil
}

// ValidateRuleDefAgainstSchema validates the given contextual policy against the
// schema for this rule type
func (r *RuleValidator) ValidateRuleDefAgainstSchema(contextualPolicy map[string]any) error {
	return validateAgainstSchema(r.schema, contextualPolicy)
}

// ValidateParamsAgainstSchema validates the given parameters against the
// schema for this rule type
func (r *RuleValidator) ValidateParamsAgainstSchema(params *structpb.Struct) error {
	if r.paramSchema == nil {
		return nil
	}

	if params == nil {
		return fmt.Errorf("params cannot be nil")
	}

	return validateAgainstSchema(r.paramSchema, params.AsMap())
}

func validateAgainstSchema(schema *gojsonschema.Schema, obj map[string]any) error {
	documentLoader := gojsonschema.NewGoLoader(obj)
	result, err := schema.Validate(documentLoader)
	if err != nil {
		return fmt.Errorf("cannot validate json schema: %w", err)
	}

	if !result.Valid() {
		return buildValidationError(result.Errors())
	}

	return nil
}

func buildValidationError(errs []gojsonschema.ResultError) error {
	problems := make([]string, 0, len(errs))
	for _, desc := range errs {
		problems = append(problems, desc.String())
	}

	return fmt.Errorf("invalid json schema: %s", strings.TrimSpace(strings.Join(problems, "\n")))
}

// RuleTypeEngine is the engine for a rule type
type RuleTypeEngine struct {
	Meta RuleMeta

	// rdi is the rule data ingest engine
	rdi engif.Ingester

	// reval is the rule evaluator
	reval engif.Evaluator

	// rrem is the rule remediator
	rrem engif.Remediator

	rval *RuleValidator

	rt *mediatorv1.RuleType

	cli *providers.ProviderBuilder
}

// NewRuleTypeEngine creates a new rule type engine
func NewRuleTypeEngine(rt *mediatorv1.RuleType, cli *providers.ProviderBuilder) (*RuleTypeEngine, error) {
	rval, err := NewRuleValidator(rt)
	if err != nil {
		return nil, fmt.Errorf("cannot create rule validator: %w", err)
	}

	rdi, err := ingester.NewRuleDataIngest(rt, cli)
	if err != nil {
		return nil, fmt.Errorf("cannot create rule data ingest: %w", err)
	}

	reval, err := eval.NewRuleEvaluator(rt, cli)
	if err != nil {
		return nil, fmt.Errorf("cannot create rule evaluator: %w", err)
	}

	rrem, err := remediate.NewRuleRemediator(rt, cli)
	if errors.Is(err, remediate.ErrNoRemediation) {
		// we should be graceful about not having a remediator
		// TODO: return a noop remediator instead that would log that there's nothing configured?
		rrem = nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot create rule remediator: %w", err)
	}

	rte := &RuleTypeEngine{
		Meta: RuleMeta{
			Name:     rt.Name,
			Provider: rt.Context.Provider,
		},
		rval:  rval,
		rdi:   rdi,
		reval: reval,
		rrem:  rrem,
		rt:    rt,
		cli:   cli,
	}

	// Set organization if it exists
	if rt.Context.Organization != nil && *rt.Context.Organization != "" {
		// We need to clone the string because the pointer is to a string literal
		// and we don't want to modify that
		org := strings.Clone(*rt.Context.Organization)
		rte.Meta.Organization = &org
	} else if rt.Context.Group != nil && *rt.Context.Group != "" {
		grp := strings.Clone(*rt.Context.Group)
		rte.Meta.Group = &grp
	} else {
		return nil, fmt.Errorf("rule type context must have an organization or group")
	}

	return rte, nil
}

// GetID returns the ID of the rule type. The ID is meant to be
// a serializable unique identifier for the rule type.
func (r *RuleTypeEngine) GetID() string {
	return r.Meta.String()
}

// GetRuleInstanceValidator returns the rule instance validator for this rule type.
// By instance we mean a rule that has been instantiated in a policy from a given rule type.
func (r *RuleTypeEngine) GetRuleInstanceValidator() *RuleValidator {
	return r.rval
}

// Eval runs the rule type engine against the given entity
func (r *RuleTypeEngine) Eval(
	ctx context.Context,
	ent protoreflect.ProtoMessage,
	pol, params map[string]any,
	remAction engif.RemediateActionOpt,
) error {
	result, err := r.rdi.Ingest(ctx, ent, params)
	if err != nil {
		return fmt.Errorf("error ingesting data: %w", err)
	}

	evalErr := r.reval.Eval(ctx, pol, result)

	remediateErr := r.tryRemediate(ctx, ent, pol, remAction, evalErr)
	if remediateErr != nil {
		// TODO(jakub): We should surface this error up and store in the rule evaluation status
		// for now we just log and do nothing else
		log.Printf("remediation error: %v", remediateErr)
	}

	return evalErr
}

func (r *RuleTypeEngine) tryRemediate(
	ctx context.Context,
	ent protoreflect.ProtoMessage,
	pol map[string]any,
	remAction engif.RemediateActionOpt,
	evalErr error,
) error {
	shouldRemediate, err := r.shouldRemediate(remAction, evalErr)
	if err != nil {
		return err
	} else if !shouldRemediate {
		return nil
	}

	return r.rrem.Remediate(ctx, remAction, ent, pol)
}

func (r *RuleTypeEngine) shouldRemediate(remAction engif.RemediateActionOpt, evalErr error) (bool, error) {
	if r.rrem == nil {
		return false, nil
	}

	var runRemediation bool

	switch remAction {
	case engif.ActionOptOff:
		return false, nil
	case engif.ActionOptUnknown:
		return false, errors.New("unknown remediation action, check your policy definition")
	case engif.ActionOptDryRun, engif.ActionOptOn:
		runRemediation = !errors.Is(evalErr, evalerrors.ErrEvaluationSkipped) ||
			errors.Is(evalErr, evalerrors.ErrEvaluationSkipSilently)
	}

	if evalErr == nil {
		runRemediation = false
	}

	return runRemediation, nil
}

// RuleDefFromDB converts a rule type definition from the database to a protobuf
// rule type definition
func RuleDefFromDB(r *db.RuleType) (*mediatorv1.RuleType_Definition, error) {
	def := &mediatorv1.RuleType_Definition{}

	if err := protojson.Unmarshal(r.Definition, def); err != nil {
		return nil, fmt.Errorf("cannot unmarshal rule type definition: %w", err)
	}
	return def, nil
}

// RuleTypePBFromDB converts a rule type from the database to a protobuf
// rule type
func RuleTypePBFromDB(rt *db.RuleType, ectx *EntityContext) (*mediatorv1.RuleType, error) {
	gname := ectx.GetGroup().GetName()

	def, err := RuleDefFromDB(rt)
	if err != nil {
		return nil, fmt.Errorf("cannot get rule type definition: %w", err)
	}

	id := rt.ID.String()

	return &mediatorv1.RuleType{
		Id:   &id,
		Name: rt.Name,
		Context: &mediatorv1.Context{
			Provider: ectx.GetProvider().Name,
			Group:    &gname,
		},
		Description: rt.Description,
		Guidance:    rt.Guidance,
		Def:         def,
	}, nil
}

// GetRulesFromPolicyOfType returns the rules from the policy of the given type
func GetRulesFromPolicyOfType(p *mediatorv1.Policy, rt *mediatorv1.RuleType) ([]*mediatorv1.Policy_Rule, error) {
	contextualRules, err := GetRulesForEntity(p, mediatorv1.EntityFromString(rt.Def.InEntity))
	if err != nil {
		return nil, fmt.Errorf("error getting rules for entity: %w", err)
	}

	rules := []*mediatorv1.Policy_Rule{}
	err = TraverseRules(contextualRules, func(r *mediatorv1.Policy_Rule) error {
		if r.Type == rt.Name {
			rules = append(rules, r)
		}
		return nil
	})

	// This shouldn't happen
	if err != nil {
		return nil, fmt.Errorf("error traversing rules: %w", err)
	}

	return rules, nil
}
