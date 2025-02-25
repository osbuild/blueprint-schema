package blueprint

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/invopop/yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

type Schema struct {
	s *jsonschema.Schema
}

// ErrUnmarshal is returned when a buffer/reader cannot be unmarshaled.
var ErrUnmarshal = errors.New("cannot unmarshal JSON/YAML")

// ErrCannotCompileSchema is returned when the schema cannot be compiled.
var ErrCannotCompileSchema = errors.New("cannot compile schema")

// ErrValidateFailed is returned when the validation fails.
var ErrValidateFailed = errors.New("validation failed")

// CompileSchema compiles the JSON schema. Uses the embedded schema
// available as blueprint.SchemaJSON. Returns the compiled schema or
// an error if the schema cannot be compiled.
func CompileSchema() (*Schema, error) {
	jsonSchema, err := jsonschema.UnmarshalJSON(bytes.NewBuffer(SchemaJSON))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.AddResource("blueprint-schema.json", jsonSchema)
	schema, err := compiler.Compile("blueprint-schema.json")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCannotCompileSchema, err)
	}

	// The validation library converts schema locations to absolute path which is not practical
	// for tests. Therefore, it is overridden here so errors do not contain pwd base path.
	schema.Location = "blueprint-schema.json"

	return &Schema{s: schema}, nil
}

func (s *Schema) ValidateMap(data any) error {
	if err := s.s.Validate(data); err != nil {
		return fmt.Errorf("%w: %v", ErrValidateFailed, err)
	}

	return nil
}

// ValidateJSON reads JSON and performs validation.
func (s *Schema) ValidateJSON(data []byte) error {
	jsonData, err := jsonschema.UnmarshalJSON(bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	return s.ValidateMap(jsonData)
}

// ReadAndValidateJSON reads JSON from the reader and performs validation.
func (s *Schema) ReadAndValidateJSON(reader io.Reader) error {
	jsonData, err := jsonschema.UnmarshalJSON(reader)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	return s.ValidateMap(jsonData)
}

// ValidateYAML reads YAML and performs validation.
func (s *Schema) ValidateYAML(data []byte) error {
	var m map[string]any
	if err := yaml.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	return s.ValidateMap(m)
}

// ReadAndValidateYAML reads YAML from the reader and calls ValidateMap.
func (s *Schema) ReadAndValidateYAML(reader io.Reader) error {
	var m map[string]any

	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(buf.Bytes(), &m); err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	return s.ValidateMap(m)
}
