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
			expected: "\n",
		},
		{
			input:    "\n\n",
			expected: "\n\n",
		},
		{
			input:    "Text with no linebreaks",
			expected: "Text with no linebreaks\n",
		},
		{
			input:    "Text with one\nlinebreak",
			expected: "Text with one linebreak\n",
		},
		{
			input:    "Text with\ntwo\nlinebreaks",
			expected: "Text with two linebreaks\n",
		},
		{
			input:    "Text with a trailing linebreak\n",
			expected: "Text with a trailing linebreak\n",
		},
		{
			input:    "Text with one\n\nparagraph\n",
			expected: "Text with one\n\nparagraph\n",
		},
		{
			input:    "Text with one\nlinebreak and\n\none paragraph",
			expected: "Text with one linebreak and\n\none paragraph\n",
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
