package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/parse"
)

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}

func main() {
	var inBuf []byte
	var err error

	ctx := context.Background()
	printJSONSchema := flag.Bool("print-json-schema", false, "print embedded schema to standard output and exit")
	printJSONExtendedSchema := flag.Bool("print-json-extended-schema", false, "print embedded schema to standard output and exit")
	printYAMLSchema := flag.Bool("print-yaml-schema", false, "print embedded schema to standard output and exit")
	validate := flag.Bool("validate", false, "validate input document (detects JSON or YAML format)")
	exportTOML := flag.Bool("export-toml", false, "convert document from UBP YAML/JSON to BP TOML")
	exportJSON := flag.Bool("export-json", false, "convert document from UBP YAML/JSON to BP JSON")
	importYAML := flag.Bool("import-yaml", false, "convert document from BP TOML/JSON to UBP YAML")
	importJSON := flag.Bool("import-json", false, "convert document from BP TOML/JSON to UBP JSON")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [input] [file...]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "If no input file is specified, the program reads from standard input.")
		fmt.Fprintln(os.Stderr, "The input file format is detected automatically (JSON or YAML).")
	}
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{""}
	}

	for _, arg := range args {
		in := os.Stdin
		if arg != "" {
			in, err = os.Open(arg)
			if err != nil {
				die("%s: %s", arg, err)
			}
			defer func() {
				_ = in.Close()
			}()
		}

		schema, err := parse.CompileSourceSchema()
		if err != nil {
			die("%s: %s", arg, err)
		}

		err = schema.ValidateSchema(ctx)
		if err != nil {
			die("%s: %s", arg, err)
		}

		if *printJSONSchema || *printYAMLSchema || *printJSONExtendedSchema {
			err = schema.Bundle(ctx)
			if err != nil {
				die("%s: %s", arg, err)
			}

			if *printYAMLSchema {
				buf, err := schema.MarshalYAML()
				if err != nil {
					die("%s: %s", arg, err)
				}

				_, _ = os.Stdout.Write(buf)
			} else if *printJSONSchema {
				buf, err := schema.MarshalJSON()
				if err != nil {
					die("%s: %s", arg, err)
				}

				_, _ = os.Stdout.Write(buf)
			} else if *printJSONExtendedSchema {
				err := schema.ApplyExtensions(ctx)
				if err != nil {
					die("%s: %s", arg, err)
				}

				buf, err := schema.MarshalJSON()
				if err != nil {
					die("%s: %s", arg, err)
				}

				_, _ = os.Stdout.Write(buf)
			}

			return
		} else if *validate {
			schema, err = parse.CompileBundledSchema()
			if err != nil {
				die("%s: %s", arg, err)
			}

			inBuf, err = io.ReadAll(in)
			if err != nil {
				die("%s: %s", arg, err)
			}

			err = schema.ValidateAny(ctx, inBuf)
			if err != nil {
				die("%s: %s", arg, err)
			}
		} else if *exportTOML || *exportJSON {
			inBuf, err = io.ReadAll(in)
			if err != nil {
				die("%s: %s", arg, err)
			}

			b, err := parse.UnmarshalYAML(inBuf)
			if err != nil {
				die("%s: %s", arg, err)
			}

			exporter := conv.NewInternalExporter(b)
			result, logs := exporter.Export()
			if logs != nil {
				fmt.Fprintln(os.Stderr, logs)
			}

			var buf []byte
			if *exportJSON {
				buf, err = json.MarshalIndent(result, "", "  ")
			} else if *exportTOML {
				buf, err = toml.Marshal(result)
			}
			if err != nil {
				die("%s: %s", arg, err)
			}
			_, _ = os.Stdout.Write(buf)
		} else if *importYAML || *importJSON {
			inBuf, err = io.ReadAll(in)
			if err != nil {
				die("%s: %s", arg, err)
			}

			details := parse.AnyDetails{}
			b, err := parse.UnmarshalAny(inBuf, &details)
			if err != nil {
				die("%s: %s", arg, err)
			}

			if details.Warnings != nil {
				fmt.Fprintln(os.Stderr, details.Warnings)
			}

			var buf []byte
			if *importJSON {
				buf, err = parse.MarshalJSON(b, true)
			} else if *importYAML {
				buf, err = parse.MarshalYAML(b)
			}
			if err != nil {
				die("%s: %s", arg, err)
			}
			_, _ = os.Stdout.Write(buf)
		}
	}
}
