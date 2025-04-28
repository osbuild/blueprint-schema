## Blueprint schema

**WORK IN PROGRESS** but please send us feedback.

This repository contains the common blueprint JSON Schema and Go types for Image Builder / osbuild customization configuration.

[![Go Reference](https://pkg.go.dev/badge/github.com/osbuild/blueprint-schema.svg)](https://pkg.go.dev/github.com/osbuild/blueprint-schema)

### WIP remarks

* We have decided to utilize OAS3 schema instead of JSON Schema to keep OpenAPI 3.0 compatibility
* The validation library `kin-openapi` does not guarantee order of hash keys therefore all outputs have random property order
* Additional checks can be done via `ext` schema

### TODO

* Import/export
* Review documentation, fix godoc rendering and prefix
* TinyGO WASM validator
* https://github.com/osbuild/blueprint-schema/tree/3e99c4560011c30ca75ef6e2df3636404147f667
