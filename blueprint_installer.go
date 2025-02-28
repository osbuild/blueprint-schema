package blueprint

import (
	"github.com/invopop/jsonschema"
)

type Installer struct {
	// Extra customizations for Anaconda installer (ISO) image types.
	//
	// Blueprint customizations that match the kickstart options (languages, keyboard, timezone)
	// will change the value in the kickstart file as well.
	Anaconda *AnacondaInstaller `json:"anaconda,omitempty" jsonschema:"nullable"`

	// CoreOS installer configuration is required by the edge-simplified-installer image. It allows to define the
	// destination device for the installation.
	CoreOS *CoreOSInstaller `json:"coreos,omitempty" jsonschema:"nullable"`
}

type AnacondaInstaller struct {
	// Unattended installation Anaconda flag. When not set, Anaconda installer will ask for user input.
	Unattended bool `json:"unattended,omitempty"`

	// Sudo users with NOPASSWD option. Adds a snippet to the kickstart file that, after installation, will create
	// drop-in files in /etc/sudoers.d to allow the specified users and groups to run sudo without a password
	// (groups must be prefixed with %).
	SudoNOPASSWD []string `json:"sudo_nopasswd,omitempty" jsonschema:"maxLength=256,pattern=^[%a-zA-Z0-9_.][a-zA-Z0-9_.$-]*$"`

	// Enabled modules. The Anaconda installer can be configured by enabling or disabling its D-Bus modules.
	// Modules must be prefixed with "org.fedoraproject.Anaconda.Modules." and only those from the schema enum
	// list are supported.
	//
	// By default, the following modules are enabled for all Anaconda ISOs: Network, Payloads, Storage.
	//
	// The disable list is processed after the enable list and therefore takes priority. In other words, adding the
	// same module in both enable and disable will result in the module being disabled. Furthermore, adding a module
	// that is enabled by default to disable will result in the module being disabled.
	EnabledModules []string `json:"enabled_modules,omitempty" jsonschema:"nullable"`

	// Disabled modules. See enabled flag for more information.
	DisabledModules []string `json:"disabled_modules,omitempty" jsonschema:"nullable"`

	// Kickstart installer configuration.
	Kickstart *KickstartInstaller `json:"kickstart,omitempty" jsonschema:"nullable"`
}

type CoreOSInstaller struct {
	// Installation device for iot/edge simplified installer image types.
	InstallationDevice string `json:"installation_device,omitempty"`
}

type KickstartInstaller struct {
	// Kickstart file formatted in plain text.
	Text string `json:"text,omitempty" jsonschema:"oneof_required=kickstart_text"`

	// Kickstart file formatted in base64.
	Base64 string `json:"base64,omitempty" jsonschema:"oneof_required=kickstart_base64"`
}

// JSONSchemaExtend can be used to extend the generated JSON schema from Go struct tags
func (AnacondaInstaller) JSONSchemaExtend(s *jsonschema.Schema) {
	ps := PartialSchema("blueprint_installer.yaml")
	s.Properties.AddPairs(
		*ps.Properties.GetPair("enabled_modules"),
		*ps.Properties.GetPair("disabled_modules"),
	)
}
