package blueprint

// Blueprint type prototype
//
// This is just a brief example of a common blueprint structure. Just few fields
// were selected to demonstrate the JSON schema.
//
// These Go comments do appear in the JSON Schema so the final version of the
// blueprint will be broken up into multiple files and the comments will be
// moch more extensive.
type Blueprint struct {
	// Name of the blueprint
	Name string `json:"name" jsonschema:"required"`

	// Description of the blueprint
	Description string `json:"description,omitempty"`

	// Registration details
	Registration *Registration `json:"registration,omitempty" jsonschema:"nullable"`

	// Networking details
	Network *Network `json:"network,omitempty" jsonschema:"nullable"`
}

type Registration struct {
	// RedHat registration details
	RedHat *RedHatRegistration `json:"redhat,omitempty"`
}

type RedHatRegistration struct {
	// Activation key
	ActivationKey string `json:"activation_key"`

	// Organization ID
	Organization string `json:"organization"`

	// Subscription manager details (internal use only)
	SubscriptionManager struct {
		Enabled              bool `json:"enabled"`
		ProductPlugin        bool `json:"product_plugin"`
		RepositoryManagement bool `json:"repository_management"`
		AutoRegistration     bool `json:"auto_registration"`
	} `json:"subscription_manager" onlyFor:"internal"`

	// Insights client details
	Insights struct {
		// Enables insights client during boot
		Enabled bool `json:"enabled"`
	} `json:"insights" onlyFor:"crc"`

	// CRC connector details
	Connector struct {
		// Enables CRC connector during boot
		Enabled bool `json:"enabled"`
	} `json:"connector" onlyFor:"crc"`
}

type Network struct {
	// Firewall details
	Firewall *NetworkFirewall `json:"firewall,omitempty"`
}

type NetworkFirewall struct {
	// Services to enable or disable
	Services []struct {
		// Service name from the IANA list
		Service string `json:"service" jsonschema:"required"`

		// Protocol (tcp, udp, any)
		Protocol string `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"services,omitempty" jsonschema:"nullable"`

	// Ports or ranges to enable or disable
	Ports []struct {
		// Service port number (or use From/To)
		Port uint16 `json:"port,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=port"`

		// From range (or use Port)
		From uint16 `json:"from,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=from_to"`

		// To range (or use Port)
		To uint16 `json:"to,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=from_to"`

		// Protocol (tcp, udp, any)
		Protocol string `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"ports,omitempty" jsonschema:"nullable"`
}
