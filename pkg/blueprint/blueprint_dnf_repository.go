package blueprint

type DNFRepository struct {
	// Repository ID. Required.
	ID string `json:"id" jsonschema:"required,maxLength=256,pattern=^[a-zA-Z0-9_-]+$"`

	// Repository name.
	Name string `json:"name,omitempty"`

	// Repository filename to use for the repository configuration file. If not provided, the ID is used.
	// Filename must be provided without the .repo extension.
	Filename string `json:"filename,omitempty" jsonschema:"maxLength=256,pattern=^[a-zA-Z0-9_-]+$"`

	// Base URLs for the repository.
	BaseURLs []string `json:"base_urls,omitempty" jsonschema:"oneof_required=dnf_repo_base_urls"`

	// Mirror list for the repository.
	MirrorList string `json:"mirror_list,omitempty" jsonschema:"oneof_required=dnf_repo_mirrorlist"`

	// Metalink for the repository.
	Metalink string `json:"metalink,omitempty" jsonschema:"oneof_required=dnf_repo_metalink"`

	// GPG keys for the repository.
	//
	// The blueprint accepts both inline GPG keys and GPG key urls. If an inline GPG key is provided
	// it will be saved to the /etc/pki/rpm-gpg directory and will be referenced accordingly in the
	// repository configuration. GPG keys are not imported to the RPM database and will only be imported
	// when first installing a package from the third-party repository.
	GPGKeys []string `json:"gpg_keys,omitempty"`

	// Enable GPG check for the repository.
	GPGCheck bool `json:"gpg_check,omitempty" jsonschema:"default=true"`

	// Enable GPG check for the repository metadata.
	GPGCheckRepo bool `json:"gpg_check_repo,omitempty" jsonschema:"default=true"`

	// Repository priority.
	Priority int `json:"priority,omitempty" jsonschema:"default=99"`

	// Enable SSL verification for the repository.
	SSLVerify bool `json:"ssl_verify,omitempty" jsonschema:"default=true"`

	// Enable module hotfixes for the repository.
	//
	// Adds module_hotfixes flag to all repo types so it can be used during osbuild.
	// This enables users to disable modularity filtering on specific repositories.
	ModuleHotfixes bool `json:"module_hotfixes,omitempty" jsonschema:"default=false"`

	// Repository usage. By default, the repository is configured on the image but not used for image build.
	Usage *DNFRepositoryUsage `json:"usage,omitempty" jsonschema:"nullable"`
}

type DNFRepositoryUsage struct {
	// Use the repository for image build.
	//
	// When this flag is set, it is possible to install third-party packages during the image build.
	Install bool `json:"install,omitempty" jsonschema:"default=false"`

	// Configure the repository for dnf.
	//
	// A repository will be saved to the /etc/yum.repos.d directory in an image.
	// An optional filename argument can be set, otherwise the repository will be saved using the the repository
	// ID, i.e. /etc/yum.repos.d/<repo-id>.repo.
	Configure bool `json:"configure,omitempty" jsonschema:"default=true"`
}
