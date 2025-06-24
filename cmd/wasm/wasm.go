package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/parse"
)

// exportTOML converts YAML blueprint input to TOML format
func exportTOML(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return map[string]interface{}{
			"error": "exportTOML expects exactly one argument (YAML string)",
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

// exportJSON converts YAML blueprint input to JSON format
func exportJSON(this js.Value, p []js.Value) interface{} {
	if len(p) != 1 {
		return map[string]interface{}{
			"error": "exportJSON expects exactly one argument (YAML string)",
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
	// Register the WASM functions
	js.Global().Set("exportTOML", js.FuncOf(exportTOML))
	js.Global().Set("exportJSON", js.FuncOf(exportJSON))

	// Keep the program running
	select {}
}
