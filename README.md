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

At the end of 2024 we held a series of meetings where the core team agreed on a common blueprint schema which is based on OAS3 specification (OpenAPI 3.0) with generated Go code.

### WIP

This is work in progress, more documentation will follow.

### TODO

* Import/export
* Review documentation, fix godoc rendering and prefix
* TinyGO WASM validator
* https://github.com/osbuild/blueprint-schema/tree/3e99c4560011c30ca75ef6e2df3636404147f667
