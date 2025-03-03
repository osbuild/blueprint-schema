## Blueprint schema

**WORK IN PROGRESS** but please send us feedback.

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

### Why

Users of Image Builder must work with several blueprint customization formats:

* composer uses TOML format also known as "on-prem" blueprint
* image-builder-cli accept TOML and JSON "on-prem" blueprint
* bootc-image-builder accept a subset of the composer TOML schema
* console.redhat.com API has a modified subset of the same but as JSON (and different from the cli JSON)
* there are few more intermediate/technical blueprint formats but you get the idea

At the end of 2024 we held a series of meetings where the core team agreed on a common blueprint schema we would like to converge to and this repository contains exactly that, it contains:

* blueprint JSON Schema
* Go library for loading/saving JSON and YAML
* Go validating library and CLI
* YAML examples with automated tests
* documentation with examples (from automated tests)
* conversion library and CLI app for both "on-prem" and "console" formats

### An example

If you just want an example of a blueprint, [here is one](fixtures/valid-000-all-fields.in.yaml) that covers all the supported fields. There are many more in the `fixtures/` directory, read on for more information.

### The schema

Latest version of the JSON Schema is available as [blueprint-schema.json](blueprint-schema.json).

The schema is generated from [Go types](blueprint.go) using [generate-schema](cmd/generate-schema/main.go) tool. To generate a new version run:

    make generate-schema

The schema is JSON Schema Draft 2020-12 compliant and can be included in OpenAPI 3.1 endpoints.

## Go types

The schema generator uses both Go struct tags `json` and `jsonschema` as well as Go documentation to create the schema. Note although YAML is supported too, no YAML Go struct tags are required since YAML is always converted to JSON first and then loaded using JSON Go struct tags to ensure consistency.

Example type:

```go
type Locale struct {
	// The languages attribute is a list of strings that contains the languages to be installed on the image.
	// To list available languages, run: localectl list-locales
	Languages []string `json:"languages,omitempty" jsonschema:"nullable,default=en_US.UTF-8"`

	// The keyboards attribute is a list of strings that contains the keyboards to be installed on the image.
	// To list available keyboards, run: localectl list-keymaps
	Keyboards []string `json:"keyboards,omitempty" jsonschema:"nullable,default=us"`
}
```

There are few types where defining schema via Go struct tags is unreadable or limiting. In that cases, files with the pattern `blueprint_*.yaml` are embedded and loaded into memory and compiled into the schema during application load. These are called schema *partials* and they can only be written in YAML which is then turned into JSON (Schema). Example:

```yaml
---
properties:
  selected:
    type: array
    items:
      type: string
  unselected:
    type: array
    items:
      type: string
  json_profile_id:
    type: string
  json_filepath:
    type: string
oneOf:
  - anyOf:
    - required:
        - selected
    - required:
        - unselected
    - required:
        - selected
        - unselected
  - required:
      - json_profile_id
      - json_filepath
```

Then used from a Go type via `JSONSchema` method which overrides all `jsonschema` Go types for that particular type. 

```go
func (OpenSCAPTailoring) JSONSchema() *jsonschema.Schema {
	return PartialSchema("blueprint_openscap.yaml")
}
```

It is also possible to achieve a hybrid approach with `JSONSchemaExtend` or `JSONSchemaProperty` methods, this is preferred because Go struct fields (properties) do automatically get description fields from Go documentation, overriding the whole schema with properties required description to be copied. Example:

```yaml
---
oneOf:
  - anyOf:
    - required:
        - selected
    - required:
        - unselected
    - required:
        - selected
        - unselected
  - required:
      - json_profile_id
      - json_filepath
```

With the following code:

```go
func (OpenSCAPTailoring) JSONSchemaExtend(s *jsonschema.Schema) {
	s.OneOf = PartialSchema("blueprint_openscap.yaml").OneOf
}
```

Read [jsonschema](https://github.com/invopop/jsonschema) library for more details about available Go struct tags and supported features. More information about the JSON Schema can be found [on the project webpage](https://json-schema.org).

## Go documentation

All Go `godoc` documentation is used for JSON Schema `description` fields. Please keep in mind that only the following syntax is supported:

* Plain text
* Lines are unwrapped (all `\n` are removed)
* Paragraphs are preserved (all `\n\n` are kept)

## Using the types in Go

This repository is a Go package that can be used to access the schema itself, load, save and validate data. The raw schema JSON is available in `blueprint.SchemaJSON` variable. The package provides marshaling functions both to JSON and YAML:

* `ReadJSON`/`WriteJSON` for `io.Reader`
* `UnmarshalJSON`/`MarshalJSON` for `[]byte`
* `ReadYAML`/`WriteYAML` for `io.Reader`
* `UnmarshalYAML`/`MarshalYAML` for `[]byte`
* `ConvertJSONtoYAML` for `[]byte`
* `ConvertYAMLtoJSON` for `[]byte`

Example use:

```go
package main

import (
    "bytes"

    blueprint "github.com/osbuild/blueprint-schema"
)

func main() {
    yf, _ := os.Open("example.yaml")
    defer yf.Close()

    bp, _ := blueprint.ReadYAML(yf)
    println(bp.Name)
}
```

JSON writing functions can optionally indent the output.

The package has minimum dependencies, only two `yaml` libraries are needed (YAML loading, YAML-JSON conversion).

## Validating the schema from the command line

The command line utility named `blueconv` can be used for validations. To validate a YAML file:

    go run ./cmd/blueconv -validate-yaml < fixtures/openscap-invalid-both.in.yaml 

To validate a JSON file:

    go run ./cmd/blueconv -validate-json < fixtures/minimal-j.in.json

Returns 0 when schema is valid, 1 otherwise with detailed information printed on the standard output. Example schema error reported by the validator:

```json
{
  "valid": false,
  "keywordLocation": "",
  "instanceLocation": "",
  "errors": [
    {
      "valid": false,
      "keywordLocation": "/properties/ignition/oneOf",
      "instanceLocation": "/ignition",
      "errors": [
        {
          "valid": false,
          "keywordLocation": "/properties/ignition/oneOf/0/$ref/oneOf",
          "instanceLocation": "/ignition",
          "error": "oneOf failed, subschemas 0, 1 matched"
        },
        {
          "valid": false,
          "keywordLocation": "/properties/ignition/oneOf/1/type",
          "instanceLocation": "/ignition",
          "error": "got object, want null"
        }
      ]
    }
  ]
}
```

To validate the JSON Schema, use `CompileSchema` function:

```go
package main

import (
    "os"

    blueprint "github.com/osbuild/blueprint-schema"
)

func main() {
    schema, _ := blueprint.CompileSchema()
    err := schema.ReadAndValidateYAML(os.Stdin)

    println(err)
}
```

Read [jsonschema](https://github.com/santhosh-tekuri/jsonschema) library documentation for more information about the error output. The CLI utility provides the same output format as the validation library.

## Testing

A fixture-based test is available in the [fixtures/](fixtures/) directory, each fixture consist of:

* `filename.in.yaml` - input file (can be YAML or JSON)
* `filename.out.yaml` - output file after parsing and write (always YAML)
* `filename.validator.json` - output of the validator (always JSON - see the example above)

Each `*.in.*` file is loaded, YAML converted to JSON (if needed), parsed into the blueprint type and written to YAML `*.out.yaml` file. At the same time, the data is loaded into `map[string]any` and validated against the JSON Schema and results written to `*.validator.json`.

To run tests do:

    make test

To regenerate `*.out.yaml` and `*.validator.json` files (after a breaking change), do:

    make write-fixtures

All `*.validator.json` files have a custom JSON marshaller that sorts arrays and some string enumerations so it is diff-friendly.

## Editor schema support

For VS Code with Red Hat's [YAML plugin](https://github.com/redhat-developer/vscode-yaml), put the following into the settings:

```
    "json.schemas": [
        {
            "fileMatch": [
                "/fixtures/*.json",
            ],
            "url": "https://raw.githubusercontent.com/osbuild/blueprint-schema/refs/heads/main/blueprint-schema.json"
        }
    ],
    "yaml.schemas": {
        "https://raw.githubusercontent.com/osbuild/blueprint-schema/refs/heads/main/blueprint-schema.json": "/fixtures/*.yaml"
    },
```

Keep in mind that relative paths are not supported, use absolute URL instead. For example: `file:///home/lzap/blueprint-schema/blueprint-schema.json`.

## Attestation

Blueprint does not carry enough information to build an image, additional input is necessary:

* Image type: `ami`, `vhd`, `gce`, `qcow2`, `tar`, `vmdk`, `image-installer` and others
* Distribution: `fedora`, `centos` or `rhel`
* Version: `42` or `9.5`
* Architecture: `x86_64` or `arm64`
* Upload information (AWS/Azure/GCP credentials)
* Additional `ostree` information (optional)

Only after some of these fields are known, then additional validation can be done. For example:

* Field `installer.anaconda` does not apply for any image types which are not installers.
* Field `openscap` cannot be used for distros older than a specific version.
* Field `ignition` is only available for iot/edge image types.

There are two options how to achieve that:

### Via JSON Schema

```yaml
---
name: "Example of a blueprint attestation"
attestation:
    type: ami
    distribution: rhel
    version: 8.3
...
```

* Pretty complex and unreadable (e.g. test for distro version string - regular expression)
* Additional input must be part of the schema itself.
* Users have a great editing experience (if they fill the attestation in)

### Via Go code

```go
aData := Attestation{
    Type: "ami"
    Distribution: "rhel"
    Version: "8.3"
}
schema.ReadAndAttestYAML(os.Stdin, aData)
```

* Limitless validations.
* Easy to work with.
* Go struct tags can be leveraged to do the job.
* When compiled to WASM/WASI could be accessible through a web browser editor.
* Conversions to/from YAML/CRC-JSON will be done in Go anyway.

### Building

To build the library and the CLI:

    make build-cli -j

Compiling WASM binaries requires TinyGo, clang, WASM tools and Go version that is compatible with TinyGo, typically few major releases behind the latest and greatest. TinyGo can be used from Fedora, but it is not compatible with its `golang` package version, so `WASM_GOROOT` variable must be explicitly passed:

    sudo dnf -y install tinygo clang17 binaryen
    make build-wasm WASM_GOROOT=$HOME/sdk/go1.21.0

To cross-build everything:

    make build WASM_GOROOT=$HOME/sdk/go1.21.0 -j

Builds the CLI for all supported platforms plus WASM built with Go and TinyGo. The latter is significantly smaller binary (2M) which thanks to HTTP compression can be as small as 1M.

```
$ ls dist/ -1
blueconv_darwin_amd64
blueconv_darwin_arm64
blueconv_linux_amd64
blueconv_linux_arm64
blueconv_windows_amd64
blueconv_windows_arm64
blueprint_go.wasm
blueprint_tgo.wasm
```

### Conversion tools

WORK IN PROGRESS

```
go run ./cmd/blueconv/ -convert < fixtures/valid-000-all-fields.in.yaml 
```

Generates something like

```
name = "Blueprint example"
description = "A complete example of a blueprint with all possible fields"

# ...
```

### Schema documentation

TBD

### Links

* https://github.com/invopop/yaml - library to convert YAML to JSON and vice versa
* https://github.com/invopop/jsonschema - library to generate JSON Schema from Go types
* https://github.com/santhosh-tekuri/jsonschema - library to validate JSON Schema

### TODO

* Implement JSON Schema attestation (folks did not like the Go-based ones)
* Make a release and CLI/WASM build pipeline
* Github page with JSON/YAML editors and few starting examples
* Explore fully YAML-based schema with optionally generated Go code
* Implement conversion tools for onprem TOML/JSON and CRC-JSON
* Generate markdown/HTML documentation for the schema with examples
* Implement WASI/WASM conversion to the on github page
* Fix default values - schema definition is not used correctly for non-pointers (e.g. bool must be pointer if defaults to true)
* Add intermediate JSON into fixtures for a quick "peek" into YAML/JSON interoperability
* Implement [common defs/refs](https://tour.json-schema.org/content/06-Combining-Subschemas/01-Reusing-and-Referencing-with-defs-and-ref) for example InstallerAnaconda.EnabledModules.
