package blueprint

type Kernel struct {
	// Kernel DNF package name to replace the standard kernel with.
	Package string `json:"package,omitempty"`

	// An optional string to append arguments to the bootloader kernel command line. The list
	// will be concatenated with spaces.
	CmdlineAppend []string `json:"cmdline_append,omitempty" jsonschema:"nullable"`
}
