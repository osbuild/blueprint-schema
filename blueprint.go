package blueprint

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
	// The name attribute is a string that contains the name of the blueprint. It can contain spaces,
	// but they may be converted to dash characters during build. It should be short and descriptive.
	Name string `json:"name" jsonschema:"required"`

	// The description attribute is a string that can be a longer description of the blueprint and is
	// only used for display purposes.
	Description string `json:"description,omitempty"`

	// Hostname is an optional string that can be used to configure the hostname of the final image.
	Hostname string `json:"hostname,omitempty"`

	// Custom Linux kernel details, optional.
	Kernel *Kernel `json:"kernel,omitempty" jsonschema:"nullable"`

	// FIPS details, optional.
	FIPS *FIPS `json:"fips,omitempty" jsonschema:"nullable"`

	// DNF package managers details. When using virtual provides as the package name the version glob
	// should be *. And be aware that you will be unable to freeze the blueprint. This is because the
	// provides will expand into multiple packages with their own names and versions.
	DNF *DNF `json:"dnf,omitempty" jsonschema:"nullable"`

	// Containers to be pulled during the image build and stored in the image at the default local
	// container storage location that is appropriate for the image type, so that all supported container
	// tools like podman and cri-o will be able to work with it.
	// The embedded containers are not started, to do so you can create systemd unit files or quadlets with
	// the files customization.
	Containers []Containers `json:"containers,omitempty" jsonschema:"nullable"`

	// Registration details
	Registration *Registration `json:"registration,omitempty" jsonschema:"nullable"`

	// Networking details
	Network *Network `json:"network,omitempty" jsonschema:"nullable"`
}

type FIPS struct {
	// Enables the system FIPS mode (disabled by default). Currently only edge-raw-image, edge-installer,
	// edge-simplified-installer, edge-ami and edge-vsphere images support this customization.
	Enabled bool `json:"enabled,omitempty"`
}

type DNF struct {
	// Packages to install. Package name or NVRA is accepted as long as DNF can
	// resolve it. Examples: vim-enhanced, vim-enhanced-9.1.866-1 or vim-enhanced-9.1.866-1.fc41.x86_64.
	// The packages can also be specified as @group_name to install all packages in the group.
	Packages []string `json:"packages,omitempty" jsonschema:"nullable"`

	// Groups to install, must match exactly. Groups describes groups of packages to be installed into
	// the image. Package groups are defined in the repository metadata. Each group has a descriptive name
	// used primarily for display in user interfaces and an ID more commonly used in kickstart files.
	// Here, the ID is the expected way of listing a group. Groups have three different ways of categorizing
	// their packages: mandatory, default, and optional. For the purposes of blueprints, only mandatory
	// and default packages will be installed. There is no mechanism for selecting optional packages.
	Groups []string `json:"groups,omitempty" jsonschema:"nullable"`

	// Additional file paths to the GPG keys to import. The files must be present in the image.
	// Does not support importing from URLs.
	ImportKeys []string `json:"import_keys,omitempty" jsonschema:"nullable"`

	// Modules to enable or disable
	Modules []string `json:"modules,omitempty" jsonschema:"nullable"`
}

type Containers struct {
	// Container image URL is a reference to a container image at a registry.
	Source string `json:"source" jsonschema:"required"`

	// Container name is an optional string to set the name under which the container image will
	// be saved in the image. If not specified name falls back to the same value as source.
	Name string `json:"name" jsonschema:"required"`

	// Verify TLS connection, default is true.
	TLSVerify bool `json:"tls_verify,omitempty" jsonschema:"default=true"`

	// Whether to pull the container image from the host's local-storage.
	LocalStorage string `json:"local_storage,omitempty"`
}

type Kernel struct {
	// Kernel DNF package name to replace the standard kernel with.
	Package string `json:"package,omitempty"`

	// An optional string to append arguments to the bootloader kernel command line. The list
	// will be concatenated with spaces.
	CmdlineAppend []string `json:"cmdline_append,omitempty" jsonschema:"nullable"`
}

type Registration struct {
	// Registration details for Red Hat operating system images.
	RedHat *RedHatRegistration `json:"redhat,omitempty"`
}

type RedHatRegistration struct {
	// Subscription manager activation key to use during registration.
	ActivationKey string `json:"activation_key"`

	// Subscription manager organization ID to use during registration.
	Organization string `json:"organization"`

	// Subscription manager details (internal use only). The customization expects that subscription-manager
	// package is installed in the image, which is by default part of the RHEL distribution bootable images.
	// To explicitly install the package, add it to the packages section in the blueprint.
	// The customization is not supported on Fedora distribution images.
	SubscriptionManager *SubscriptionManagerRegistration `json:"subscription_manager" jsonschema:"nullable"`

	// Red Hat Insights client details.
	Insights struct {
		// Enables insights client during boot.
		Enabled bool `json:"enabled"`
	} `json:"insights" jsonschema:"nullable"`

	// Red Hat console.redhat.com connector (rhc) details.
	Connector struct {
		// Enables rhc (Red Hat Connector) during boot.
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
	// Enables the subscription-manager DNF plugin.
	Enabled bool `json:"enabled" jsonschema:"default=true"`

	// Enables the product-id DNF plugin.
	ProductPluginEnabled bool `json:"product_plugin_enabled" jsonschema:"default=true"`

	// Enabled repository_management plugin configuration.
	RepositoryManagement bool `json:"repository_management" jsonschema:"default=true"`

	// Enabled auto_registration plugin configuration.
	AutoRegistration bool `json:"auto_registration" jsonschema:"default=true"`
}

type Network struct {
	// Firewall details - package firewalld must be installed in the image.
	Firewall *NetworkFirewall `json:"firewall,omitempty"`
}

type NetworkFirewall struct {
	// Services to enable or disable. The service name must be from the IANA list.
	// Alternatively, you can specify a port or range using the ports field.
	Services []struct {
		// Service name from the IANA list.
		//
		// A service must appear only single time in both services and ports fields to prevent
		// conflicts.
		Service string `json:"service" jsonschema:"required"`

		// Protocol (tcp, udp, any) - all lowercase.
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service. If a port or service is disabled and enabled at
		// the same time either using services or ports fields, the service will be disabled.
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"services,omitempty" jsonschema:"nullable"`

	// Ports or ranges to enable or disable
	Ports []struct {
		// Service port number. Alternatively, a port range via from/to fields can be used.
		//
		// A service must appear only single time in both services and ports fields to prevent
		// conflicts.
		Port uint16 `json:"port,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=port"`

		// From in a port range. Must be used with To and must be less than or equal to To.
		//
		// Port ranges are not subject to validation and may cause conflicts with services defined
		// via IANA service names or ports.
		From uint16 `json:"from,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=from_to"`

		// To in a port range. Must be used with From and must be greater than or equal to From.
		To uint16 `json:"to,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=from_to"`

		// Protocol (tcp, udp, any) - all lowercase.
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service. If a port or service is disabled and enabled at
		// the same time either using services or ports fields, the service will be disabled.
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"ports,omitempty" jsonschema:"nullable"`
}
