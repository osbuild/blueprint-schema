package blueprint

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/invopop/yaml"
	"github.com/osbuild/blueprint-schema"
)

type Schema struct {
	doc *openapi3.T
}

// ErrUnmarshal is returned when a buffer/reader cannot be unmarshaled.
var ErrUnmarshal = errors.New("cannot unmarshal JSON/YAML")

// ErrCannotCompileSchema is returned when the schema cannot be compiled.
var ErrCannotCompileSchema = errors.New("cannot compile schema")

// ErrValidateFailed is returned when the validation fails.
var ErrValidateFailed = errors.New("validation failed")

// CompileSourceSchema compiles the JSON schema. Uses the embedded schema
// from the oas/ directory. Returns the compiled schema or
// an error if the schema cannot be compiled.
//
// Do not use this schema for validation, use bundled schema instead.
func CompileSourceSchema() (*Schema, error) {
	return compileSchema(blueprint.SchemaSource())
}

// CompileBundledSchema compiles the JSON schema. Uses the bundled schema
// with extensions. Use this schema for validation.
func CompileBundledSchema() (*Schema, error) {
	return compileSchema(blueprint.BundledSchema())
}

func compileSchema(buf []byte) (*Schema, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return blueprint.SchemaFS.ReadFile(filepath.Join("oas", uri.Path))
	}

	location, _ := url.Parse(".")
	doc, err := loader.LoadFromDataWithPath(buf, location)
	if err != nil {
		panic(err)
	}

	return &Schema{doc: doc}, nil
}

func (s *Schema) Document() *openapi3.T {
	return s.doc
}

// Bundle resolves all references in the schema. It modifies the
// schema in place. It is not necessary to call this function
// if the schema is already bundled.
func (s *Schema) Bundle(ctx context.Context) error {
	s.doc.InternalizeRefs(ctx, func(s *openapi3.T, c openapi3.ComponentRef) string {
		return strings.TrimSuffix(c.RefString(), ".yaml")
	})

	delete(s.doc.Components.Schemas, "./components/blueprint")

	return nil
}

// ValidateSchema validates the schema. It checks if the schema is
// valid and if all references are valid.
func (s *Schema) ValidateSchema(ctx context.Context) error {
	if err := s.doc.Validate(ctx); err != nil {
		return fmt.Errorf("%w: %v", ErrValidateFailed, err)
	}

	return nil
}

// MarshalJSON marshals the schema to JSON.
func (s *Schema) MarshalJSON() ([]byte, error) {
	x, err := s.doc.MarshalYAML()
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(x, "", "  ")
}

// MarshalYAML marshals the schema to YAML.
func (s *Schema) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(s.doc)
}

// ValidateJSON reads JSON and performs validation.
func (s *Schema) ValidateJSON(ctx context.Context, data []byte) error {
	req, err := http.NewRequest(http.MethodGet, "/items", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	rvi := &openapi3filter.RequestValidationInput{
		Request: req,
	}
	rb := s.doc.Paths.Find("/validate_blueprint").GetOperation(http.MethodPost).RequestBody
	err = openapi3filter.ValidateRequestBody(context.Background(), rvi, rb.Value)
	if err != nil {
		return err
	}

	return nil
}

// ValidateYAML reads YAML and performs validation.
func (s *Schema) ValidateYAML(ctx context.Context, data []byte) error {
	json, err := ConvertYAMLtoJSON(data)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnmarshal, err)
	}

	return s.ValidateJSON(ctx, json)
}

// ReadAndValidateYAML reads YAML from the reader and calls ValidateMap.
// TODO: implement this function for JSON as well
func (s *Schema) ReadAndValidateYAML(ctx context.Context, reader io.Reader) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return err
	}

	return s.ValidateYAML(ctx, buf.Bytes())
}

func (s *Schema) ApplyExtensions(ctx context.Context) error {
	dir, err := blueprint.SchemaFS.ReadDir("oas/extensions/*.yaml")
	if err != nil {
		return err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		f, err := blueprint.SchemaFS.Open(filepath.Join("oas/extensions", file.Name()))
		if err != nil {
			return err
		}

		b, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			return err
		}

		j, err := ConvertYAMLtoJSON(b)
		if err != nil {
			return err
		}

		var ts openapi3.Schema
		ts.UnmarshalJSON(j)

		schemaName := filepath.Base(file.Name())
		if ts.AnyOf != nil {
			s.doc.Components.Schemas[schemaName].Value.AnyOf = ts.AnyOf
		}

		if ts.AllOf != nil {
			s.doc.Components.Schemas[schemaName].Value.AllOf = ts.AllOf
		}

		if ts.OneOf != nil {
			s.doc.Components.Schemas[schemaName].Value.OneOf = ts.OneOf
		}
	}
	return s.ValidateSchema(ctx)
}
