package blueprint

type FIPS struct {
	// Enables the system FIPS mode (disabled by default). Currently only edge-raw-image, edge-installer,
	// edge-simplified-installer, edge-ami and edge-vsphere images support this customization.
	Enabled bool `json:"enabled,omitempty"`
}
