// Package blueprint contains primitives for representing weldr blueprints
package blueprint

// A Blueprint is a high-level description of an image.
type Blueprint struct {
	Name           string          `json:"name" toml:"name"`
	Description    string          `json:"description" toml:"description"`
	Version        string          `json:"version,omitempty" toml:"version,omitempty"`
	Packages       []Package       `json:"packages" toml:"packages"`
	Modules        []Package       `json:"modules" toml:"modules"`
	Groups         []Group         `json:"groups" toml:"groups"`
	Containers     []Container     `json:"containers,omitempty" toml:"containers,omitempty"`
	Customizations *Customizations `json:"customizations,omitempty" toml:"customizations"`
	Distro         string          `json:"distro" toml:"distro"`

	// EXPERIMENTAL
	Minimal bool `json:"minimal" toml:"minimal"`
}

// A Package specifies an RPM package.
type Package struct {
	Name    string `json:"name" toml:"name"`
	Version string `json:"version,omitempty" toml:"version,omitempty"`
}

// A group specifies an package group.
type Group struct {
	Name string `json:"name" toml:"name"`
}

type Container struct {
	Source string `json:"source" toml:"source"`
	Name   string `json:"name,omitempty" toml:"name,omitempty"`

	TLSVerify    *bool `json:"tls-verify,omitempty" toml:"tls-verify,omitempty"`
	LocalStorage bool  `json:"local-storage,omitempty" toml:"local-storage,omitempty"`
}
