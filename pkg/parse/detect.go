package parse

import (
	"bytes"
	"encoding/json"
	"slices"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// detectedFormat represents the detected data format.
type detectedFormat string

const (
	formatJSON    detectedFormat = "json"
	formatYAML    detectedFormat = "yaml"
	formatTOML    detectedFormat = "toml"
	formatUnknown detectedFormat = "unknown"
)

// Type represents the detected structure type.
type Type string

const (
	TypeUBP     Type = "ubp"
	TypeBP      Type = "bp"
	TypeUnknown Type = "unknown"
)

// detectFormat analyzes a byte slice and returns whether it is likely
// TOML, JSON, or YAML. It prioritizes JSON over YAML since any valid
// JSON is also valid YAML.
func detectFormat(data []byte) (detectedFormat, map[string]any) {
	obj := make(map[string]any)
	tdata := bytes.TrimSpace(data)
	if len(tdata) == 0 {
		return formatUnknown, obj
	}

	obj = make(map[string]any)
	if json.Unmarshal(tdata, &obj) == nil && len(obj) > 0 {
		return formatJSON, obj
	}

	obj = make(map[string]any)
	if err := toml.Unmarshal(tdata, &obj); err == nil && len(obj) > 0 {
		return formatTOML, obj
	}

	obj = make(map[string]any)
	if yaml.Unmarshal(tdata, &obj) == nil && len(obj) > 0 {
		return formatYAML, obj
	}

	return formatUnknown, make(map[string]any)
}

var uniqueBPKeys []string = []string{
	"version",
	"packages",
	"modules",
	"enabled_modules",
	"groups",
	"customizations",
	"distro",
	"minimal",
}

var uniqueUBPKeys []string = []string{
	"accounts",
	"cacerts",
	"distribution",
	"dnf",
	"fips",
	"fsnodes",
	"hostname",
	"ignition",
	"installer",
	"kernel",
	"locale",
	"network",
	"openscap",
	"registration",
	"storage",
	"systemd",
	"timedate",
}

func init() {
	for _, key := range uniqueBPKeys {
		if slices.Contains(uniqueUBPKeys, key) {
			panic("Detected common keys between UBP and BP: " + key)
		}
	}
}

// DetectType analyzes a map and returns whether it is likely
// a UBP or BP structure based on the presence of unique keys.
// It uses a naive approach to check for the presence of top-level
// keys that are unique to each structure type. It is not 100% accurate.
func DetectType(data map[string]any) Type {
	for _, key := range uniqueUBPKeys {
		if _, exists := data[key]; exists {
			//println("matched UBP key:", key)
			return TypeUBP
		}
	}
	for _, key := range uniqueBPKeys {
		if _, exists := data[key]; exists {
			//println("matched BP key:", key)
			return TypeBP
		}
	}

	return TypeUnknown
}
