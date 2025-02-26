package blueprint

// Image Builder new blueprint schema.
//
// THIS IS WORK IN PROGRESS
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

	// Registration details for various registration types, namely Red Hat Subscription Manager.
	Registration *Registration `json:"registration,omitempty" jsonschema:"nullable"`

	// Users and groups details
	Accounts *Accounts `json:"accounts,omitempty" jsonschema:"nullable"`

	// Time and date details allowing configuration of the timezone and NTP servers. The timezone is
	// set by default to UTC.
	TimeDate *TimeDate `json:"timedate,omitempty" jsonschema:"nullable"`

	// An optional object that contains the following attributes to customize the locale settings for the system.
	// If the locale is not specified, the default locale and keyboard settings are used: en_US.UTF-8 and us.
	Locale *Locale `json:"locale,omitempty" jsonschema:"nullable"`

	// Networking details including firewall configuration.
	Network *Network `json:"network,omitempty" jsonschema:"nullable"`

	// OpenSCAP policy to be applied on the operating system. Added in RHEL 8.7 & RHEL 9.1. It is possible to either
	// list policy rules (enable or disable) or to provide a full policy file.
	OpenSCAP *OpenSCAP `json:"openscap,omitempty" jsonschema:"nullable"`

	// Systemd unit configuration.
	//
	// This section can be used to control which services are enabled at boot time. Some image types already
	// have services enabled or disabled in order for the image to work correctly, and cannot be overridden.
	// For example, ami image type requires sshd, chronyd, and cloud-init services. Blueprint services do not
	// replace these services, but add them to the list of services already present in the templates, if any.
	//
	Systemd *Systemd `json:"systemd,omitempty" jsonschema:"nullable"`

	// File system nodes details.
	//
	// You can use the customization to create new files or to replace existing ones, if not restricted by
	// the policy specified below. If the target path is an existing symlink to another file, the symlink
	// will be replaced by the custom file.
	//
	// Please note that the parent directory of a specified file must exist. If it does not exist, the image
	// build will fail. One can ensure that the parent directory exists by specifying "ensure_parents".
	//
	// In addition, the following files are not allowed to be created or replaced by policy: /etc/fstab,
	// /etc/shadow, /etc/passwd and /etc/group.
	//
	// Using the files customization comes with a high chance of creating an image that doesn't boot. Use this
	// feature only if you know what you are doing. Although the files customization can be used to configure
	// parts of the OS which can also be configured by other blueprint customizations, this use is discouraged.
	// If possible, users should always default to using the specialized blueprint customizations. Note that
	// if you combine the files customizations with other customizations, the other customizations may not work
	// as expected or may be overridden by the files customizations.
	//
	// You can create custom directories as well. The existence of a specified directory is handled gracefully
	// only if no explicit mode, user or group is specified. If any of these customizations are specified and
	// the directory already exists in the image, the image build will fail. The intention is to prevent changing
	// the ownership or permissions of existing directories.
	FSNodes []FSNode `json:"fsnodes,omitempty" jsonschema:"nullable"`

	// Provides Ignition configuration files to be used in edge-raw-image and edge-simplified-installer images.
	// Check the RHEL for Edge butane specification for a description of the supported configuration options.
	//
	// The blueprint configuration can be done either by embedding an Ignition configuration file into the image,
	// or providing a provisioning URL that will be fetched at first boot.
	Ignition *Ignition `json:"ignition,omitempty" jsonschema:"nullable"`

	// Extra customization for Anaconda installer (ISO) and Edge/IOT simplified installer image types.
	Installer *Installer `json:"installer,omitempty" jsonschema:"nullable"`
}
