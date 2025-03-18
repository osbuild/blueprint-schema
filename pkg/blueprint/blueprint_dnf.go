package blueprint

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

	// Third-party repositories are supported by the blueprint customizations.
	//
	// All fields reflect configuration values of dnf, see man dnf.conf(5) for more information.
	Repositories []DNFRepository `json:"repositories,omitempty" jsonschema:"nullable"`
}
