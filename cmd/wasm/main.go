package main

import (
	"strings"
	"unsafe"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/blueprint"
	"github.com/osbuild/blueprint-schema/pkg/conv/notes"
	"github.com/osbuild/blueprint-schema/pkg/conv/onprem"
	"github.com/osbuild/blueprint-schema/pkg/conv/ptr"
)

func main() {}

func getJsString(pointer, length int) string {
	// Convert memory pointer and length to Go string
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(*(*uint8)(unsafe.Pointer(uintptr(pointer + i))))
	}
	return string(bytes)
}

//go:wasmexport BlueprintValidateJSON
func BlueprintValidateJSON(pointer, length int) *string {
	input := getJsString(pointer, length)

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
//export BlueprintExportTOML
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
