package conv

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

// splitEnVr splits a NVRE string to Epoch+Name and Version+Release. Performs "best effort"
// heuristic, can generate incorrect split.
func splitEnVr(s string) (string, string) {
	lastDash := strings.LastIndex(s, "-")
	if lastDash == -1 {
		return s, ""
	}

	secondLastDash := strings.LastIndex(s[:lastDash], "-")
	if secondLastDash == -1 {
		if strings.Contains(s[secondLastDash+1:], ".") {
			return s[:lastDash], s[lastDash+1:]
		} else {
			return s, ""
		}
	}

	if strings.Contains(s[secondLastDash+1:], ".") {
		return s[:secondLastDash], s[secondLastDash+1:]
	} else {
		return s, ""
	}
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
