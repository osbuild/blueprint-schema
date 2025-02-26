package blueprint

type Ignition struct {
	// The embedded configuration to be used by Ignition as base64-encoded contents.
	//
	// Cannot be used with embedded_text or firstboot_url.
	EmbeddedBase64 string `json:"embedded_base64,omitempty" jsonschema:"oneof_required=ignition_base64"`

	// The embedded configuration to be used by Ignition as plain text.
	//
	// Cannot be used with embedded_base64 or firstboot_url.
	EmbeddedText string `json:"embedded_text,omitempty" jsonschema:"oneof_required=ignition_text"`

	// The URL to the Ignition configuration to be used by Ignition. This configuration is a URL to a remote Ignition
	// configuration. The firstboot_url is used if the embedded configuration is not specified.
	//
	// Cannot be used with embedded_base64 or embedded_text.
	FirstbootURL string `json:"firstboot_url,omitempty" jsonschema:"oneof_required=ignition_url"`
}
