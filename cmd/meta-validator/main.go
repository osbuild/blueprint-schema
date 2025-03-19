package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	//go:embed oas3-2024-10-18.json
	oas3 []byte

	//go:embed draft5.json
	draft5 []byte
)

func main() {
	input := flag.String("input", "", "input file (defaults to standard input)")
	inSchema := flag.String("schema", "oas3", "meta schema (oas3 or draft5)")

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

	var jsonSchema any
	var err error
	if inSchema == nil || *inSchema == "oas3" {
		jsonSchema, err = jsonschema.UnmarshalJSON(bytes.NewBuffer(oas3))
	} else if *inSchema == "draft5" {
		jsonSchema, err = jsonschema.UnmarshalJSON(bytes.NewBuffer(draft5))
	} else {
		panic("invalid schema")
	}
	if err != nil {
		panic(err)
	}

	compiler := jsonschema.NewCompiler()
	compiler.AddResource("meta-schema.json", jsonSchema)
	schema, err := compiler.Compile("meta-schema.json")
	if err != nil {
		panic(err)
	}

	doc, err := jsonschema.UnmarshalJSON(in)
	if err != nil {
		panic(err)
	}

	err = schema.Validate(doc)
	if err != nil {
		fmt.Println(err)
	}
}
