package blueprint

import "github.com/invopop/jsonschema"

type OpenSCAP struct {
	// The desired security profile ID.
	ProfileID string `json:"profile_id,omitempty" jsonschema:"required"`

	// Datastream to use for the scan. The datastream is the path to the SCAP datastream file to use for the scan.
	// If the datastream parameter is not provided, a sensible default based on the selected distro will be chosen.
	Datastream string `json:"datastream,omitempty"`

	// An optional OpenSCAP tailoring information. Can be done via profile selection or tailoring JSON file.
	//
	// In case of profile selection, a tailoring file with a new tailoring profile ID is created and saved to the image.
	// The new tailoring profile ID is created by appending the _osbuild_tailoring suffix to the base profile.
	// For example, given tailoring options for the cis profile, tailoring profile
	// xccdf_org.ssgproject.content_profile_cis_osbuild_tailoring will be created. The default namespace of the rules
	// is org.ssgproject.content, so the prefix may be omitted for rules under this namespace, i.e.
	// org.ssgproject.content_grub2_password and grub2_password are functionally equivalent.
	// The generated tailoring file is saved to the image as /usr/share/xml/osbuild-oscap-tailoring/tailoring.xml or,
	// for newer releases, in the /oscap_data directory, this is the location used for other OpenSCAP related artifacts.
	//
	// It is also possible to use JSON tailoring. In that case, custom JSON file must be provided using the blueprint and
	// used in json_filepath field alongside with json_profile_id field. The generated XML tailoring file is saved to the
	// image as /oscap_data/tailoring.xml.
	Tailoring *OpenSCAPTailoring `json:"tailoring,omitempty" jsonschema:"nullable"`
}

type OpenSCAPTailoring struct {
	// Selected rules, cannot be used with json_profile_id and json_filepath.
	Selected []string `json:"selected,omitempty"`

	// Unselected rules, cannot be used with json_profile_id and json_filepath.
	Unselected []string `json:"unselected,omitempty"`

	// JSON profile ID, must be used with json_filepath and cannot be used with selected and unselected fields.
	JSONProfileID string `json:"json_profile_id,omitempty"`

	// JSON filepath, must be used with json_profile_id and cannot be used with selected and unselected fields.
	JSONFilepath string `json:"json_filepath,omitempty"`
}

// JSONSchema can be used to generate JSON schema programmatically.
func (OpenSCAPTailoring) JSONSchema() *jsonschema.Schema {
	return PartialSchema("blueprint_openscap.yaml")
}
