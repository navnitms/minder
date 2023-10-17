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

// Package noop provides a fallback remediation engine for cases where
// no remediation is set.
package noop

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	enginerr "github.com/stacklok/mediator/internal/engine/errors"
	"github.com/stacklok/mediator/internal/engine/interfaces"
	pb "github.com/stacklok/mediator/pkg/api/protobuf/go/mediator/v1"
)

// Remediator is the structure backing the noop remediator
type Remediator struct {
	actionType interfaces.ActionType
}

// NewNoopRemediate creates a new noop remediation engine
func NewNoopRemediate(actionType interfaces.ActionType) (*Remediator, error) {
	return &Remediator{actionType: actionType}, nil
}

// ParentType returns the action type of the noop engine
func (r *Remediator) ParentType() interfaces.ActionType {
	return r.actionType
}

// SubType returns the action subtype of the remediation engine
func (_ *Remediator) SubType() string {
	return "noop"
}

// GetOnOffState returns the off state of the noop engine
func (_ *Remediator) GetOnOffState(_ *pb.Profile) interfaces.ActionOpt {
	return interfaces.ActionOptOff
}

// Do perform the remediation
func (r *Remediator) Do(
	_ context.Context,
	_ interfaces.ActionCmd,
	_ interfaces.ActionOpt,
	_ protoreflect.ProtoMessage,
	_ *interfaces.EvalStatusParams,
	_ *json.RawMessage,
) (json.RawMessage, error) {
	return nil, fmt.Errorf("%s:%w", r.ParentType(), enginerr.ErrActionNotAvailable)
}
