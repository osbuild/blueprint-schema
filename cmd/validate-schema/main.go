package main

import (
	"flag"
	"os"

	"github.com/osbuild/blueprint-schema"
)

func main() {
	quiet := flag.Bool("quiet", false, "do not print details to the output")
	json := flag.Bool("json", false, "input is JSON (default: YAML)")

	flag.Parse()

	schema, err := blueprint.CompileSchema()
	if err != nil {
		panic(err)
	}

	if *json {
		err = schema.ReadAndValidateJSON(os.Stdin)
	} else {
		err = schema.ReadAndValidateYAML(os.Stdin)
	}

	if !*quiet && err != nil {
		os.Stdout.WriteString(err.Error())
	}

	if err != nil {
		os.Exit(1)
	}
}
