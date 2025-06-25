package main

import (
	"context"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/parse"
)

func wasmValidateUBP(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return map[string]interface{}{
			"error": "Validation expects exactly one argument (UBP string)",
		}
	}

	schema, err := parse.CompileBundledSchema()
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to compile schema: %v", err),
		}
	}

	result := "Validation succeeded"
	err = schema.ValidateAny(context.Background(), []byte(p[0].String()))
	if err != nil {
		result = fmt.Sprintf("Validation failed:\n%s", err)
	}

	return map[string]interface{}{
		"report": result,
	}
}

func wasmExportTOML(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return map[string]interface{}{
			"error": "Export TOML expects exactly one argument (UBP string)",
		}
	}

	b, err := parse.UnmarshalYAML([]byte(p[0].String()))
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to unmarshal YAML: %v", err),
		}
	}

	// Export using the internal exporter
	exporter := conv.NewInternalExporter(b)
	if logs := exporter.Export(); logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	// Marshal to TOML
	buf, err := toml.Marshal(exporter.Result())
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to marshal TOML: %v", err),
		}
	}

	return map[string]interface{}{
		"toml": string(buf),
	}
}

func wasmExportJSON(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return map[string]interface{}{
			"error": "Export JSON expects exactly one argument (UBP string)",
		}
	}

	yamlInput := p[0].String()
	inBuf := []byte(yamlInput)

	// Parse the YAML input
	b, err := parse.UnmarshalYAML(inBuf)
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to unmarshal YAML: %v", err),
		}
	}

	// Export using the internal exporter
	exporter := conv.NewInternalExporter(b)
	if logs := exporter.Export(); logs != nil {
		js.Global().Get("console").Call("warn", logs.Error())
	}

	// Marshal to JSON
	buf, err := json.MarshalIndent(exporter.Result(), "", "  ")
	if err != nil {
		return map[string]interface{}{
			"error": fmt.Sprintf("Failed to marshal JSON: %v", err),
		}
	}

	return map[string]interface{}{
		"json": string(buf),
	}
}

func main() {
	js.Global().Set("wasmValidateUBP", js.FuncOf(wasmValidateUBP))
	js.Global().Set("wasmExportTOML", js.FuncOf(wasmExportTOML))
	js.Global().Set("wasmExportJSON", js.FuncOf(wasmExportJSON))

	// Keep the program running
	select {}
}
