//go:build js
// +build js

package main

import (
	"strings"

	"github.com/osbuild/blueprint-schema"
)

//go:wasmexport
func BlueprintValidateJSON(input string) string {
	schema, err := blueprint.CompileSchema()
	if err != nil {
		return err.Error()
	}

	err = schema.ReadAndValidateJSON(strings.NewReader(input))
	if err != nil {
		return err.Error()
	}

	return ""
}

//go:wasmexport
func BlueprintValidateYAML(input string) string {
	schema, err := blueprint.CompileSchema()
	if err != nil {
		return err.Error()
	}

	err = schema.ReadAndValidateYAML(strings.NewReader(input))
	if err != nil {
		return err.Error()
	}

	return ""
}
