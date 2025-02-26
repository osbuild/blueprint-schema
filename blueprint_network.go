package blueprint

type Network struct {
	// Firewall details - package firewalld must be installed in the image.
	Firewall *NetworkFirewall `json:"firewall,omitempty"`
}

type NetworkFirewall struct {
	// Services to enable or disable. The service can be defined via an assigned IANA name,
	// port number or port range.
	//
	// Services are processed in order, when a service is disabled and then accidentally enabled
	// via a port or a port range, the service will be enabled in the end.
	//
	// By default the firewall blocks all access, except for services that enable their ports
	// explicitly such as the sshd.
	Services []struct {
		// Service name from the IANA list. Examples: ssh, http, https, etc.
		//
		// This field is mutually exclusive with service, port and from/to pair.
		Service string `json:"service,omitempty" jsonschema:"minLength=2,oneof_required=firewall_service"`

		// Service port number. Alternatively, a port range via from/to fields can be used.
		//
		// This field is mutually exclusive with service, port and from/to pair.
		Port uint16 `json:"port,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_port"`

		// From in a port range. Must be used with "to" and must be less than or equal to "to".
		//
		// This field is mutually exclusive with service, port and from/to pair.
		From uint16 `json:"from,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_from_to"`

		// To in a port range. Must be used with "from" and must be greater than or equal to "from".
		//
		// This field is mutually exclusive with service, port and from/to pair.
		To uint16 `json:"to,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_from_to"`

		// Protocol (tcp, udp, any) - all lowercase.
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service. If a port or service is disabled and enabled at
		// the same time either using services or ports fields, the service will be disabled.
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"services,omitempty" jsonschema:"nullable"`
}
