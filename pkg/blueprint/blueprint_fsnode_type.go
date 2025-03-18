package blueprint

import (
	"errors"
	"fmt"
)

// FSNodeType type, one of: file, dir
type FSNodeType string

// ErrInvalidFSNodeType is returned when the FSNodeType is invalid
var ErrInvalidFSNodeType = errors.New("invalid file system node type")

func (np *FSNodeType) UnmarshalJSON(data []byte) error {
	if data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("%w: %q", ErrInvalidFSNodeType, string(data))
	}
	switch string(data[1 : len(data)-1]) {
	case "file", "dir":
		*np = FSNodeType(data)
		return nil
	default:
		return fmt.Errorf("%w: %q", ErrInvalidFSNodeType, data)
	}
}

func (np *FSNodeType) MarshalJSON() ([]byte, error) {
	if np == nil {
		return []byte("null"), nil
	}

	return []byte(*np), nil
}

func (np *FSNodeType) String() string {
	return string(*np)
}

func (np FSNodeType) IsDir() bool {
	return np == "dir"
}

func (np FSNodeType) IsFile() bool {
	return np == "file"
}
