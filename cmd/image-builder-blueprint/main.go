package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gabriel-vasile/mimetype"
	"github.com/osbuild/blueprint-schema/pkg/blueprint"
	"github.com/osbuild/blueprint-schema/pkg/convert"
)

func main() {
	ctx := context.Background()
	input := flag.String("input", "", "input JSON or YAML file (defaults to standard input, detects format)")
	printJSONSchema := flag.Bool("print-json-schema", false, "print embedded schema to standard output and exit")
	printJSONExtendedSchema := flag.Bool("print-json-extended-schema", false, "print embedded schema to standard output and exit")
	printYAMLSchema := flag.Bool("print-yaml-schema", false, "print embedded schema to standard output and exit")
	validate := flag.Bool("validate", false, "validate input document (detects JSON or YAML format)")
	exportTOML := flag.Bool("export-toml", false, "convert document into legacy TOML")
	flag.Parse()

	convert.SetLogger(log.Default())

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

	schema, err := blueprint.CompileSourceSchema()
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
		schema, err = blueprint.CompileBundledSchema()
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
		} else if mime.Is("application/x-yaml") || mime.Is("text/yaml") {
			err = schema.ValidateYAML(ctx, inBuf)
		} else {
			err = errors.New("unsupported format, only JSON and YAML are supported")
		}
		if err != nil {
			panic(err)
		}

	} else if exportTOML != nil {
		inBuf, err = io.ReadAll(in)
		if err != nil {
			panic(err)
		}

		b, err := blueprint.UnmarshalYAML(inBuf)
		if err != nil {
			panic(err)
		}
		eb := convert.ExportBlueprint(b)
		buf, err := toml.Marshal(eb)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(buf)
	}

	_ = inBuf
}
