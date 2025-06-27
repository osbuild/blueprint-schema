package parse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/ubp"
	bp "github.com/osbuild/blueprint/pkg/blueprint"
	"sigs.k8s.io/yaml"
)

var ErrParsingExplanation = errors.New("parsing errors, some are false positives")

func UnmarshalAny(buf []byte) (*ubp.Blueprint, *bp.Blueprint, error, error) {
	// Try UBP YAML
	ubpYAML, ubpErrYAML := UnmarshalYAML(buf)
	ubpCountYAML := countSetFieldsRecursive(ubpYAML)

	// Try UBP JSON
	ubpJSON, ubpErrJSON := UnmarshalJSON(buf)
	ubpCountJSON := countSetFieldsRecursive(ubpJSON)

	// Try BP TOML
	bpTempTOML := new(bp.Blueprint)
	bpErrTOML := toml.Unmarshal(buf, bpTempTOML)
	imTOML := conv.NewInternalImporter(bpTempTOML)
	bpWarnTOML := imTOML.Import()
	bpTOML := imTOML.Result()
	bpCountTOML := countSetFieldsRecursive(bpTOML)

	// Try BP JSON
	bpTempJSON := new(bp.Blueprint)
	bpErrJSON := json.Unmarshal(buf, bpTempJSON)
	imJSON := conv.NewInternalImporter(bpTempJSON)
	bpWarnJSON := imJSON.Import()
	bpJSON := imJSON.Result()
	bpCountJSON := countSetFieldsRecursive(bpJSON)

	maxCount := max(ubpCountYAML, ubpCountJSON, bpCountTOML, bpCountJSON)
	err := errors.Join(
		fmt.Errorf("YAML: %w", ubpErrYAML),
		fmt.Errorf("JSON: %w", ubpErrJSON),
		fmt.Errorf("TOML: %w", bpErrTOML),
		fmt.Errorf("JSON: %w", bpErrJSON),
	)
	warn := errors.Join(bpWarnTOML, bpWarnJSON)

	if ubpErrYAML == nil && ubpCountYAML == maxCount {
		return ubpYAML, nil, nil, warn
	} else if ubpErrJSON == nil && ubpCountJSON == maxCount {
		return ubpJSON, nil, nil, warn
	} else if bpErrTOML == nil && bpCountTOML == maxCount {
		return bpTOML, bpTempTOML, nil, warn
	} else if bpErrJSON == nil && bpCountJSON == maxCount {
		return bpJSON, bpTempJSON, nil, warn
	} else {
		countErr := fmt.Errorf("fields set: UBPY:%d UBPJ:%d BPT:%d BPJ:%d", ubpCountYAML, ubpCountJSON, bpCountTOML, bpCountJSON)
		return nil, nil, errors.Join(ErrParsingExplanation, err, countErr), warn
	}
}

// UnmarshalYAML loads a blueprint from YAML data. It converts YAML into JSON first,
// and then unmarshals it into a Blueprint object. This is done to ensure that the
// YAML representation is consistent with the JSON representation.
//
// Uses sigs.k8s.io/yaml package for YAML parsing, for the API guarantees and
// compatibility read https://pkg.go.dev/sigs.k8s.io/yaml#Unmarshal.
func UnmarshalYAML(buf []byte) (*ubp.Blueprint, error) {
	b := new(ubp.Blueprint)

	if err := yaml.Unmarshal(buf, b); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML blueprint: %w", err)
	}

	return b, nil
}

// ReadYAML reads into a buffer and calls UnmarshalYAML. Read UnmarshalYAML for more details.
func ReadYAML(reader io.Reader) (*ubp.Blueprint, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return UnmarshalYAML(buf.Bytes())
}

// MarshalYAML uses JSON encoder to marshal the object into JSON and then converts JSON to YAML.
// No YAML Go struct tags are necessary as JSON tags are used.
//
// Uses sigs.k8s.io/yaml package for YAML encoding, for the API guarantees and
// compatibility read https://pkg.go.dev/sigs.k8s.io/yaml#Unmarshal.
func MarshalYAML(b *ubp.Blueprint) ([]byte, error) {
	return yaml.Marshal(b)
}

// WriteYAML calls MarshalYAML and writes the result to the writer.
func WriteYAML(b *ubp.Blueprint, writer io.Writer) error {
	data, err := yaml.Marshal(b)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}

// UnmarshalJSON uses JSON decoder to unmarshal into an object.
//
// Do not use this function for user-facing data.
func UnmarshalJSON(data []byte) (*ubp.Blueprint, error) {
	b := new(ubp.Blueprint)
	return b, json.Unmarshal(data, b)
}

// ReadJSON calls UnmarshalJSON after reading into a buffer.
//
// Do not use this function for user-facing data.
func ReadJSON(reader io.Reader) (*ubp.Blueprint, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return UnmarshalJSON(buf.Bytes())
}

// MarshalJSON uses JSON encoder to marshal the object into JSON.
// Output can be optionaly indented.
//
// Do not use this function for user-facing data.
func MarshalJSON(b *ubp.Blueprint, indent bool) ([]byte, error) {
	if indent {
		return json.MarshalIndent(b, "", "  ")
	}

	return json.Marshal(b)
}

// WriteJSON calls MarshalJSON and writes the result to the writer.
// Output can be optionaly indented.
//
// Do not use this function for user-facing data.
func WriteJSON(b *ubp.Blueprint, writer io.Writer, indent bool) error {
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
