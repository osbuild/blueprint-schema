package main

import (
	"context"
	"flag"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/osbuild/blueprint-schema"
	"gopkg.in/yaml.v3"
)

func main() {
	input := flag.String("input", "", "input file (defaults to standard input)")
	printJSONSchema := flag.Bool("print-json-schema", false, "print embedded schema to standard output and exit")
	printJSONExtendedSchema := flag.Bool("print-json-extended-schema", false, "print embedded schema to standard output and exit")
	printYAMLSchema := flag.Bool("print-yaml-schema", false, "print embedded schema to standard output and exit")
	//validateJSON := flag.Bool("validate-json", false, "validate JSON standard input")
	//validateYAML := flag.Bool("validate-yaml", false, "validate YAML standard input (default behavior)")
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

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return blueprint.SchemaFS.ReadFile(filepath.Join("oas", uri.Path))
	}

	location, _ := url.Parse(".")
	doc, err := loader.LoadFromDataWithPath(blueprint.Schema(), location)
	if err != nil {
		panic(err)
	}

	if err = doc.Validate(loader.Context); err != nil {
		panic(err)
	}

	if *printJSONSchema || *printYAMLSchema || *printJSONExtendedSchema {
		doc.InternalizeRefs(context.Background(), func(s *openapi3.T, c openapi3.ComponentRef) string {
			return strings.TrimSuffix(c.RefString(), ".yaml")
		})

		if *printYAMLSchema {
			buf, err := yaml.Marshal(doc)
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(buf)
		} else if *printJSONSchema {
			buf, err := doc.MarshalJSON()
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
			doc.Components.Schemas["fsnode"].Value.AnyOf = []*openapi3.SchemaRef{
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

			buf, err := doc.MarshalJSON()
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(buf)
		}
		return
	}

	_ = inBuf
}
