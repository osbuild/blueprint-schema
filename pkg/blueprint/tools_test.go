package blueprint

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

func TestInt64ToVersion(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "1.0.0"},
		{1, "1.0.1"},
		{2, "1.0.2"},
		{3, "1.0.3"},
		{1368473400, "1.20881.16184"},
		{1747392766, "1.26663.6398"},
		{0x00000001FFFFFFFF, "2.65535.65535"},
	}

	for _, test := range tests {
		result := int64ToVersion(test.input)

		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Errorf("int64ToVersion(%d) mismatch (-want +got):\n%s", test.input, diff)
		}
	}
}
