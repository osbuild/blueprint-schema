package blueprint

import (
	"testing"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

func TestExpireDateToEpoch(t *testing.T) {
	tests := []struct {
		input    ExpireDate
		expected int
	}{
		{"1970-01-01", 0},
		{"1970-01-02", 1},
		{"1980-01-15", 3666},
		{"2023-10-01", 19631},
		{"2023-10-01T00:00:00Z", 19631},
		{"2023-10-01T12:34:56Z", 19631},
	}

	for _, test := range tests {
		result, err := ExpireDateToEpochDays(test.input)
		if err != nil && test.expected != 0 {
			t.Errorf("Expected no error for input %s, got %v", test.input, err)
		}

		if result != test.expected {
			t.Errorf("Expected %d, got %d for input %s", test.expected, result, test.input)
		}
	}
}

func TestParseExpireDate(t *testing.T) {
	tests := []struct {
		input    *int
		expected *ExpireDate
	}{
		{ptr.To(0), ptr.To("1970-01-01T00:00:00Z")},
		{ptr.To(1), ptr.To("1970-01-02T00:00:00Z")},
		{ptr.To(3666), ptr.To("1980-01-15T00:00:00Z")},
		{ptr.To(19631), ptr.To("2023-10-01T00:00:00Z")},
	}

	for _, test := range tests {
		result := ParseExpireDate(test.input)
		if result == nil && test.expected != nil {
			t.Errorf("Expected %s, got nil for input %v", *test.expected, test.input)
			continue
		}

		if result != nil && *result != *test.expected {
			t.Errorf("Expected %s, got %s for input %v", *test.expected, *result, test.input)
		}
	}
}
