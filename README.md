## Blueprint schema

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

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

    blueprint "github.com/lzap/common-blueprint-example"
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

    go run ./cmd/validate-schema < example.yaml

To validate a JSON file:

    go run ./cmd/validate-schema -json < example.json

Returns 0 when schema is valid, 1 otherwise with detailed information formatted as JSON on the standard output. Example schema error reported by the validator:

```json
{
  "valid": false,
  "evaluationPath": "",
  "schemaLocation": "",
  "instanceLocation": "",
  "annotations": {
    "description": "Blueprint type prototype\n\nThis is just a brief example of a common blueprint structure. Just few fields\nwere selected to demonstrate the JSON schema."
  },
  "errors": {
    "properties": "Property 'name' does not match the schema",
    "required": "Required property 'name' is missing"
  },
  "details": [
    {
      "valid": false,
      "evaluationPath": "/properties/name",
      "schemaLocation": "https://github.com/lzap/common-blueprint-example/blueprint#/properties/name",
      "instanceLocation": "/name",
      "annotations": {
        "description": "Name of the blueprint"
      },
      "errors": {
        "type": "Value is null but should be string"
      }
    }
  ]
}
```

## Validating the schema in Go

To minimize dependencies of the main `blueprint` package, a separate package named `validate` must be used. Read [jsonschema](https://github.com/kaptinlin/jsonschema) library documentation for more information about the error output.

```go
package main

import (
    "os"

    "github.com/lzap/common-blueprint-example/validate"
)

func main() {
    // compile the schema which embedded as part of this package
    schema, _ := validate.CompileSchema()

    // returns bool, string and err
    valid, out, _ := schema.ReadAndValidateYAML(os.Stdin)

    println(valid, out)
}
```

The CLI utility provides the same output format as the validation library.

## Testing

A fixture-based test is available in the [validate/fixtures/](validate/fixtures/) directory, each fixture consist of:

* `filename.in.yaml` - input file (can be YAML or JSON)
* `filename.out.yaml` - output file after parsing and write (always YAML)
* `filename.valid.json` - output of the validator (always JSON)

Each `*.in.*` file is loaded, parsed, YAML converted to JSON (if needed), validated and written to YAML `*.out.yaml` file. At the same time, the data is loaded into `map[string]any` and validated against the JSON Schema and results written to `*.valid.json`.

To run tests do:

    make test

To regenerate `*.out.yaml` and `*.valid.json` files (after a breaking change), do:

    make write-fixtures

## Links

* https://github.com/invopop/jsonschema - library to generate JSON Schema from Go types
* https://github.com/kaptinlin/jsonschema - library to validate JSON Schema

## TODO

* Only fraction of the main `example.yaml` is currently implemented as Go type - this will be done after we agree on the final version of the example data.
* Multiple schemas could be conditionally generated per-service: crc and cli (some fields could be hidden/marked as deprecated). This could be implemented via `onlyFor:"crc"` Go struct tag for example. Examples are in some types but this is not implemented.
