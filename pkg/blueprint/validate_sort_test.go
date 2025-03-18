package blueprint

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortSchemaSlice(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "",
			output: "",
		},
		{
			input:  "untouched string",
			output: "untouched string",
		},
		{
			input:  "untouched 'quoted string'",
			output: "untouched 'quoted string'",
		},
		{
			input:  "sort: 'c', 'b', 'a'",
			output: "sort: 'a', 'b', 'c'",
		},
		{
			input:  "sort: 'aaa', 'b', 'aaa'",
			output: "sort: 'aaa', 'aaa', 'b'",
		},
		{
			input:  "invalid case: ' invalid",
			output: "invalid case: ' invalid",
		},
	}

	for _, tt := range tests {
		got := sortQuotedSubstrings(tt.input)

		if diff := cmp.Diff(tt.output, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
