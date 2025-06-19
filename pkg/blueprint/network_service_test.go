package blueprint

import "testing"

func TestNetworkProtocolJSON(t *testing.T) {
	tests := []struct {
		input string
		proto    NetworkProtocol
		output string
	}{
		{`""`, ProtocolAny, `""`},
		{`"any"`, ProtocolAny, `""`},
		{`"tcp"`, ProtocolTCP, `"tcp"`},
		{`"udp"`, ProtocolUDP, `"udp"`},
		{`"icmp"`, ProtocolICMP, `"icmp"`},
	}

	for _, test := range tests {
		t.Run(test.proto.String(), func(t *testing.T) {
			var proto NetworkProtocol
			err := proto.UnmarshalJSON([]byte(test.input))
			if err != nil {
				t.Fatalf("unexpected error unmarshaling protocol %q: %v", test.input, err)
			}
			if proto != test.proto {
				t.Errorf("expected %s, got %s", test.proto, proto)
			}

			data, err := test.proto.MarshalJSON()
			if err != nil {
				t.Fatalf("unexpected error marshaling protocol %q: %v", test.proto, err)
			}
			if string(data) != test.output {
				t.Errorf("expected %s, got %s", test.output, data)
			}
		})
	}
}

func TestParseFirewalldPort(t *testing.T) {
	tests := []struct {
		input    string
		expected FirewallPort
	}{
		{"80", FirewallPort{Port: 80, Protocol: ProtocolAny}},
		{"80:tcp", FirewallPort{Port: 80, Protocol: ProtocolTCP}},
		{"80:udp", FirewallPort{Port: 80, Protocol: ProtocolUDP}},
		{"443:tcp", FirewallPort{Port: 443, Protocol: ProtocolTCP}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			port, err := ParseFirewalldPort(test.input)
			if err != nil {
				t.Fatalf("unexpected error parsing port %q: %v", test.input, err)
			}
			if port != test.expected {
				t.Errorf("expected %v, got %v", test.expected, port)
			}
		})
	}
}

func TestParseFromTo(t *testing.T) {
	tests := []struct {
		input    string
		expected FirewallFromTo
	}{
		{"1000-2000", FirewallFromTo{From: 1000, To: 2000, Protocol: ProtocolAny}},
		{"5000-6000:tcp", FirewallFromTo{From: 5000, To: 6000, Protocol: ProtocolTCP}},
		{"30000-32767:tcp", FirewallFromTo{From: 30000, To: 32767, Protocol: ProtocolTCP}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			fromTo, err := ParseFirewalldFromTo(test.input)
			if err != nil {
				t.Fatalf("unexpected error parsing from-to %q: %v", test.input, err)
			}
			if fromTo != test.expected {
				t.Errorf("expected %v, got %v", test.expected, fromTo)
			}
		})
	}
}
