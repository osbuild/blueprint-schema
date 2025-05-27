package blueprint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"sigs.k8s.io/yaml"
	goyaml "sigs.k8s.io/yaml/goyaml.v3"
)

// SplitYAMLDocuments takes a byte slice containing one or more YAML documents
// separated by "---" and returns a slice of byte slices, each representing
// a single YAML document.
func SplitYAMLDocuments(input []byte) ([][]byte, error) {
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

// UnmarshalYAML splits multiple YAML documents, convert them one by one to JSON then
// uses the standard library JSON decoder to unmarshal them into Blueprint. 
//
// Note the blueprint types do not use any YAML Go struct tags, this is because
// the JSON tags are used instead. This ensures consistency between JSON and YAML
// representations, as YAML is a superset of JSON.
func UnmarshalYAML(data []byte) (*Blueprint, error) {
	b := new(Blueprint)

	docs, err := SplitYAMLDocuments(data)
	if err != nil {
		return nil, fmt.Errorf("failed to split YAML documents: %w", err)
	}

	for _, doc := range docs {
		if err := yaml.Unmarshal(doc, b); err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML document: %w", err)
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
