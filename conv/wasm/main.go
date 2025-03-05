package main

import (
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema"
	"github.com/osbuild/blueprint-schema/conv/notes"
	"github.com/osbuild/blueprint-schema/conv/onprem"
	"github.com/osbuild/blueprint-schema/conv/ptr"
)

func main() {}

//go:wasmexport BlueprintValidateJSON
func BlueprintValidateJSON(input string) *string {
	schema, err := blueprint.CompileSchema()
	if err != nil {
		return ptr.To(err.Error())
	}

	err = schema.ReadAndValidateJSON(strings.NewReader(input))
	if err != nil {
		return ptr.To(err.Error())
	}

	return ptr.To("")
}

//go:wasmexport BlueprintValidateYAML
func BlueprintValidateYAML(input string) *string {
	schema, err := blueprint.CompileSchema()
	if err != nil {
		return ptr.To(err.Error())
	}

	err = schema.ReadAndValidateYAML(strings.NewReader(input))
	if err != nil {
		return ptr.To(err.Error())
	}

	return ptr.To("")
}

//go:wasmexport BlueprintExportTOML
//export BlueprintValidateJSON
func BlueprintExportTOML(input string) *string {
	from, err := blueprint.ReadYAML(strings.NewReader(input))
	if err != nil {
		return ptr.To(err.Error())
	}

	nts := notes.ConversionNotes{}
	to := onprem.ExportBlueprint(from, &nts)

	// XXX ignore errors for now

	output := strings.Builder{}
	err = toml.NewEncoder(&output).Encode(to)
	if err != nil {
		return ptr.To(err.Error())
	}

	return ptr.To(output.String())
}
