package ubp

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

func TestJSONUnmarshalDNFRepoUsage(t *testing.T) {
	tests := []struct {
		input    string
		expected DNFRepoUsage
	}{
		{
			input:    `{}`,
			expected: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(true)},
		},
		{
			input:    `{"configure": false, "install": false}`,
			expected: DNFRepoUsage{Configure: ptr.To(false), Install: ptr.To(false)},
		},
		{
			input:    `{"configure": true, "install": false}`,
			expected: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(false)},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var usage DNFRepoUsage
			err := json.Unmarshal([]byte(test.input), &usage)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if diff := cmp.Diff(test.expected, usage); diff != "" {
				t.Errorf("Unmarshal mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestJSONMarshalDNFRepoUsage(t *testing.T) {
	tests := []struct {
		input  DNFRepoUsage
		output string
	}{
		{
			input:  DNFRepoUsage{},
			output: `{}`,
		},
		{
			input:  DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(true)},
			output: `{}`,
		},
		{
			input:  DNFRepoUsage{Configure: ptr.To(false), Install: ptr.To(false)},
			output: `{"configure":false,"install":false}`,
		},
		{
			input:  DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(false)},
			output: `{"install":false}`,
		},
	}

	for _, test := range tests {
		t.Run(test.output, func(t *testing.T) {
			output, err := json.Marshal(test.input)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			if diff := cmp.Diff(test.output, string(output)); diff != "" {
				t.Errorf("Marshal mismatch (-input +got):\n%s", diff)
			}
		})
	}
}
