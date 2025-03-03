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
	"text/scanner"

	"github.com/invopop/yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"golang.org/x/text/message"
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

	return &Schema{s: schema}, nil
}

// ValidateMap validates the given map. Returns an error if the validation fails.
func (s *Schema) ValidateMap(data any) error {
	if err := s.s.Validate(data); err != nil {
		return fmt.Errorf("%w: %v", ErrValidateFailed, err)
	}

	return nil
}

// Validates blueprint and returns an error if the validation fails.
func (s *Schema) Validate(bp *Blueprint) error {
	buf, err := json.Marshal(bp)
	if err != nil {
		return fmt.Errorf("cannot marshal blueprint: %w", err)
	}

	var bpMap map[string]any
	if err := json.Unmarshal(buf, &bpMap); err != nil {
		return fmt.Errorf("cannot unmarshal blueprint: %w", err)
	}

	return s.ValidateMap(bpMap)
}

var sortSchemaRE = regexp.MustCompile(`'[^']+'`)

func sortQuotedSubstrings(str string) string {
	subs := sortSchemaRE.FindAllString(str, -1)
	if len(subs) <= 1 {
		return str
	}
	slices.Sort(subs)

	var s scanner.Scanner
	var sb strings.Builder
	var quoted bool
	s.Init(strings.NewReader(str))
	s.Whitespace = 0
	for _, c := range str {
		if c == '\'' {
			quoted = !quoted

			if quoted {
				sb.WriteString(subs[0])
				subs = subs[1:]
			}
			continue
		}
		if !quoted {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}

type wrappedErrorKind struct {
	jsonschema.ErrorKind
}

var _ jsonschema.ErrorKind = (*wrappedErrorKind)(nil)

func (w *wrappedErrorKind) KeywordPath() []string {
	// After upgrade to Go 1.23 can be replaced just with ... range slices.Sorted(arr)
	sarr := make([]string, len(w.ErrorKind.KeywordPath()))
	copy(sarr, w.ErrorKind.KeywordPath())
	sort.Strings(sarr)
	return sarr
}

func (w *wrappedErrorKind) LocalizedString(p *message.Printer) string {
	//println("LocalizedString", p.Sprintf(w.ErrorKind.LocalizedString(p)))
	return sortQuotedSubstrings(p.Sprintf(w.ErrorKind.LocalizedString(p)))
}

func deepSortResult(result *jsonschema.OutputUnit) {
	// The validation library converts schema locations to absolute path which is not practical
	// for tests. Therefore, it is overridden here so errors do not contain pwd base path.
	result.AbsoluteKeywordLocation = ""

	if result.Error != nil {
		result.Error.Kind = &wrappedErrorKind{result.Error.Kind}
	}

	slices.SortFunc(result.Errors, func(a, b jsonschema.OutputUnit) int {
		return strings.Compare(a.InstanceLocation, b.InstanceLocation)
	})

	for i := range result.Errors {
		deepSortResult(&result.Errors[i])
	}
}

// ValidateMapJSONResult validates the given map and returns the result as JSON.
// The JSON output is sorted to make it deterministic for JSON/text comparison.
// This method is only useful for testing, use ValidateMap for regular validation.
func (s *Schema) validateMapJSONResult(data any) []byte {
	err := s.s.Validate(data)

	var validationErr *jsonschema.ValidationError
	if err != nil && errors.As(err, &validationErr) {
		outputUnit := validationErr.DetailedOutput()
		deepSortResult(outputUnit)

		result, err := json.MarshalIndent(outputUnit, "", "  ")
		if err != nil {
			return fmt.Appendf(nil, "Error marshalling error output: %v", err)
		}

		return result
	}

	return []byte("")
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
