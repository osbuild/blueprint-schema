package blueprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func (t NetworkService) SelectUnion() (FirewallService, FirewallPort, FirewallFromTo, error) {
	var fs FirewallService
	err := json.Unmarshal(t.union, &fs)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	var fp FirewallPort
	err = json.Unmarshal(t.union, &fp)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	var fft FirewallFromTo
	err = json.Unmarshal(t.union, &fft)
	if err != nil {
		return FirewallService{}, FirewallPort{}, FirewallFromTo{}, err
	}

	return fs, fp, fft, nil
}

var ErrInvalidNetworkProtocol = errors.New("invalid network protocol")

func ParseNetworkProtocol(s string) (NetworkProtocol, error) {
	switch strings.ToLower(s) {
	case "", "any":
		return ProtocolAny, nil
	case "tcp":
		return ProtocolTCP, nil
	case "udp":
		return ProtocolUDP, nil
	case "icmp":
		return ProtocolICMP, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidNetworkProtocol, s)
	}
}

func NetworkServiceFromService(node FirewallService) *NetworkService {
	u, _ := json.Marshal(node)
	return &NetworkService{union: u}
}

func NetworkServiceFromPort(node FirewallPort) *NetworkService {
	u, _ := json.Marshal(node)
	return &NetworkService{union: u}
}

func NetworkServiceFromFromTo(node FirewallFromTo) *NetworkService {
	u, _ := json.Marshal(node)
	return &NetworkService{union: u}
}

var ErrFirewallParseError = errors.New("firewall parse error")

// ParseFirewalldPort parses port strings in the firewall-cmd format:
// 22:tcp imap:tcp 53:udp 30000-32767:tcp
func ParseFirewalldPort(port string) (FirewallPort, error) {
	var fp FirewallPort

	parts := strings.Split(port, ":")
	if len(parts) < 1 || len(parts) > 2 {
		return FirewallPort{}, fmt.Errorf("%w: expected format 'port[:protocol]', got '%s'", ErrFirewallParseError, port)
	}

	iport, err := strconv.ParseUint(parts[0], 10, 16)
	if err != nil {
		return FirewallPort{}, fmt.Errorf("%w: invalid port '%s': %v", ErrFirewallParseError, parts[0], err)
	}

	fp.Port = int(iport)
	if len(parts) == 2 {
		proto, err := ParseNetworkProtocol(parts[1])
		if err != nil {
			return FirewallPort{}, fmt.Errorf("%w: invalid protocol '%s': %v", ErrFirewallParseError, parts[1], err)
		}
		fp.Protocol = proto
	} else {
		fp.Protocol = ProtocolAny
	}

	return fp, nil
}

// ParseFirewalldPort parses port strings in the firewall-cmd format: FROM-TO[:PROTOCOL]
// It does not support single ports, only ranges.
func ParseFirewalldFromTo(port string) (FirewallFromTo, error) {

	parts := strings.Split(port, ":")
	if len(parts) < 1 || len(parts) > 2 {
		return FirewallFromTo{}, fmt.Errorf("%w: expected format 'range[:protocol]', got '%s'", ErrFirewallParseError, port)
	}

	rangeParts := strings.Split(parts[0], "-")
	if len(rangeParts) != 2 {
		return FirewallFromTo{}, fmt.Errorf("%w: expected format 'from-to[:protocol]', got '%s'", ErrFirewallParseError, port)
	}
	from, err := strconv.ParseUint(rangeParts[0], 10, 16)
	if err != nil {
		return FirewallFromTo{}, fmt.Errorf("%w: invalid from port '%s': %v", ErrFirewallParseError, rangeParts[0], err)
	}
	to, err := strconv.ParseUint(rangeParts[1], 10, 16)
	if err != nil {
		return FirewallFromTo{}, fmt.Errorf("%w: invalid to port '%s': %v", ErrFirewallParseError, rangeParts[1], err)
	}
	var proto NetworkProtocol
	if len(parts) == 2 {
		proto, err = ParseNetworkProtocol(parts[1])
		if err != nil {
			return FirewallFromTo{}, fmt.Errorf("%w: invalid protocol '%s': %v", ErrFirewallParseError, parts[1], err)
		}
	} else {
		proto = ProtocolAny
	}
	return FirewallFromTo{
		From:     int(from),
		To:       int(to),
		Protocol: proto,
	}, nil
}

// PortProtoToFirewalld converts a port and protocol to the firewalld format: PORT[:PROTOCOL]
func PortProtoToFirewalld(port int, proto NetworkProtocol) string {
	if proto == ProtocolAny || proto == "" {
		return strconv.Itoa(port)
	}

	return fmt.Sprintf("%d:%s", port, proto)
}

// PortsProtoToFirewalld converts a range of ports and protocol to the firewalld format: FROM-TO[:PROTOCOL]
func PortsProtoToFirewalld(from int, to int, proto NetworkProtocol) string {
	if proto == ProtocolAny || proto == "" {
		return fmt.Sprintf("%d-%d", from, to)
	}

	return fmt.Sprintf("%d-%d:%s", from, to, proto)
}
