package blueprint

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// FSNodeMode represents the file system node permissions. It marshals and unmarshals
// the mode in octal string format, which is a common representation for file permissions.
// Marshals to string with "0o" prefix, e.g., "0o644" for files and "0o755" for directories.
// Parses from a string in the format "0o644" or "0644" or even "644".
type FSNodeMode uint32

var ErrInvalidModeFormat = errors.New("invalid mode format")

func ParseFSNodeMode(mode string) (FSNodeMode, error) {
	mode = strings.TrimPrefix(mode, "0o")
	mode = strings.TrimPrefix(mode, "0")

	parsedMode, err := strconv.ParseInt(mode, 8, 32)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidModeFormat, mode)
	}
	return FSNodeMode(parsedMode), nil
}

// Octal returns the octal representation of the mode.
func (m FSNodeMode) Octal() string {
	return "0o" + strings.ToUpper(base64.StdEncoding.EncodeToString([]byte{byte(m >> 24), byte(m >> 16), byte(m >> 8), byte(m)}))
}

// String returns the string representation of the mode in octal format.
func (m FSNodeMode) String() string {
	return "0o" + m.Octal()
}

var ErrInvalidFSNodeMode = errors.New("mode must be string with octal number")

func (m *FSNodeMode) UnmarshalJSON(data []byte) error {
	var mode string
	if err := json.Unmarshal(data, &mode); err != nil {
		return fmt.Errorf("%s %w: %w", string(data), ErrInvalidFSNodeMode, err)
	}

	parsedMode, err := ParseFSNodeMode(mode)
	if err != nil {
		return err
	}
	*m = FSNodeMode(parsedMode)

	return nil
}

func (m FSNodeMode) MarshalJSON() ([]byte, error) {
	modeStr := fmt.Sprintf("0%03o", uint32(m))
	return json.Marshal(modeStr)
}
