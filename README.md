## Blueprint schema

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

**WORK IN PROGRESS**

### Terminology

* **UBP**: Unified Blueprint (this repo). Also known as Blueprint v2.
* **Blueprint**: The blueprint as it exists in [osbuild/blueprint](https://github.com/osbuild/blueprint), is user-facing in osbuild-composer, and is documented on osbuild.org
* **CRC Blueprint**: The blueprint as it exists in [image-builder-crc](https://github.com/osbuild/image-builder-crc), is user-facing in the service API, and is documented on osbuild.org.  This is very close but sometimes slightly different from the Blueprint.
* **Images Blueprint**: The format and code that currently exists in [osbuild/images](https://github.com/osbuild/images) in `pkg/blueprint`, which was never meant to be user-facing, but through sloppiness (on my part as well) initially (but not anymore) ended up being the user-facing blueprint in bootc-image-builder and, for a short period, ib-cli.  This slightly differs from the Old Schema in that it was more aggressive in dropping deprecated options (like SSHKey), because the user-facing blueprints (Old and CRC) were responsible for backwards compatibility.

### Schema source

All schema source files are in `oas/` directory, each component resides in its own YAML file in `oas/components`. Make sure to create component for each object that is supposed to be a Go type (`struct`). There is a README in the `oas/` directory with some tips on how to write OAS3 schemas.

### Schema files

* `blueprint-oas3.yaml` - OpenAPI 3.0 schema, the blueprint document is at `#/components/schemas/blueprint`
* `blueprint-oas3.json` - the same schema but in JSON format
* `blueprint-oas3-ext.json` - the schema with additional "extensions" which would confuse Go code generator but are useful for validation purposes, they are defined in `oas/extensions`

All the mentioned files are generated using `make schema`

### Go code

Go code is generated from `blueprint-oas3.json` via `oapi-codegen` using `make schema`.

All the code resides in `pkg/blueprint` except embedded schemas from above which are in the top-level directory for technical reasons (Go embedding limitations). Direct access to schema files is not required for any scenario, so only import the former package.

### CLI tool

A simple CLI tool for schema bundling, schema validation or conversion is part of this library:

```
go run github.com/osbuild/blueprint-schema/cmd/image-builder-blueprint/ -h
```

The usage is self-explanatory:

```
Usage of image-builder-blueprint:
  -export-json
        convert document into legacy JSON
  -export-toml
        convert document into legacy TOML
  -input string
        input JSON or YAML file (defaults to standard input, detects format)
  -print-json-extended-schema
        print embedded schema to standard output and exit
  -print-json-schema
        print embedded schema to standard output and exit
  -print-yaml-schema
        print embedded schema to standard output and exit
  -validate
        validate input document (detects JSON or YAML format)
```

Example validation:

```
go run github.com/osbuild/blueprint-schema/cmd/image-builder-blueprint/ -validate -input testdata/all-fields.in.yaml
```

The return value is non-zero when validation fails and error is printed on the standard error. Example export:

```
go run github.com/osbuild/blueprint-schema/cmd/image-builder-blueprint/ -export-toml -input testdata/valid-kernel.in.yaml 
```

Output (TOML):

```toml
name = "Blueprint example: kernel"
version = "408.5848.48376"

[customizations]
  [customizations.kernel]
    name = "kernel-debug-6.11.5-300"
    append = "nosmt=force crashkernel=1G-4G:192M,4G-64G:256M,64G-:512M"
```

### Parsing blueprints

Blueprint types have only JSON Go struct tags because mixing them with other tags like YAML will create inconsistencies, specifically for some Go types (date, time). To prevent that, loading from YAML is done via converting to JSON first. There are several helper functions available in the package which take/return the `Blueprint` type:

* ReadYAML
* WriteYAML
* ReadJSON
* WriteJSON
* MarshalYAML
* UnmarshalYAML
* MarshalJSON
* UnmarshalJSON

Example:

```go
package main

import "github.com/osbuild/blueprint-schema/pkg/blueprint"

func main() {
    blueprint, err := schema.ReadYAML(os.Stdin)
}
```

### Validation

To validate JSON or YAML buffers, use `CompileBundledSchema` function:

```go
package main

import "github.com/osbuild/blueprint-schema/pkg/blueprint"

func main() {
    schema = blueprint.CompileBundledSchema()
    err = schema.ValidateYAML(context.Background(), buffer)
}
```

### Extensions

Various advanced validation rules do not work well with Go code generator, therefore these are kept separate in `oas/extensions` directory and are only applied to `blueprint-oas3-ext.json` bundled schema. This schema must be only used for validation purposes and not for code generation.

### Tests

To run tests against fixtures: `make test`

To regenerate fixtures: `make write-fixtures`

To print diff between two YAML files via round-trip conversion `UBP-YAML > TOML > UBP-YAML`: `make test-diff`
