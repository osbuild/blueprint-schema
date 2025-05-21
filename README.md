## Blueprint schema

**WORK IN PROGRESS** but please send us feedback.

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

### Terminology

* UBP: Unified Blueprint (this repo). Also known as Blueprint v2.
* Blueprint: The blueprint as it exists in osbuild/blueprint, is user-facing in osbuild-composer, and is documented on osbuild.org
* CRC Blueprint: The blueprint as it exists in image-builder-crc, is user-facing in the service API, and is documented on osbuild.org.  This is very close but sometimes slightly different from the Blueprint.
* Images Blueprint: The format and code that currently exists in osbuild/images/pkg/blueprint, which was never meant to be user-facing, but through sloppiness (on my part as well) initially (but not anymore) ended up being the user-facing blueprint in bootc-image-builder and, for a short period, ib-cli.  This slightly differs from the Old Schema in that it was more aggressive in dropping deprecated options (like SSHKey), because the user-facing blueprints (Old and CRC) were responsible for backwards compatibility.

### Schema files

* `blueprint-oas3.yaml` - OpenAPI 3.0 schema, the blueprint document is at `#/components/schemas/blueprint`
* `blueprint-oas3.json` - the same schema but in JSON format
* `blueprint-oas3-ext.json` - the schema with additional "extensions" which would confuse Go code generator but are useful for validation purposes, they are defined in `oas/extensions`

All the mentioned files are generated using `make schema`

### Schema source

All schema source files are in `oas/` directory, each component resides in its own YAML file in `oas/components`. Make sure to create component for each object that is supposed to be a Go type (`struct`). There is a README in the `oas/` directory with some tips on how to write OAS3 schemas.

### Go code

Go code is generated from `blueprint-oas3.json` via `oapi-codegen` using `make schema`.

All the code resides in `pkg/blueprint` except embedded schemas from above which are in the top-level directory for technical reasons (Go embedding limitations). Direct access to schema files is not required for any scenario, so only import the former package.

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
