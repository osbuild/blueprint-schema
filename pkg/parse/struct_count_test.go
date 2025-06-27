package parse

import (
	"reflect"
	"testing"
)

func TestCountSetFieldsRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
	}{
		{
			name:     "Empty struct",
			input:    struct{}{},
			expected: 0,
		},
		{
			name:     "Single set field",
			input:    struct{ A string }{A: "value"},
			expected: 1,
		},
		{
			name: "Multiple set fields",
			input: struct {
				A string
				B int
			}{A: "value", B: 42},
			expected: 2,
		},
		{
			name: "Nested struct",
			input: struct {
				A string
				B struct{ C int }
			}{A: "value", B: struct{ C int }{C: 10}},
			expected: 2,
		},
		{
			name:     "Pointer to struct",
			input:    struct{ A *string }{A: func(s string) *string { return &s }("value")},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := reflect.ValueOf(tt.input)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}
			count := countRecursive(val)
			if count != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, count)
			}
		})
	}
}
