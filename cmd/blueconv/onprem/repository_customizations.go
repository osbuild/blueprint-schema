package blueprint

type RepositoryCustomization struct {
	Id             string   `json:"id" toml:"id"`
	BaseURLs       []string `json:"baseurls,omitempty" toml:"baseurls,omitempty"`
	GPGKeys        []string `json:"gpgkeys,omitempty" toml:"gpgkeys,omitempty"`
	Metalink       string   `json:"metalink,omitempty" toml:"metalink,omitempty"`
	Mirrorlist     string   `json:"mirrorlist,omitempty" toml:"mirrorlist,omitempty"`
	Name           string   `json:"name,omitempty" toml:"name,omitempty"`
	Priority       *int     `json:"priority,omitempty" toml:"priority,omitempty"`
	Enabled        *bool    `json:"enabled,omitempty" toml:"enabled,omitempty"`
	GPGCheck       *bool    `json:"gpgcheck,omitempty" toml:"gpgcheck,omitempty"`
	RepoGPGCheck   *bool    `json:"repo_gpgcheck,omitempty" toml:"repo_gpgcheck,omitempty"`
	SSLVerify      *bool    `json:"sslverify,omitempty" toml:"sslverify,omitempty"`
	ModuleHotfixes *bool    `json:"module_hotfixes,omitempty" toml:"module_hotfixes,omitempty"`
	Filename       string   `json:"filename,omitempty" toml:"filename,omitempty"`
}
