package blueprint

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kaptinlin/jsonschema"
)

func TestSortSchemaSlice(t *testing.T) {
	tests := []struct {
		input  []jsonschema.List
		output []jsonschema.List
	}{
		{
			input: []jsonschema.List{
				{
					Errors: map[string]string{
						"Properties": "Do not touch",
					},
				},
			},
			output: []jsonschema.List{
				{
					Errors: map[string]string{
						"Properties": "Do not touch",
					},
				},
			},
		},
		{
			input: []jsonschema.List{
				{
					Errors: map[string]string{
						"properties": "Property 'a' do not match its schema",
					},
				},
			},
			output: []jsonschema.List{
				{
					Errors: map[string]string{
						"properties": "Property 'a' do not match its schema",
					},
				},
			},
		},
		{
			input: []jsonschema.List{
				{
					Errors: map[string]string{
						"properties": "Properties 'b', 'a' do not match their schemas",
					},
				},
			},
			output: []jsonschema.List{
				{
					Errors: map[string]string{
						"properties": "Properties 'a', 'b' have problems",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		sortSchemaSlice(tt.input)

		if diff := cmp.Diff(tt.output, tt.input); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
