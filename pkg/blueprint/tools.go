package blueprint

import (
	"fmt"
	"strings"
)

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

// int64ToVersion converts a uint64 value to a version string in the format "x.y.z".
func int64ToVersion(input uint64) string {
	z := uint16(input & 0xFFFF)
	y := uint16(input >> 16 & 0xFFFF)
	x := uint32(input >> 32 & 0xFFFFFFFF)

	return fmt.Sprintf("%d.%d.%d", x+1, y, z)
}
