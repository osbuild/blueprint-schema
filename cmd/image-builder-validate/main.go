package main

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/osbuild/blueprint-schema/pkg/blueprint"
)

func main() {
	ctx := context.Background()
	input := flag.String("input", "", "input file (defaults to standard input)")
	printJSONSchema := flag.Bool("print-json-schema", false, "print embedded schema to standard output and exit")
	printJSONExtendedSchema := flag.Bool("print-json-extended-schema", false, "print embedded schema to standard output and exit")
	printYAMLSchema := flag.Bool("print-yaml-schema", false, "print embedded schema to standard output and exit")
	validateJSON := flag.Bool("validate-json", false, "validate JSON standard input")
	validateYAML := flag.Bool("validate-yaml", false, "validate YAML standard input (default behavior)")
	flag.Parse()

	var inBuf []byte
	var err error

	in := os.Stdin
	if *input != "" {
		in, err = os.Open(*input)
		if err != nil {
			panic(err)
		}
		defer in.Close()
	}

	if !*printJSONSchema && !*printYAMLSchema && !*printJSONExtendedSchema {
		inBuf, err = io.ReadAll(in)
		if err != nil {
			panic(err)
		}
	}

	schema, err := blueprint.CompileSchema()
	if err != nil {
		panic(err)
	}

	err = schema.ValidateSchema(ctx)
	if err != nil {
		panic(err)
	}

	if *printJSONSchema || *printYAMLSchema || *printJSONExtendedSchema {
		err = schema.Bundle(ctx)
		if err != nil {
			panic(err)
		}

		if *printYAMLSchema {
			buf, err := schema.MarshalYAML()
			if err != nil {
				panic(err)
			}

			os.Stdout.Write(buf)
		} else if *printJSONSchema {
			buf, err := schema.MarshalJSON()
			if err != nil {
				panic(err)
			}

			os.Stdout.Write(buf)
		} else if *printJSONExtendedSchema {
			// fsnodes: if type is "dir", contents must not be present
			//
			// anyOf:
			//   - not:
			//       properties:
			//         type:
			//           enum: ["dir"]
			//       required:
			//         - type
			//   - not:
			//       required:
			//       - contents
			schema.Document().Components.Schemas["fsnode"].Value.AnyOf = []*openapi3.SchemaRef{
				{
					Value: &openapi3.Schema{
						Not: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Properties: openapi3.Schemas{
									"type": &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Enum: []any{"dir"},
										},
									},
								},
								Required: []string{"type"},
							},
						},
					},
				},
				{
					Value: &openapi3.Schema{
						Not: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Required: []string{"contents"},
							},
						},
					},
				},
			}

			buf, err := schema.MarshalJSON()
			if err != nil {
				panic(err)
			}

			os.Stdout.Write(buf)
		}

		return
	} else if *validateJSON {
		err = schema.ValidateJSON(ctx, inBuf)
		if err != nil {
			panic(err)
		}
	} else if *validateYAML {
		err = schema.ValidateYAML(ctx, inBuf)
		if err != nil {
			panic(err)
		}
	}

	_ = inBuf
}
