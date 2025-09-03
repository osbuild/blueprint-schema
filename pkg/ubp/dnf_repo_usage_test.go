package ubp

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

func TestDefaultValues(t *testing.T) {
	ubp := &Blueprint{
		DNF: DNF{
			Repositories: []DNFRepository{
				{
					Usage: DNFRepoUsage{},
				},
			},
		},
	}

	if err := PopulateDefaults(ubp); err != nil {
		t.Fatalf("Failed to populate defaults: %v", err)
	}

	if ubp.DNF.Repositories[0].Usage.Configure == nil {
		t.Error("Expected Configure to be set")
	} else if *ubp.DNF.Repositories[0].Usage.Configure != true {
		t.Errorf("Expected Configure to be true, got %v", *ubp.DNF.Repositories[0].Usage.Configure)
	}

	if ubp.DNF.Repositories[0].Usage.Install == nil {
		t.Error("Expected Install to be set")
	} else if *ubp.DNF.Repositories[0].Usage.Install != true {
		t.Errorf("Expected Install to be true, got %v", *ubp.DNF.Repositories[0].Usage.Install)
	}
}

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
	type DRU = DNFRepoUsage

	type tru struct {
		Usage DRU `json:"usage,omitzero"`
	}

	tests := []struct {
		input  tru
		output string
	}{
		{
			input:  tru{Usage: DNFRepoUsage{}},
			output: `{}`,
		},
		{
			input:  tru{Usage: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(true)}},
			output: `{}`,
		},
		{
			input:  tru{Usage: DNFRepoUsage{Configure: ptr.To(false), Install: ptr.To(false)}},
			output: `{"usage":{"configure":false,"install":false}}`,
		},
		{
			input:  tru{Usage: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(false)}},
			output: `{"usage":{"install":false}}`,
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

func TestDNFRepoUsageIsZero(t *testing.T) {
	tests := []struct {
		n    string
		in   DNFRepoUsage
		zero bool
	}{
		// zero
		{n: "empty", in: DNFRepoUsage{}, zero: true},
		{n: "true-true", in: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(true)}, zero: true},
		{n: "true-nil", in: DNFRepoUsage{Configure: ptr.To(true)}, zero: true},

		// non-zero
		{n: "false-nil", in: DNFRepoUsage{Configure: ptr.To(false)}, zero: false},
		{n: "false-false", in: DNFRepoUsage{Configure: ptr.To(false), Install: ptr.To(false)}, zero: false},
		{n: "true-false", in: DNFRepoUsage{Configure: ptr.To(true), Install: ptr.To(false)}, zero: false},
		{n: "false-true", in: DNFRepoUsage{Configure: ptr.To(false), Install: ptr.To(true)}, zero: false},
	}

	for _, test := range tests {
		t.Run(test.n, func(t *testing.T) {
			if got := test.in.IsZero(); got != test.zero {
				t.Errorf("IsZero() = %v, want %v", got, test.zero)
			}
		})
	}
}
