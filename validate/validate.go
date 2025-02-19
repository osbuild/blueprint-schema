package validate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	gojson "github.com/goccy/go-json"
	"github.com/kaptinlin/jsonschema"
	blueprint "github.com/lzap/common-blueprint-example"
)

type Schema struct {
	s *jsonschema.Schema
}

var ErrCannotCompileSchema = errors.New("cannot compile schema")

func CompileSchema() (*Schema, error) {
	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(blueprint.SchemaJSON)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCannotCompileSchema, err)
	}

	return &Schema{s: schema}, nil
}

func (s *Schema) ValidateMap(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	if !result.IsValid() {
		details, _ := json.MarshalIndent(result.ToList(), "", "  ")
		return false, string(details)
	}

	return true, ""
}

func (s *Schema) ValidateMap2(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	if !result.IsValid() {
		details, _ := gojson.MarshalIndent(result.ToList(), "", "  ")
		return false, string(details)
	}

	return true, ""
}

func (s *Schema) ValidateJSON(data []byte) (bool, string, error) {
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}

func (s *Schema) ReadAndValidateJSON(reader io.Reader) (bool, string, error) {
	var m map[string]any
	if err := json.NewDecoder(reader).Decode(&m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}
