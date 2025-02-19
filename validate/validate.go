package validate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	gojson "github.com/goccy/go-json"
	"github.com/invopop/yaml"
	"github.com/kaptinlin/jsonschema"
	blueprint "github.com/lzap/common-blueprint-example"
)

type Schema struct {
	s *jsonschema.Schema
}

// ErrCannotCompileSchema is returned when the schema cannot be compiled.
var ErrCannotCompileSchema = errors.New("cannot compile schema")

// CompileSchema compiles the JSON schema. Uses the embedded schema
// available as blueprint.SchemaJSON.
func CompileSchema() (*Schema, error) {
	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile(blueprint.SchemaJSON)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCannotCompileSchema, err)
	}

	return &Schema{s: schema}, nil
}

// ValidateMap validates the map against the schema. Returns true if the
// data is valid, otherwise false and a string with details.
func (s *Schema) ValidateMap(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	if !result.IsValid() {
		details, _ := json.MarshalIndent(result.ToList(), "", "  ")
		return false, string(details)
	}

	return true, ""
}

// https://github.com/kaptinlin/jsonschema/issues/27
func (s *Schema) ValidateMap2(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	if !result.IsValid() {
		details, _ := gojson.MarshalIndent(result.ToList(), "", "  ")
		return false, string(details)
	}

	return true, ""
}

// ValidateJSON unmarshals JSON and calls ValidateMap.
func (s *Schema) ValidateJSON(data []byte) (bool, string, error) {
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}

// ReadAndValidateJSON reads JSON from the reader and calls ValidateMap.
func (s *Schema) ReadAndValidateJSON(reader io.Reader) (bool, string, error) {
	var m map[string]any
	if err := json.NewDecoder(reader).Decode(&m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}

// ValidateYAML unmarshals YAML and calls ValidateMap.
func (s *Schema) ValidateYAML(data []byte) (bool, string, error) {
	var m map[string]any
	if err := yaml.Unmarshal(data, &m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}

// ReadAndValidateYAML reads YAML from the reader and calls ValidateMap.
func (s *Schema) ReadAndValidateYAML(reader io.Reader) (bool, string, error) {
	var m map[string]any

	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return false, "", err
	}

	if err := yaml.Unmarshal(buf.Bytes(), &m); err != nil {
		return false, "", err
	}

	valid, details := s.ValidateMap(m)
	return valid, details, nil
}
