package blueprint

// DirectoryCustomization represents a directory to be created in the image
type DirectoryCustomization struct {
	// Absolute path to the directory
	Path string `json:"path" toml:"path"`
	// Owner of the directory specified as a string (user name), int64 (UID) or nil
	User interface{} `json:"user,omitempty" toml:"user,omitempty"`
	// Owner of the directory specified as a string (group name), int64 (UID) or nil
	Group interface{} `json:"group,omitempty" toml:"group,omitempty"`
	// Permissions of the directory specified as an octal number
	Mode string `json:"mode,omitempty" toml:"mode,omitempty"`
	// EnsureParents ensures that all parent directories of the directory exist
	EnsureParents bool `json:"ensure_parents,omitempty" toml:"ensure_parents,omitempty"`
}

// FileCustomization represents a file to be created in the image
type FileCustomization struct {
	// Absolute path to the file
	Path string `json:"path" toml:"path"`
	// Owner of the directory specified as a string (user name), int64 (UID) or nil
	User interface{} `json:"user,omitempty" toml:"user,omitempty"`
	// Owner of the directory specified as a string (group name), int64 (UID) or nil
	Group interface{} `json:"group,omitempty" toml:"group,omitempty"`
	// Permissions of the file specified as an octal number
	Mode string `json:"mode,omitempty" toml:"mode,omitempty"`
	// Data is the file content in plain text
	Data string `json:"data,omitempty" toml:"data,omitempty"`
}
