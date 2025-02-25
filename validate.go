package blueprint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/invopop/yaml"
	"github.com/kaptinlin/jsonschema"
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
	schema, err := compiler.Compile(SchemaJSON)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCannotCompileSchema, err)
	}

	return &Schema{s: schema}, nil
}

// ValidateMap validates the map against the schema. Returns true if the
// data is valid, otherwise false and a string with details. The result
// details can be randomly sorted.
func (s *Schema) ValidateMap(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	list := result.ToList(true)

	if !result.IsValid() {
		details, _ := json.MarshalIndent(list, "", "  ")
		return false, string(details)
	}

	return true, ""
}

// For testing purposes, this returns the details with removed errors and annotations.
// This is to make the output stable for comparison in test cases. For more info:
// https://github.com/kaptinlin/jsonschema/issues/28
func (s *Schema) ValidateMapStable(data map[string]any) (bool, string) {
	result := s.s.Validate(data)
	list := result.ToList(true)
	sortSchemaSlice(list.Details)

	if !result.IsValid() {
		// Must use encoding/json here: // https://github.com/kaptinlin/jsonschema/issues/27
		details, _ := json.MarshalIndent(list, "", "  ")
		return false, string(details)
	}

	return true, ""
}

var sortSchemaRE = regexp.MustCompile(`'[^']+'`)

func sortSchemaSlice(list []jsonschema.List) {
	if len(list) == 0 {
		return
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].EvaluationPath < list[j].EvaluationPath
	})

	for i := range list {
		for k, v := range list[i].Errors {
			if k != "properties" {
				continue
			}
			pOrig := sortSchemaRE.FindAllString(v, -1)
			if len(pOrig) <= 1 {
				continue
			}
			pNew := make([]string, len(pOrig))
			copy(pNew, pOrig)
			slices.Sort(pNew)
			list[i].Errors[k] = fmt.Sprintf("Properties %s have problems", strings.Join(pNew, ", "))
		}

		sortSchemaSlice(list[i].Details)
	}
}

// ValidateJSON unmarshal JSON and calls ValidateMap.
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

// ValidateYAML unmarshal YAML and calls ValidateMap.
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
