package blueprint

type Ignition struct {
	// The URL to the Ignition configuration to be used by Ignition. This configuration is a URL to a remote Ignition
	// configuration. The firstboot_url is used if the embedded configuration is not specified.
	//
	// Cannot be used with embedded_base64 or embedded_text.
	FirstbootURL string `json:"firstboot_url,omitempty" jsonschema:"oneof_required=ignition_url"`

	// The embedded Ignition configuration to be used by Ignition. This configuration is embedded in the blueprint.
	//
	// Cannot be used with firstboot_url.
	Embedded *IgnitionEmbedded `json:"embedded,omitempty" jsonschema:"oneof_required=ignition_embedded"`
}

type IgnitionEmbedded struct {
	// Ignition data formatted in plain text.
	Text string `json:"text,omitempty" jsonschema:"oneof_required=ignition_text"`

	// Ignition data formatted in base64.
	Base64 string `json:"base64,omitempty" jsonschema:"oneof_required=ignition_base64"`
}
