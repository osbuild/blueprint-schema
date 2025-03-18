package blueprint

type Systemd struct {
	// The enabled attribute is a list of strings that contains the systemd units to be enabled.
	Enabled []string `json:"enabled,omitempty" jsonschema:"nullable"`

	// The disabled attribute is a list of strings that contains the systemd units to be disabled.
	Disabled []string `json:"disabled,omitempty" jsonschema:"nullable"`

	// The masked attribute is a list of strings that contains the systemd units to be masked.
	Masked []string `json:"masked,omitempty" jsonschema:"nullable"`
}
