package blueprint

import (
	"bytes"
	"encoding/json"
	"io"

	"sigs.k8s.io/yaml"
)

// UnmarshalYAML converts YAML to JSON then uses JSON decoder to unmarshal into an object.
// No YAML Go struct tags are necessary as JSON tags are used.
func UnmarshalYAML(data []byte) (*Blueprint, error) {
	b := new(Blueprint)
	return b, yaml.Unmarshal(data, b)
}

// ReadYAML reads into a buffer and calls UnmarshalYAML.
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

// UnmarshalJSON uses JSON decoder to unmarshal into an object.
func UnmarshalJSON(data []byte) (*Blueprint, error) {
	b := new(Blueprint)
	return b, json.Unmarshal(data, b)
}

// ReadJSON calls UnmarshalJSON after reading into a buffer.
func ReadJSON(reader io.Reader) (*Blueprint, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return UnmarshalJSON(buf.Bytes())
}

// MarshalJSON uses JSON encoder to marshal the object into JSON.
// Output can be optionaly indented.
func MarshalJSON(b *Blueprint, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(b, "", "  ")
	}

	return json.Marshal(b)
}

// WriteJSON calls MarshalJSON and writes the result to the writer.
// Output can be optionaly indented.
func WriteJSON(b *Blueprint, writer io.Writer, indent bool) error {
	data, err := MarshalJSON(b, indent)
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
