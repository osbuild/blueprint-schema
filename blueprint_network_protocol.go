package blueprint

import (
	"errors"
	"fmt"
)

// NetworkProtocol type, one of: tcp, udp, any
type NetworkProtocol string

// ErrInvalidNetworkProtocol is returned when the NetworkProtocol is invalid
var ErrInvalidNetworkProtocol = errors.New("invalid network protocol")

func (np *NetworkProtocol) UnmarshalJSON(data []byte) error {
	if len(data) < 5 {
		return fmt.Errorf("%w: %q", ErrInvalidNetworkProtocol, string(data))
	}
	switch string(data[1 : len(data)-1]) {
	case "tcp", "udp", "any":
		*np = NetworkProtocol(data)
		return nil
	default:
		return fmt.Errorf("%w: %q", ErrInvalidNetworkProtocol, data)
	}
}

func (np *NetworkProtocol) MarshalJSON() ([]byte, error) {
	if np == nil {
		return []byte("null"), nil
	}

	return []byte(*np), nil
}

func (np *NetworkProtocol) String() string {
	return string(*np)
}

func (np NetworkProtocol) IsAny() bool {
	return np == "any"
}

func (np NetworkProtocol) IsTCP() bool {
	return np == "tcp"
}

func (np NetworkProtocol) IsUDP() bool {
	return np == "udp"
}
