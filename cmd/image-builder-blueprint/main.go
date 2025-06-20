package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gabriel-vasile/mimetype"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/parse"
)

func main() {
	ctx := context.Background()
	input := flag.String("input", "", "input JSON or YAML file (defaults to standard input, detects format)")
	printJSONSchema := flag.Bool("print-json-schema", false, "print embedded schema to standard output and exit")
	printJSONExtendedSchema := flag.Bool("print-json-extended-schema", false, "print embedded schema to standard output and exit")
	printYAMLSchema := flag.Bool("print-yaml-schema", false, "print embedded schema to standard output and exit")
	validate := flag.Bool("validate", false, "validate input document (detects JSON or YAML format)")
	exportTOML := flag.Bool("export-toml", false, "convert document into legacy TOML")
	exportJSON := flag.Bool("export-json", false, "convert document into legacy JSON")
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

	schema, err := parse.CompileSourceSchema()
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
			err := schema.ApplyExtensions(ctx)
			if err != nil {
				panic(err)
			}

			buf, err := schema.MarshalJSON()
			if err != nil {
				panic(err)
			}

			os.Stdout.Write(buf)
		}

		return
	} else if *validate {
		schema, err = parse.CompileBundledSchema()
		if err != nil {
			panic(err)
		}

		inBuf, err = io.ReadAll(in)
		if err != nil {
			panic(err)
		}

		mime := mimetype.Detect(inBuf)
		if mime.Is("application/json") {
			err = schema.ValidateJSON(ctx, inBuf)
		} else if mime.Is("application/x-yaml") || mime.Is("text/yaml") || mime.Is("text/plain") {
			err = schema.ValidateYAML(ctx, inBuf)
		} else {
			err = fmt.Errorf("unsupported format: %s, only JSON and YAML are supported", mime.String())
		}
		if err != nil {
			panic(err)
		}

	} else if *exportTOML || *exportJSON {
		inBuf, err = io.ReadAll(in)
		if err != nil {
			panic(err)
		}

		b, err := parse.UnmarshalYAML(inBuf)
		if err != nil {
			panic(err)
		}

		exporter := conv.NewInternalExporter(b)
		if logs := exporter.Export(); logs != nil {
			fmt.Fprintln(os.Stderr, logs)
		}

		var buf []byte
		if *exportJSON {
			buf, err = json.MarshalIndent(exporter.Result(), "", "  ")
		} else if *exportTOML {
			buf, err = toml.Marshal(exporter.Result())
		}
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(buf)
	}

	_ = inBuf
}
