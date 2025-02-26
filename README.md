## Blueprint schema

**WORK IN PROGRESS** but please send us feedback.

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

### The schema

Latest version of the JSON Schema is available as [blueprint-schema.json](blueprint-schema.json).

The schema is generated from [Go types](blueprint.go) using [generate-schema](cmd/generate-schema/main.go) tool. To generate a new version run:

    make generate-schema

The schema is JSON Schema Draft 2020-12 compliant and can be included in OpenAPI 3.1 endpoints.

## Go types

The schema generator uses both Go struct tags `json` and `jsonschema` as well as Go documentation to create the schema. Note although YAML is supported too, no YAML Go struct tags are required since YAML is always converted to JSON first and then loaded using JSON Go struct tags to ensure consistency.

Read [jsonschema](https://github.com/invopop/jsonschema) library for more details about available Go struct tags and supported features.

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

To validate a YAML file:

    go run ./cmd/validate-schema < fixtures/openscap-invalid-both.in.yaml 

To validate a JSON file:

    go run ./cmd/validate-schema -json < fixtures/minimal-j.in.json

Returns 0 when schema is valid, 1 otherwise with detailed information printed on the standard output. Example schema error reported by the validator:

```
validation failed: jsonschema validation failed with 'blueprint-schema.json'
- at '/openscap': oneOf failed, none matched
  - at '/openscap/tailoring': oneOf failed, none matched
    - at '/openscap/tailoring': oneOf failed, subschemas 0, 1 matched
    - at '/openscap/tailoring': got object, want null
  - at '/openscap': got object, want nullexit status 1
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
* `filename.valid.out` - output of the validator (always JSON)

Each `*.in.*` file is loaded, YAML converted to JSON (if needed), parsed into the blueprint type and written to YAML `*.out.yaml` file. At the same time, the data is loaded into `map[string]any` and validated against the JSON Schema and results written to `*.valid.out`.

To run tests do:

    make test

To regenerate `*.out.yaml` and `*.valid.out` files (after a breaking change), do:

    make write-fixtures

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
* When compiled to WASM/WASI could be accessible through a web browser editor.
* Conversions to/from YAML/CRC-JSON will be done in Go anyway.

## Links

* https://github.com/invopop/yaml - library to convert YAML to JSON and vice versa
* https://github.com/invopop/jsonschema - library to generate JSON Schema from Go types
* https://github.com/santhosh-tekuri/jsonschema - library to validate JSON Schema

## TODO

* Finalize the schema
* Implement conversion tools in both crc/images repos in ./cmd subdirectories and use those tools via "go run" command to generate a nice example set:
* https://github.com/osbuild/image-builder-crc/blob/main/internal/v1/api.go#L663
* https://github.com/osbuild/blueprint-schema/blob/main/blueprint.go#L63
* Generate markdown/HTML documentation for the schema with examples
* Attestations
* Github page with JSON/YAML editors
* Example loading on github page
* WASI/WASM conversion API and convertor on github page
