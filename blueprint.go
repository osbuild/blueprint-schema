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

	// Registration details
	Registration *Registration `json:"registration,omitempty" jsonschema:"nullable"`

	// Users and groups details
	Accounts *Accounts `json:"accounts,omitempty" jsonschema:"nullable"`

	// Time and date details
	TimeDate *TimeDate `json:"timedate,omitempty" jsonschema:"nullable"`

	Locale *Locale `json:"locale,omitempty" jsonschema:"nullable"`
	
	// Networking details
	Network *Network `json:"network,omitempty" jsonschema:"nullable"`

	// From RHEL 8.7 & RHEL 9.1 support has been added for OpenSCAP build-time remediation
	OpenSCAP *OpenSCAP `json:"openscap,omitempty" jsonschema:"nullable"`
}
