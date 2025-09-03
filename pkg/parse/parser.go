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

// AnyFormat represents the format of a blueprint that can be parsed.
type AnyFormat int

const (
	AnyFormatUnknown AnyFormat = iota // Unknown format
	AnyFormatUBPYAML                  // UBP YAML format
	AnyFormatUBPJSON                  // UBP JSON format
	AnyFormatBPTOML                   // BP TOML format
	AnyFormatBPJSON                   // BP JSON format
)

// AnyDetails contains details about the parsing process of a blueprint in any format.
// It includes the format detected, warnings encountered, whether the blueprint was converted,
// and the intermediate blueprint if conversion was necessary.
type AnyDetails struct {
	Format       AnyFormat     // Detected format of the blueprint
	Warnings     error         // Warnings encountered during conversion
	Intermediate *bp.Blueprint // Intermediate blueprint if conversion was necessary
}

func (f AnyFormat) String() string {
	switch f {
	case AnyFormatUBPYAML:
		return "UBP-YAML"
	case AnyFormatUBPJSON:
		return "UBP-JSON"
	case AnyFormatBPTOML:
		return "BP-TOML"
	case AnyFormatBPJSON:
		return "BP-JSON"
	default:
		return "Unknown"
	}
}

var ErrParsingDetection = errors.New("parsing and detecting blueprint failed")

// UnmarshalAny attempts to unmarshal a blueprint from a byte slice in any format.
// It performs heuristic detection of UBP YAML, UBP JSON, BP TOML, and BP JSON formats
// via strict unmarshalling. If parsing fails, it returns a list of wrapped errors
// indicating the failure of each format attempt.
//
// When BP format is detected, it is converted to UBP automatically and default values
// are populated.
//
// To get some insights about the parsing process, you can pass an `AnyDetails` pointer as an
// argument. This allows finding out which format was detected, whether there were any warnings
// during the conversion, and what the intermediate blueprint was if conversion was necessary.
func UnmarshalAny(buf []byte, details *AnyDetails) (*ubp.Blueprint, error) {
	if details == nil {
		details = &AnyDetails{}
	}

	errs := make([]error, 0, 5)
	errs = append(errs, ErrParsingDetection)

	// Try UBP YAML
	ubpData, err := UnmarshalStrictYAML(buf)
	if err == nil {
		details.Format = AnyFormatUBPYAML

		return ubpData, nil
	} else {
		errs = append(errs, fmt.Errorf("%w (UBP-YAML)", err))
	}

	// Try UBP JSON
	ubpData, err = UnmarshalStrictJSON(buf)
	if err == nil {
		details.Format = AnyFormatUBPJSON

		return ubpData, nil
	} else {
		errs = append(errs, fmt.Errorf("%w (UBP-JSON)", err))
	}

	// Try BP JSON
	bpData := new(bp.Blueprint)
	dec := json.NewDecoder(bytes.NewReader(buf))
	dec.DisallowUnknownFields()
	err = dec.Decode(bpData)
	if err == nil {
		details.Format = AnyFormatBPJSON

		importer := conv.NewInternalImporter(bpData)
		imData, warn := importer.Import()
		details.Intermediate = bpData
		details.Warnings = warn

		return imData, nil
	} else {
		errs = append(errs, fmt.Errorf("%w (BP-JSON)", err))
	}

	// Try BP TOML (this parser does not have a strict mode, so it comes last)
	bpData = new(bp.Blueprint)
	err = toml.Unmarshal(buf, bpData)
	if err == nil {
		details.Format = AnyFormatBPTOML

		importer := conv.NewInternalImporter(bpData)
		imData, warn := importer.Import()
		details.Intermediate = bpData
		details.Warnings = warn

		return imData, nil
	} else {
		errs = append(errs, fmt.Errorf("%w (BP-TOML)", err))
	}

	return nil, errors.Join(errs...)
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

// UnmarshalStrictYAML loads a blueprint from YAML data, but it does not allow unknown fields.
// Read UnmarshalYAML for more details.
func UnmarshalStrictYAML(buf []byte) (*ubp.Blueprint, error) {
	b := new(ubp.Blueprint)

	if err := yaml.UnmarshalStrict(buf, b); err != nil {
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

// UnmarshalStrictJSON uses JSON decoder to unmarshal into an object, but it does
// not allow unknown fields. See UnmarshalJSON for more details.
func UnmarshalStrictJSON(data []byte) (*ubp.Blueprint, error) {
	b := new(ubp.Blueprint)
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(b); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON blueprint: %w", err)
	}

	return b, nil
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
