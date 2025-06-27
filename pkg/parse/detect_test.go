package parse

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDetectObject(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		format   detectedFormat
		expected any
	}{
		{
			"Empty", "", formatUnknown,
			map[string]any{},
		},
		{
			"Invalid", "this is not a valid format", formatUnknown,
			map[string]any{},
		},
		{
			"JSON", `{"key": "value"}`, formatJSON,
			map[string]any{
				"key": "value",
			},
		},
		{
			"YAML", "key: value\n", formatYAML,
			map[string]any{
				"key": "value",
			},
		},
		{
			"TOML", `key = "value"`, formatTOML,
			map[string]any{
				"key": "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, obj := detectFormat([]byte(tt.input))
			if obj == nil && tt.expected != nil {
				t.Errorf("expected non-nil object, got nil")
			} else if diff := cmp.Diff(tt.expected, obj); diff != "" {
				t.Logf("expected format: %s, got: %s", tt.format, f)
				t.Errorf("object mismatch (-expected +got):\n%s", diff)
			} else if f != tt.format {
				t.Errorf("expected format %s, got %s", tt.format, f)
			}
		})
	}
}

func TestDetectUBP(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected Type
	}{
		{
			map[string]any{}, TypeUnknown,
		},
		{
			map[string]any{"key": "value"}, TypeUnknown,
		},
		{
			map[string]any{
				"distribution": "fedora-13",
			}, TypeUBP,
		},
		{
			map[string]any{
				"distro": "fedora-13",
			}, TypeBP,
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.expected), func(t *testing.T) {
			result := DetectType(tt.input)
			if result != tt.expected {
				t.Errorf("expected struct type %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestDetectFixtures(t *testing.T) {
	tests := []struct {
		filename string
		format   detectedFormat
		str      Type
	}{
		{"../../testdata/all-fields.in.yaml", formatYAML, TypeUBP},
		{"../../testdata/invalid-all-empty.in.yaml", formatYAML, TypeUBP},
		{"../../testdata/valid-empty.in.yaml", formatUnknown, TypeUnknown},
		{"../../testdata/valid-empty-j.in.json", formatUnknown, TypeUnknown},
		{"../../testdata/small.json", formatJSON, TypeUBP},
		{"../../testdata/legacy-small.json", formatJSON, TypeBP},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			buf, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", tt.filename, err)
			}

			format, data := detectFormat(buf)
			if format != tt.format {
				t.Errorf("expected format %s, got %s", tt.format, format)
			}

			str := DetectType(data)
			if str != tt.str {
				t.Errorf("expected struct type %s, got %s, data: %+v", tt.str, str, data)
			}
		})
	}
}
