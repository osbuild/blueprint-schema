package main

import (
	"flag"
	"os"

	"github.com/osbuild/blueprint-schema/validate"
)

func main() {
	quiet := flag.Bool("quiet", false, "do not print details to the output")
	json := flag.Bool("json", false, "input is JSON (default: YAML)")

	flag.Parse()

	schema, err := validate.CompileSchema()
	if err != nil {
		panic(err)
	}

	var valid bool
	var out string

	if *json {
		valid, out, err = schema.ReadAndValidateJSON(os.Stdin)
	} else {
		valid, out, err = schema.ReadAndValidateYAML(os.Stdin)
	}
	if err != nil {
		panic(err)
	}

	if !*quiet {
		os.Stdout.WriteString(out)
	}

	if !valid {
		os.Exit(1)
	}
}
