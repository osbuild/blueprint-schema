package blueprint

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrParsing = errors.New("parsing error")

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

// parseUGIDstr parses a user/group ID from a string. It returns the
// user/group ID as an int64 if it is a number, or the string itself
// if it is not a number. If the string is empty, it returns nil.
func parseUGIDstr(s string) any {
	if s == "" {
		return nil
	}

	if i, err := strconv.ParseInt(s, 10, 0); err == nil {
		return i
	}

	return s
}

// parseUGIDany parses a string as either a username/groupname or a int64 as UID/GID.
// Returns empty string if the input is empty or nil.
func parseUGIDany(a any) string {
	if a == nil {
		return ""
	}

	if num, ok := a.(int64); ok && num >= 0 {
		return fmt.Sprintf("%d", num)
	}

	if str, ok := a.(string); ok && str != "" {
		return str
	}

	return ""
}

// parseOctalString parses a string as an octal number and returns the integer value.
// It parses strings not starting with '0' as valid octal numbers.
func parseOctalString(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	var s string
	if !strings.HasPrefix(str, "0") {
		s = "0" + str
	} else {
		s = str
	}

	value, err := strconv.ParseInt(s, 8, 0)
	if err != nil {
		return 0, fmt.Errorf("%w: string %q is not a valid octal number", ErrParsing, str)
	}

	return int(value), nil
}

func joinNonEmpty(delimiter string, parts ...string) string {
	var sb strings.Builder
	for _, part := range parts {
		if part != "" {
			if sb.Len() > 0 {
				sb.WriteString(delimiter)
			}
			sb.WriteString(part)
		}
	}
	return sb.String()
}
