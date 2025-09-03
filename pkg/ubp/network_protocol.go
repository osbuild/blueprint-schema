package ubp

// ProtocolAny represents any network protocol.
const ProtocolAny NetworkProtocol = ""

func (np NetworkProtocol) String() string {
	return string(np)
}
