module github.com/osbuild/blueprint-schema/cmd/meta-validator

go 1.23

toolchain go1.23.6

replace github.com/osbuild/blueprint-schema => ../..

require github.com/santhosh-tekuri/jsonschema/v6 v6.0.1

require golang.org/x/text v0.19.0 // indirect
