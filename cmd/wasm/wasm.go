package main

import (
	"context"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/parse"
	"sigs.k8s.io/yaml"
)

func wasmValidateUBP(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return js.ValueOf("Validation expects exactly one argument (UBP string)")
	}

	schema, err := parse.CompileBundledSchema()
	if err != nil {
		return js.ValueOf(fmt.Sprintf("Failed to compile schema: %v", err))
	}

	err = schema.ValidateAny(context.Background(), []byte(p[0].String()))
	if err != nil {
		return js.ValueOf(fmt.Sprintf("Validation failed:\n\n%s", err.Error()))
	}

	return js.ValueOf("")
}

func wasmExportTOML(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return js.ValueOf([]any{"", "Export TOML expects exactly one argument (UBP string)"})
	}

	b, _, err, warn := parse.UnmarshalAny([]byte(p[0].String()))
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to unmarshal: %v", err)})
	}
	if warn != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Unexpected warning(s): %v", err)})
	}

	exporter := conv.NewInternalExporter(b)
	result, logs := exporter.Export()
	if logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	buf, err := toml.Marshal(result)
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to marshal TOML: %v", err)})
	}

	return js.ValueOf([]any{string(buf), ""})
}

func wasmExportJSON(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return js.ValueOf([]any{"", "Export TOML expects exactly one argument (UBP string)"})
	}

	b, _, err, warn := parse.UnmarshalAny([]byte(p[0].String()))
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to unmarshal: %v", err)})
	}
	if warn != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Unexpected warning(s): %v", err)})
	}

	exporter := conv.NewInternalExporter(b)
	result, logs := exporter.Export()
	if logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	buf, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to marshal TOML: %v", err)})
	}

	return js.ValueOf([]any{string(buf), ""})
}

func wasmImportYAML(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return js.ValueOf([]any{"", "Import TOML expects exactly one argument (BP string)"})
	}

	_, b, err, warn := parse.UnmarshalAny([]byte(p[0].String()))
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to unmarshal: %v", err)})
	}
	if warn != nil {
		js.Global().Get("console").Call("warn", warn.Error())
	}

	importer := conv.NewInternalImporter(b)
	result, logs := importer.Import()
	if logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	buf, err := yaml.Marshal(result)
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to marshal YAML: %v", err)})
	}

	return js.ValueOf([]any{string(buf), ""})
}

func wasmImportJSON(this js.Value, p []js.Value) any {
	if len(p) != 1 {
		return js.ValueOf([]any{"", "Import TOML expects exactly one argument (BP string)"})
	}

	_, b, err, warn := parse.UnmarshalAny([]byte(p[0].String()))
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to unmarshal: %v", err)})
	}
	if warn != nil {
		js.Global().Get("console").Call("warn", warn.Error())
	}

	importer := conv.NewInternalImporter(b)
	result, logs := importer.Import()
	if logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	buf, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return js.ValueOf([]any{"", fmt.Sprintf("Failed to marshal JSON: %v", err)})
	}

	return js.ValueOf([]any{string(buf), ""})
}

func main() {
	js.Global().Set("wasmValidateUBP", js.FuncOf(wasmValidateUBP))
	js.Global().Set("wasmExportTOML", js.FuncOf(wasmExportTOML))
	js.Global().Set("wasmExportJSON", js.FuncOf(wasmExportJSON))
	js.Global().Set("wasmImportYAML", js.FuncOf(wasmImportYAML))
	js.Global().Set("wasmImportJSON", js.FuncOf(wasmImportJSON))

	// Keep the program running
	select {}
}
