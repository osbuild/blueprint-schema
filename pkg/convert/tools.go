package convert

import "strings"

// splitStringN splits a string into n parts using the specified delimiter.
// If the string has fewer than n parts, the remaining parts are filled with empty strings.
func splitStringEmptyN(s string, delimiter string, n int) []string {
	parts := strings.SplitN(s, delimiter, n)
	result := make([]string, n)
	for i := range n {
		if i < len(parts) {
			result[i] = parts[i]
		} else {
			result[i] = ""
		}
	}
	return result
}
