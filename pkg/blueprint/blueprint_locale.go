package blueprint

type Locale struct {
	// The languages attribute is a list of strings that contains the languages to be installed on the image.
	// To list available languages, run: localectl list-locales
	Languages []string `json:"languages,omitempty" jsonschema:"nullable,default=en_US.UTF-8"`

	// The keyboards attribute is a list of strings that contains the keyboards to be installed on the image.
	// To list available keyboards, run: localectl list-keymaps
	Keyboards []string `json:"keyboards,omitempty" jsonschema:"nullable,default=us"`
}
