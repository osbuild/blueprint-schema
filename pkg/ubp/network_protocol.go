package ubp

import (
	"encoding/json"
	"fmt"
)

// ProtocolAny represents any network protocol.
const ProtocolAny NetworkProtocol = ""

func (np NetworkProtocol) String() string {
	return string(np)
}

// UnmarshalJSON handles default values.
func (np *NetworkProtocol) UnmarshalJSON(data []byte) error {
	var proto string
	if err := json.Unmarshal(data, &proto); err != nil {
		return fmt.Errorf("unmarshalling network protocol: %w", err)
	}

	parsedProto, err := ParseNetworkProtocol(proto)
	if err != nil {
		return fmt.Errorf("parsing network protocol %q: %w", proto, err)
	}

	*np = parsedProto
	return nil
}

// MarshalJSON handles default values.
func (np NetworkProtocol) MarshalJSON() ([]byte, error) {
	if np == ProtocolAny {
		return json.Marshal("")
	}

	return json.Marshal(string(np))
}
