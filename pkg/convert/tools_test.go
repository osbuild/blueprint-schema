package convert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitStringEmptyN(t *testing.T) {
	tests := []struct {
		input     string
		delimiter string
		n         int
		expected  []string
	}{
		{"", ":", 1, []string{""}},
		{"a", ":", 1, []string{"a"}},
		{"a", ":", 2, []string{"a", ""}},
		{"a:b", ":", 2, []string{"a", "b"}},
		{"a:b", ":", 3, []string{"a", "b", ""}},
		{"a:b:c", ":", 2, []string{"a", "b:c"}},
	}

	for _, test := range tests {
		result := splitStringEmptyN(test.input, test.delimiter, test.n)

		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Errorf("splitStringEmptyN(%q, %q, %d) mismatch (-want +got):\n%s", test.input, test.delimiter, test.n, diff)
		}
	}
}
