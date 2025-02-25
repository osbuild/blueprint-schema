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

	// From RHEL 8.7 & RHEL 9.1 support has been added for OpenSCAP build-time remediation
	OpenSCAP *OpenSCAP `json:"openscap,omitempty" jsonschema:"nullable"`
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
		Port uint16 `json:"port,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_port"`

		// From in a port range. Must be used with To and must be less than or equal to To.
		//
		// Port ranges are not subject to validation and may cause conflicts with services defined
		// via IANA service names or ports.
		From uint16 `json:"from,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_from_to"`

		// To in a port range. Must be used with From and must be greater than or equal to From.
		To uint16 `json:"to,omitempty" jsonschema:"minimum=1,maximum=65535,oneof_required=firewall_from_to"`

		// Protocol (tcp, udp, any) - all lowercase.
		Protocol NetworkProtocol `json:"protocol,omitempty" jsonschema:"default=any,enum=tcp,enum=udp,enum=any"`

		// Enable (default) or disable the service. If a port or service is disabled and enabled at
		// the same time either using services or ports fields, the service will be disabled.
		Enabled bool `json:"enabled,omitempty" jsonschema:"default=true"`
	} `json:"ports,omitempty" jsonschema:"nullable"`
}

type OpenSCAP struct {
	// The desired securinty profile ID.
	ProfileID string `json:"profile_id,omitempty" jsonschema:"required"`

	// Datastream to use for the scan. The datastream is the path to the SCAP datastream file to use for the scan.
	// If the datastream parameter is not provided, a sensible default based on the selected distro will be chosen.
	Datastream string `json:"datastream,omitempty"`

	// An optional OpenSCAP tailoring information. Can be done via profile selection or tailoring JSON file.
	//
	// In case of profile selection, a tailoring file with a new tailoring profile ID is created and saved to the image.
	// The new tailoring profile ID is created by appending the _osbuild_tailoring suffix to the base profile.
	// For example, given tailoring options for the cis profile, tailoring profile
	// xccdf_org.ssgproject.content_profile_cis_osbuild_tailoring will be created. The default namespace of the rules
	// is org.ssgproject.content, so the prefix may be omitted for rules under this namespace, i.e.
	// org.ssgproject.content_grub2_password and grub2_password are functionally equivalent.
	// The generated tailoring file is saved to the image as /usr/share/xml/osbuild-oscap-tailoring/tailoring.xml or,
	// for newer releases, in the /oscap_data directory, this is the location used for other OpenSCAP related artifacts.
	//
	// It is also possible to use JSON tailoring. In that case, custom JSON file must be provided using the blueprint and
	// used in json_filepath field alongside with json_profile_id field. The generated XML tailoring file is saved to the
	// image as /oscap_data/tailoring.xml.
	Tailoring *OpenSCAPTailoring `json:"tailoring,omitempty" jsonschema:"nullable"`
}

type OpenSCAPTailoring struct {
	// Selected profiles, cannot be used with json_profile_id and json_filepath.
	Selected []string `json:"selected,omitempty" jsonschema:"nullable,oneof_required=tailoring_selection"`

	// Unselected profiles, cannot be used with json_profile_id and json_filepath.
	Unselected []string `json:"unselected,omitempty" jsonschema:"nullable,oneof_required=tailoring_selection"`

	// JSON profile ID, must be used with json_filepath and cannot be used with selected and unselected fields.
	JSONProfileID string `json:"json_profile_id,omitempty" jsonschema:"oneof_required=tailoring_json"`

	// JSON filepath, must be used with json_profile_id and cannot be used with selected and unselected fields.
	JSONFilepath string `json:"json_filepath,omitempty" jsonschema:"oneof_required=tailoring_json"`
}
