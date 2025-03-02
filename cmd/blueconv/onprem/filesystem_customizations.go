package blueprint

type FilesystemCustomization struct {
	Mountpoint string `json:"mountpoint,omitempty" toml:"mountpoint,omitempty"`
	MinSize    string `json:"min_size,omitempty" toml:"min_size,omitempty"`
}
