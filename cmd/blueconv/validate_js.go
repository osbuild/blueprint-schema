//go:build js
// +build js

package main

import (
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema"
	"github.com/osbuild/blueprint-schema/conv/onprem"
	onprem_blueprint "github.com/osbuild/blueprint-schema/conv/onprem/blueprint"
	"github.com/osbuild/blueprint-schema/conv/ptr"
)

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
func BlueprintExportTOML(input string) *string {
	from, err := blueprint.ReadYAML(strings.NewReader(input))
	if err != nil {
		return ptr.To(err.Error())
	}

	to := onprem_blueprint.Blueprint{}
	errs := onprem.Errors{}
	onprem.ExportBlueprint(&to, from, &errs)

	// XXX ignore errors for now

	output := strings.Builder{}
	err = toml.NewEncoder(&output).Encode(to)
	if err != nil {
		return ptr.To(err.Error())
	}

	return ptr.To(output.String())
}
