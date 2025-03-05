package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRewrapText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "\n",
			expected: "",
		},
		{
			input:    "\n\n",
			expected: "",
		},
		{
			input:    "Text with no linebreaks",
			expected: "Text with no linebreaks",
		},
		{
			input:    "Text with one\nlinebreak",
			expected: "Text with one linebreak",
		},
		{
			input:    "Text with\ntwo\nlinebreaks",
			expected: "Text with two linebreaks",
		},
		{
			input:    "Text with a trailing linebreak\n",
			expected: "Text with a trailing linebreak",
		},
		{
			input:    "Text with one\n\nparagraph\n",
			expected: "Text with one\n\nparagraph",
		},
		{
			input:    "Text with one\nlinebreak and\n\none paragraph",
			expected: "Text with one linebreak and\n\none paragraph",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := rewrapText(test.input)
			if cmp.Diff(test.expected, actual) != "" {
				t.Errorf("unexpected output (-want +got):\n%s", cmp.Diff(test.expected, actual))
			}
		})
	}
}
