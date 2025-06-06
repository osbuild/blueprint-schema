package blueprint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	"sigs.k8s.io/yaml"
	goyaml "sigs.k8s.io/yaml/goyaml.v3"
)

// splitYAMLDocuments takes a byte slice containing one or more YAML documents
// separated by "---" and returns a slice of byte slices, each representing
// a single YAML document.
func splitYAMLDocuments(input []byte) ([][]byte, error) {
	var documents [][]byte
	decoder := goyaml.NewDecoder(bytes.NewReader(input))

	for {
		var node goyaml.Node
		err := decoder.Decode(&node)

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to decode YAML document: %w", err)
		}

		var buffer bytes.Buffer
		encoder := goyaml.NewEncoder(&buffer)
		encoder.SetIndent(2)

		if err := encoder.Encode(&node); err != nil {
			return nil, fmt.Errorf("failed to re-encode YAML node: %w", err)
		}

		documents = append(documents, buffer.Bytes())
	}

	return documents, nil
}

// isFieldSet determines if a reflect.Value is "set", meaning it's not its
// type's zero value. For pointers, slices, maps, interfaces, channels,
// and functions, this means it's not nil. For other types (bool, int,
// string, struct, array), it means it's not the standard zero value (e.g.,
// false, 0, "", or a struct/array with all fields/elements zero).
func isFieldSet(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return !v.IsNil()
	default:
		return !v.IsZero()
	}
}

// nonZeroFields returns a slice of strings containing the names of all exported fields
// that are not their zero value.
func nonZeroFields(s *Blueprint) []string {
	val := reflect.ValueOf(s).Elem()

	var nonZeroFieldNames []string
	structType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		structField := structType.Field(i)

		// Skip unexported fields.
		if structField.PkgPath != "" {
			continue
		}

		// Check if the field is set (not its zero value).
		if isFieldSet(fieldVal) {
			nonZeroFieldNames = append(nonZeroFieldNames, structField.Name)
		}
	}

	return nonZeroFieldNames
}

var ErrMergeField = errors.New("cannot merge field into blueprint")

// UnmarshalYAML splits multiple YAML documents, convert them one by one to JSON then
// uses the standard library JSON decoder to unmarshal them into Blueprint.
//
// Document buffers can be passed as variadic arguments, in a single ocument concatenated
// with "---" or any combination of multiple documents.
//
// Documents are safely merged into each other sequentially. Only selected top-level fields
// are supported merged.
//
// Note the blueprint types do not use any YAML Go struct tags, this is because
// the JSON tags are used instead. This ensures consistency between JSON and YAML
// representations, as YAML is a superset of JSON.
func UnmarshalYAML(yamlBufs ...[]byte) (*Blueprint, error) {
	var b *Blueprint
	var done struct {
		Registration bool
	}

	for _, buf := range yamlBufs {
		docs, err := splitYAMLDocuments(buf)
		if err != nil {
			return nil, fmt.Errorf("failed to split YAML documents: %w", err)
		}

		for i, doc := range docs {
			if len(doc) == 0 {
				continue
			}

			tmp := Blueprint{}
			if err := yaml.Unmarshal(doc, &tmp); err != nil {
				return nil, fmt.Errorf("failed to unmarshal YAML document %d: %w", i, err)
			}

			// The first document initializes the blueprint.
			if b == nil {
				b = &tmp
				continue
			}

			// The merging code
			if tmp.Registration != nil {
				if done.Registration {
					return nil, fmt.Errorf("%w: cannot merge twice: %q", ErrMergeField, "Registration")
				}
				b.Registration = tmp.Registration
				tmp.Registration = nil
				done.Registration = true
			}

			// With all the fields that can be merged gone, the rest must be nil or empty.
			fields := nonZeroFields(&tmp)
			if len(fields) > 0 {
				return nil, fmt.Errorf("%w: %q", ErrMergeField, fields)
			}
		}
	}

	return b, nil
}

// ReadYAML reads into a buffer and calls UnmarshalYAML. Read UnmarshalYAML for more details.
func ReadYAML(reader io.Reader) (*Blueprint, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return UnmarshalYAML(buf.Bytes())
}

// MarshalYAML uses JSON encoder to marshal the object into JSON and then converts JSON to YAML.
// No YAML Go struct tags are necessary as JSON tags are used.
func MarshalYAML(b *Blueprint) ([]byte, error) {
	return yaml.Marshal(b)
}

// WriteYAML calls MarshalYAML and writes the result to the writer.
func WriteYAML(b *Blueprint, writer io.Writer) error {
	data, err := yaml.Marshal(b)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}

// unmarshalJSON uses JSON decoder to unmarshal into an object.
func unmarshalJSON(data []byte) (*Blueprint, error) {
	b := new(Blueprint)
	return b, json.Unmarshal(data, b)
}

// readJSON calls UnmarshalJSON after reading into a buffer.
func readJSON(reader io.Reader) (*Blueprint, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalJSON(buf.Bytes())
}

// marshalJSON uses JSON encoder to marshal the object into JSON.
// Output can be optionaly indented.
func marshalJSON(b *Blueprint, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(b, "", "  ")
	}

	return json.Marshal(b)
}

// writeJSON calls MarshalJSON and writes the result to the writer.
// Output can be optionaly indented.
func writeJSON(b *Blueprint, writer io.Writer, indent bool) error {
	data, err := marshalJSON(b, indent)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	return err
}

// ConvertJSONtoYAML converts JSON to YAML.
func ConvertJSONtoYAML(data []byte) ([]byte, error) {
	return yaml.JSONToYAML(data)
}

// ConvertYAMLtoJSON converts YAML to JSON. Output is not indented.
func ConvertYAMLtoJSON(data []byte) ([]byte, error) {
	return yaml.YAMLToJSON(data)
}
