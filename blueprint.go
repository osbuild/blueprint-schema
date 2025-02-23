package blueprint

import (
	"errors"
	"fmt"
)

// Blueprint type prototype
//
// This is just a brief example of a common blueprint structure. Just few fields
// were selected to demonstrate the JSON schema.
//
// These Go comments do appear in the JSON Schema so the final version of the
// blueprint will be broken up into multiple files and the comments will be
// moch more extensive.
//
// TODO: Break all anonymous struct into named structs. Break this .go file into
// multiple files.
type Blueprint struct {
	// Name of the blueprint
	Name string `json:"name" jsonschema:"required"`

	// Description of the blueprint
	Description string `json:"description,omitempty"`

	// Registration details
	Registration *Registration `json:"registration,omitempty" jsonschema:"nullable"`

	// Networking details
	Network *Network `json:"network,omitempty" jsonschema:"nullable"`

	// OS hostname
	Hostname string `json:"hostname,omitempty"`

	// FIPS details
	FIPS *FIPS `json:"fips,omitempty" jsonschema:"nullable"`

	// DNF package managers details
	DNF *DNF `json:"dnf,omitempty" jsonschema:"nullable"`

	// Containers details
	Containers []Containers `json:"containers,omitempty" jsonschema:"nullable"`

	// Linux kernel details
	Kernel *Kernel `json:"kernel,omitempty" jsonschema:"nullable"`
}

type FIPS struct {
	// Enable FIPS mode
	Enabled bool `json:"enabled,omitempty"`
}

type DNF struct {
	// Packages to install. Package name or NVRA is accepted as long as DNF can
	// resolve it.
	Packages []string `json:"packages,omitempty" jsonschema:"nullable"`

	// Groups to install. Groups can also be specificed via the packages field
	// prefixed with @.
	Groups []string `json:"groups,omitempty" jsonschema:"nullable"`

	// Import GPG keys, default is true.
	ImportKeys bool `json:"import_keys,omitempty" jsonschema:"default=true"`

	// Modules to enable or disable
	Modules []string `json:"modules,omitempty" jsonschema:"nullable"`
}

type Containers struct {
	Source       string `json:"source" jsonschema:"required"`
	Name         string `json:"name" jsonschema:"required"`
	TlsVerify    bool   `json:"tls_verify,omitempty"`
	LocalStorage string `json:"local_storage,omitempty"`
}

type Kernel struct {
	Package       string   `json:"package,omitempty"`
	CmdlineAppend []string `json:"cmdline_append,omitempty"`
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
	SubscriptionManager *SubscriptionManagerRegistration `json:"subscription_manager" jsonschema:"nullable"`

	// Insights client details
	Insights struct {
		// Enables insights client during boot
		Enabled bool `json:"enabled"`
	} `json:"insights" jsonschema:"nullable"`

	// CRC connector details
	Connector struct {
		// Enables CRC connector during boot
		Enabled bool `json:"enabled"`
	} `json:"connector" jsonschema:"nullable"`

	// FDO details
	FDO struct {
		ManufacturingServerURL  string `json:"manufacturing_server_url,omitempty"`
		DiunPubKeyInsecure      bool   `json:"diun_pub_key_insecure,omitempty"`
		DiunPubKeyHash          string `json:"diun_pub_key_hash,omitempty"`
		DiunPubKeyRootCerts     string `json:"diun_pub_key_root_certs,omitempty"`
		DiMfgStringTypeMacIface string `json:"di_mfg_string_type_mac_iface,omitempty"`
	} `json:"fdo,omitempty" jsonschema:"nullable"`
}

type SubscriptionManagerRegistration struct {
	Enabled              bool `json:"enabled"`
	ProductPlugin        bool `json:"product_plugin"`
	RepositoryManagement bool `json:"repository_management"`
	AutoRegistration     bool `json:"auto_registration"`
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
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

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
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"ports,omitempty" jsonschema:"nullable"`
}

type NetworkProtocol string

var ErrInvalidNetworkProtocol = errors.New("invalid network protocol")

func (np *NetworkProtocol) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "tcp", "udp", "any":
		*np = NetworkProtocol(data)
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrInvalidNetworkProtocol, data)
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
