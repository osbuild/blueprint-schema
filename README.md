## Blueprint schema

**WORK IN PROGRESS** but please send us feedback.

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

### Schema files

* `blueprint-oas3.yaml` - OpenAPI 3.0 schema, the blueprint document is at `#/components/schemas/blueprint`
* `blueprint-oas3.json` - the same schema but in JSON format
* `blueprint-oas3-ext.json` - the schema with additional "extensions" which would confuse Go code generator but are useful for validation purposes, they are defined in `oas/extensions`

All the mentioned files are generated using `make schema`

### Go code

All the code resides in `pkg/blueprint` except embedded schemas from above which are in the top-level directory for technical reasons (Go embedding limitations). Direct access to schema files is not required for any scenario, so only import the former package:

```go
package main

import "github.com/osbuild/blueprint-schema/pkg/blueprint"

func main() {
    var b Blueprint
}
```

### WIP remarks

* We have decided to utilize OAS3 schema instead of JSON Schema to keep OpenAPI 3.0 compatibility
* The validation library `kin-openapi` does not guarantee order of hash keys therefore all outputs have random property order
* Additional checks can be done via `ext` schema

### TODO

* Import/export
* Review documentation, fix godoc rendering and prefix
* TinyGO WASM validator
* https://github.com/osbuild/blueprint-schema/tree/3e99c4560011c30ca75ef6e2df3636404147f667
