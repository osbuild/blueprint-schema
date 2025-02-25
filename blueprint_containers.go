package blueprint

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
