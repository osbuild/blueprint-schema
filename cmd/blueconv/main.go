//go:build !js
// +build !js

package main

import (
	"flag"
	"os"

	"github.com/osbuild/blueprint-schema"
)

func main() {
	quiet := flag.Bool("quiet", false, "do not print details to the output")
	validateJSON := flag.Bool("validate-json", false, "validate JSON input")
	validateYAML := flag.Bool("validate-yaml", false, "validate YAML input (default behavior)")

	flag.Parse()

	schema, err := blueprint.CompileSchema()
	if err != nil {
		panic(err)
	}

	if *validateJSON {
		err = schema.ReadAndValidateJSON(os.Stdin)
	} else if *validateYAML {
		err = schema.ReadAndValidateYAML(os.Stdin)
	} else {
		// default behavior
		err = schema.ReadAndValidateYAML(os.Stdin)
	}

	if !*quiet && err != nil {
		os.Stdout.WriteString(err.Error())
	}

	if err != nil {
		os.Exit(1)
	}
}
