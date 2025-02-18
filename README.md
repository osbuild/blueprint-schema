This repo contains an example of a possible common blueprint format for osbuild.

Feel free to create a PR if you want to propose or correct anything.

## The schema

Ideas:

* Go-native type with schema definition via Go struct tags and doc
* Easy to work with in Python
* Generation for multiple variants: crc and cli
* Supported fields: name, type, description, default, required, deprecated (as only, variant name)
* Additional fields: minLength/maxLength/pattern for strings, minimum/maximum for numbers
* Description read from Go documentation optionally
* Examples rendered into OpenAPI and used in tests
* Generator and validator for OpenAPI 3.0 based on kin-openapi
* Generator and validator for OpenAPI 3.1 based on libopenapi
* Validation function that covers use cases where schema will not help (e.g. firewall port or from/to)
* Conversion tool from/to TOML
* Makes use of `go.work` so no dependencies for the main Go structure
* Works with YAML and JSON payload documents
* Flag "deprecate" that can be set to "true" or "crc" or "cli"
