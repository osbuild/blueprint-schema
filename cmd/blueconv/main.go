//go:build !js
// +build !js

package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/blueprint"
	"github.com/osbuild/blueprint-schema/pkg/conv/notes"
	"github.com/osbuild/blueprint-schema/pkg/conv/onprem"
)

func main() {
	quiet := flag.Bool("quiet", false, "do not print details to the output")
	input := flag.String("input", "", "input file (defaults to standard input)")
	validateJSON := flag.Bool("validate-json", false, "validate JSON standard input")
	validateYAML := flag.Bool("validate-yaml", false, "validate YAML standard input (default behavior)")
	convert := flag.Bool("convert", false, "convert standard input to standard output")
	inputFormat := flag.String("input-format", "yaml", "input format (json or yaml - default is yaml)")
	outputFormat := flag.String("output-format", "toml", "output format (toml)")

	flag.Parse()

	in := os.Stdin
	if *input != "" {
		var err error
		in, err = os.Open(*input)
		if err != nil {
			panic(err)
		}
		defer in.Close()
	}

	schema, schemaErr := blueprint.CompileSchema()
	if schemaErr != nil {
		panic(schemaErr)
	}

	if *convert {
		var from *blueprint.Blueprint
		var err error

		if *inputFormat == "json" {
			from, err = blueprint.ReadJSON(in)
			if err != nil {
				panic(err)
			}
		} else if *inputFormat == "yaml" {
			from, err = blueprint.ReadYAML(in)
			if err != nil {
				panic(err)
			}
		} else {
			panic("invalid input format")
		}

		schemaErr = schema.Validate(from)
		if schemaErr != nil {
			os.Stdout.WriteString(schemaErr.Error())
			os.Stdout.WriteString("\n")
			os.Exit(1)
		}

		errs := notes.ConversionNotes{}
		if *outputFormat == "toml" {
			to := onprem.ExportBlueprint(from, &errs)
			err = toml.NewEncoder(os.Stdout).Encode(to)
			if err != nil {
				panic(err)
			}
		} else {
			panic("invalid output format")
		}

		if !*quiet {
			os.Stdout.WriteString("\n")
			for _, err := range errs.Errors() {
				os.Stdout.WriteString(err.Error())
				os.Stdout.WriteString("\n")
			}
		}
	} else if *validateJSON {
		schemaErr = schema.ReadAndValidateJSON(in)
	} else if *validateYAML {
		schemaErr = schema.ReadAndValidateYAML(in)
	} else {
		// default behavior
		schemaErr = schema.ReadAndValidateYAML(in)
	}

	if !*quiet && schemaErr != nil {
		os.Stdout.WriteString(schemaErr.Error())
		os.Stdout.WriteString("\n")
	}

	if schemaErr != nil {
		os.Exit(1)
	}
}
