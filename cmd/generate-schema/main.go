package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	blueprint "github.com/lzap/common-blueprint-example"
	strcase "github.com/stoewer/go-strcase"
)

func main() {
	pkgPath := flag.String("src-path", ".", "path to Go source package with structs to reflect")

	flag.Parse()

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	r.ExpandedStruct = true

	if _, err := os.Stat(filepath.Join(*pkgPath, "/blueprint.go")); errors.Is(err, os.ErrNotExist) {
		panic("must be run from the root of the project in order to load Go comments via Go AST parser")
	}
	if err := r.AddGoComments("github.com/lzap/common-blueprint-example", ".", jsonschema.WithFullComment()); err != nil {
		panic(err)
	}

	schema := r.Reflect(&blueprint.Blueprint{})

	minimizedSchema, err := schema.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var prettySchema bytes.Buffer
	err = json.Indent(&prettySchema, minimizedSchema, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(prettySchema.String())
}
