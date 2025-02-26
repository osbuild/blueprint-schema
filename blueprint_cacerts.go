package blueprint

type CACerts struct {
	// The PEM-encoded certificate.
	Cert string `json:"cert,omitempty" jsonschema:"required"`
}
