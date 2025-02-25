package blueprint

type Registration struct {
	// Registration details for Red Hat operating system images.
	RedHat *RedHatRegistration `json:"redhat,omitempty"`

	// FDO allows users to configure FIDO Device Onboard device initialization parameters. It is only available
	// with the edge-simplified-installer or iot-simplified-installer image types.
	FDO *FDORegistration `json:"fdo,omitempty" jsonschema:"nullable"`
}

type RedHatRegistration struct {
	// Subscription manager activation key to use during registration. A list of keys to use to redeem or apply
	// specific subscriptions to the system.
	ActivationKey string `json:"activation_key,omitempty"`

	// Subscription manager organization name to use during registration.
	Organization string `json:"organization,omitempty"`

	// Subscription manager details (internal use only). The customization expects that subscription-manager
	// package is installed in the image, which is by default part of the RHEL distribution bootable images.
	// To explicitly install the package, add it to the packages section in the blueprint.
	// The customization is not supported on Fedora distribution images.
	SubscriptionManager *SubscriptionManagerRegistration `json:"subscription_manager,omitempty" jsonschema:"nullable"`

	// Red Hat Insights client details.
	Insights *InsightsRegistration `json:"insights,omitempty" jsonschema:"nullable"`

	// Red Hat console.redhat.com connector (rhc) details.
	Connector *ConnectorRegistration `json:"connector,omitempty" jsonschema:"nullable"`
}

type InsightsRegistration struct {
	// Enables insights client during boot.
	Enabled bool `json:"enabled"`
}

type ConnectorRegistration struct {
	// Enables rhc (Red Hat Connector) during boot.
	Enabled bool `json:"enabled"`
}

type FDORegistration struct {
	// FDO manufacturing server URL.
	ManufacturingServerURL string `json:"manufacturing_server_url,omitempty" jsonschema:"required"`

	// FDO insecure option. When set, both hash or root certs must not be set.
	DiunPubKeyInsecure bool `json:"diun_pub_key_insecure,omitempty" jsonschema:"oneof_required=fdo_insecure"`

	// FDO server public key hex-encoded hash. Cannot be used together with insecure option or root certs.
	DiunPubKeyHash string `json:"diun_pub_key_hash,omitempty" jsonschema:"oneof_required=fdo_hash"`

	// FDO server public key root certificate path. Cannot be used together with insecure option or hash.
	DiunPubKeyRootCerts string `json:"diun_pub_key_root_certs,omitempty" jsonschema:"oneof_required=fdo_rootcerts"`

	// Optional interface name for the MAC address.
	DiMfgStringTypeMacIface string `json:"di_mfg_string_type_mac_iface,omitempty"`
}

type SubscriptionManagerRegistration struct {
	// Enables the subscription-manager DNF plugin.
	Enabled bool `json:"enabled" jsonschema:"default=true"`

	// Enables the product-id DNF plugin.
	ProductPluginEnabled bool `json:"product_plugin_enabled" jsonschema:"default=true"`

	// Enabled repository_management plugin configuration.
	RepositoryManagement bool `json:"repository_management" jsonschema:"default=true"`

	// Enabled auto_registration plugin configuration.
	AutoRegistration bool `json:"auto_registration" jsonschema:"default=true"`
}
