// SPDX-FileCopyrightText: Copyright 2024 The Minder Authors
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
	uritemplate "github.com/std-uritemplate/std-uritemplate/go/v2"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/mindersec/minder/internal/util"
	"github.com/mindersec/minder/internal/util/schemaupdate"
	"github.com/mindersec/minder/internal/util/schemavalidate"
	minderv1 "github.com/mindersec/minder/pkg/api/protobuf/go/minder/v1"
)

const (
	// MaxBytesLimit is the maximum number of bytes to read from the response body
	// We limit to 1MB to prevent abuse
	MaxBytesLimit int64 = 1 << 20
)

type restHandler struct {
	rawInputSchema *structpb.Struct
	inputSchema    *jsonschema.Schema
	endpointTmpl   string
	method         string
	body           string
	headers        map[string]string
	parse          string
	// TODO implement fallback
	// TODO implement auth
}

func newHandlerFromDef(def *minderv1.RestDataSource_Def) (*restHandler, error) {
	if def == nil {
		return nil, errors.New("rest data source handler definition is nil")
	}

	// schema may be nil
	schema, err := schemavalidate.CompileSchemaFromPB(def.GetInputSchema())
	if err != nil {
		return nil, err
	}

	return &restHandler{
		rawInputSchema: def.GetInputSchema(),
		inputSchema:    schema,
		endpointTmpl:   def.GetEndpoint(),
		method:         util.HttpMethodFromString(def.GetMethod(), http.MethodGet),
		headers:        def.GetHeaders(),
		body:           parseRequestBodyConfig(def),
		parse:          def.GetParse(),
	}, nil
}

func (h *restHandler) ValidateArgs(args any) error {
	if h.inputSchema == nil {
		return errors.New("input schema cannot be nil")
	}

	mapobj, ok := args.(map[string]any)
	if !ok {
		return errors.New("args is not a map")
	}

	return schemavalidate.ValidateAgainstSchema(h.inputSchema, mapobj)
}

func (h *restHandler) ValidateUpdate(obj any) error {
	if obj == nil {
		return errors.New("update schema cannot be nil")
	}

	switch castedobj := obj.(type) {
	case *structpb.Struct:
		if _, err := schemavalidate.CompileSchemaFromPB(castedobj); err != nil {
			return fmt.Errorf("update validation failed due to invalid schema: %w", err)
		}
		return schemaupdate.ValidateSchemaUpdate(h.rawInputSchema, castedobj)
	case map[string]any:
		if _, err := schemavalidate.CompileSchemaFromMap(castedobj); err != nil {
			return fmt.Errorf("update validation failed due to invalid schema: %w", err)
		}
		return schemaupdate.ValidateSchemaUpdateMap(h.rawInputSchema.AsMap(), castedobj)
	default:
		return errors.New("invalid type")
	}
}

func (h *restHandler) Call(_ context.Context, args any) (any, error) {
	argsMap, ok := args.(map[string]any)
	if !ok {
		return nil, errors.New("args is not a map")
	}

	expandedEndpoint, err := uritemplate.Expand(h.endpointTmpl, argsMap)
	if err != nil {
		return nil, err
	}

	// TODO: Add option to use custom client
	cli := http.Client{
		// TODO: Make timeout configurable
		Timeout: 5 * time.Second,
	}

	var b io.Reader
	if h.body != "" {
		b = strings.NewReader(h.body)
	}

	req, err := http.NewRequest(h.method, expandedEndpoint, b)
	if err != nil {
		return nil, err
	}

	for k, v := range h.headers {
		req.Header.Add(k, v)
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bout, err := h.parseResponseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	// TODO: Handle fallback here.

	return buildRestOutput(resp.StatusCode, bout), nil
}

func (h *restHandler) parseResponseBody(body io.Reader) (any, error) {
	var data any

	if body == nil {
		return nil, nil
	}

	lr := io.LimitReader(body, MaxBytesLimit)

	if h.parse == "json" {
		var jsonData any
		dec := json.NewDecoder(lr)
		if err := dec.Decode(&jsonData); err != nil {
			return nil, fmt.Errorf("cannot decode json: %w", err)
		}

		data = jsonData
	} else {
		bytedata, err := io.ReadAll(lr)
		if err != nil {
			return nil, fmt.Errorf("cannot read response body: %w", err)
		}

		data = string(bytedata)
	}

	return data, nil
}

// body may be unset, in which case it is nil
// or it may be an object or a string. We are using
// a oneof in the protobuf definition to represent this.
func parseRequestBodyConfig(def *minderv1.RestDataSource_Def) string {
	defBody := def.GetBody()
	if defBody == nil {
		return ""
	}

	if def.GetBodyobj() != nil {
		// stringify the object
		obj, err := json.Marshal(def.GetBodyobj())
		if err != nil {
			return ""
		}

		return string(obj)
	}

	return def.GetBodystr()
}

func buildRestOutput(statusCode int, body any) any {
	return map[string]any{
		"status_code": statusCode,
		"body":        body,
	}
}
