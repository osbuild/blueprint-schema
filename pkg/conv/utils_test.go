package conv

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

func TestSplitEnVr(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantPart1 string
		wantPart2 string
	}{
		{
			name:      "simple package",
			input:     "vim-1.0-1",
			wantPart1: "vim",
			wantPart2: "1.0-1",
		},
		{
			name:      "package with dash",
			input:     "vim-enhanced-1.0-1",
			wantPart1: "vim-enhanced",
			wantPart2: "1.0-1",
		},
		{
			name:      "package with everything",
			input:     "vim-enhanced-9.1.866-1.fc41.x86_64",
			wantPart1: "vim-enhanced",
			wantPart2: "9.1.866-1.fc41.x86_64",
		},
		{
			name:      "number and two hyphens",
			input:     "grub2-utils",
			wantPart1: "grub2-utils",
			wantPart2: "",
		},
		{
			name:      "exactly one hyphen",
			input:     "name-version",
			wantPart1: "name-version",
			wantPart2: "",
		},
		{
			name:      "exactly two hyphens",
			input:     "name-version-release",
			wantPart1: "name-version-release",
			wantPart2: "",
		},
		{
			name:      "no hyphens",
			input:     "package",
			wantPart1: "package",
			wantPart2: "",
		},
		{
			name:      "empty string",
			input:     "",
			wantPart1: "",
			wantPart2: "",
		},
		{
			name:      "trailing hyphen with version",
			input:     "a-b-1.0-",
			wantPart1: "a-b",
			wantPart2: "1.0-",
		},
		{
			name:      "trailing hyphen",
			input:     "a-b-",
			wantPart1: "a-b-",
			wantPart2: "",
		},
		{
			name:      "leading hyphen", // suboptimal case
			input:     "-a-b-1.0",
			wantPart1: "-a",
			wantPart2: "b-1.0",
		},
		{
			name:      "consecutive hyphens", // suboptimal case
			input:     "a--b-1.0",
			wantPart1: "a-",
			wantPart2: "b-1.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotPart1, gotPart2 := splitEnVr(tc.input)
			if gotPart1 != tc.wantPart1 || gotPart2 != tc.wantPart2 {
				t.Errorf("SplitAtSecondLastDash(%q):\ngot:  part1=%q, part2=%q\nwant: part1=%q, part2=%q",
					tc.input, gotPart1, gotPart2, tc.wantPart1, tc.wantPart2)
			}
		})
	}
}

func TestParseUGID(t *testing.T) {
	tests := []struct {
		str string
		an  any
	}{
		{"", nil},
		{"1000", int64(1000)},
		{"0", int64(0)},
		{"1000.1", "1000.1"},
		{"user", "user"},
		{"group", "group"},
		{"1000user", "1000user"},
		{"user1000", "user1000"},
		{"1000user1000", "1000user1000"},
	}

	for _, test := range tests {
		result := parseUGIDstr(test.str)

		if diff := cmp.Diff(test.an, result); diff != "" {
			t.Errorf("parseUGIDstr(%q) mismatch (-want +got):\n%s", test.str, diff)
		}
	}
	for _, test := range tests {
		result := parseUGIDany(test.an)

		if diff := cmp.Diff(test.str, result); diff != "" {
			t.Errorf("parseUGIDany(%v) mismatch (-want +got):\n%s", test.an, diff)
		}
	}
}

func TestJoinNonEmptyStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", ""}, "a"},
		{[]string{"", "b"}, "b"},
		{[]string{"a", "b"}, "a-b"},
		{[]string{"a", "", "b"}, "a-b"},
		{[]string{"", "a", "b"}, "a-b"},
	}

	for _, test := range tests {
		result := joinNonEmpty("-", test.input...)

		if diff := cmp.Diff(test.expected, result); diff != "" {
			t.Errorf("joinNonEmpty(%v) mismatch (-want +got):\n%s", test.input, diff)
		}
	}
}
