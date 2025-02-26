package blueprint

import (
	"errors"
	"fmt"
)

// FSNodeState type, one of: file, dir
type FSNodeState string

const (
	FSNodeStatePresent FSNodeState = "present"
	FSNodeStateAbsent  FSNodeState = "absent"
)

// ErrInvalidFSNodeState is returned when enum value is invalid
var ErrInvalidFSNodeState = errors.New("unexpected file system state")

func (np *FSNodeState) UnmarshalJSON(data []byte) error {
	if data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("%w: %q", ErrInvalidFSNodeState, string(data))
	}
	switch string(data[1 : len(data)-1]) {
	case "present", "absent":
		*np = FSNodeState(data)
		return nil
	default:
		return fmt.Errorf("%w: %q", ErrInvalidFSNodeState, data)
	}
}

func (np *FSNodeState) MarshalJSON() ([]byte, error) {
	if np == nil {
		return []byte("null"), nil
	}

	return []byte(*np), nil
}

func (np *FSNodeState) String() string {
	return string(*np)
}

func (np FSNodeState) IsPresent() bool {
	return np == "present"
}

func (np FSNodeState) IsAbsent() bool {
	return np == "absent"
}
