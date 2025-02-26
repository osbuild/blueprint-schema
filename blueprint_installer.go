package blueprint

import (
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
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
	EnabledModules []string `json:"enabled_modules,omitempty"`

	// Disabled modules. See enabled flag for more information.
	DisabledModules []string `json:"disabled_modules,omitempty"`

	// Kickstart file formatted in plain text.
	KickstartText string `json:"kickstart_text,omitempty" jsonschema:"oneof_required=kickstart_text"`

	// Kickstart file formatted in base64.
	KickstartBase64 string `json:"kickstart_base64,omitempty" jsonschema:"oneof_required=kickstart_base64"`
}

type CoreOSInstaller struct {
	// Installation device for iot/edge simplified installer image types.
	InstallationDevice string `json:"installation_device,omitempty"`
}

// JSONSchemaExtend can be used to extend the generated JSON schema from Go struct tags
func (AnacondaInstaller) JSONSchemaExtend(s *jsonschema.Schema) {
	enum := &jsonschema.Schema{
		Type: "array",
		Items: &jsonschema.Schema{
			Type: "string",
			Enum: []any{
				"org.fedoraproject.Anaconda.Modules.Localization",
				"org.fedoraproject.Anaconda.Modules.Network",
				"org.fedoraproject.Anaconda.Modules.Payloads",
				"org.fedoraproject.Anaconda.Modules.Runtime",
				"org.fedoraproject.Anaconda.Modules.Security",
				"org.fedoraproject.Anaconda.Modules.Services",
				"org.fedoraproject.Anaconda.Modules.Storage",
				"org.fedoraproject.Anaconda.Modules.Subscription",
				"org.fedoraproject.Anaconda.Modules.Timezone",
				"org.fedoraproject.Anaconda.Modules.Users",
			},
		},
	}

	enabledPair := orderedmap.Pair[string, *jsonschema.Schema]{
		Key:   "enabled_modules",
		Value: enum,
	}
	disabledPair := orderedmap.Pair[string, *jsonschema.Schema]{
		Key:   "disabled_modules",
		Value: enum,
	}
	s.Properties.AddPairs(enabledPair, disabledPair)
}
