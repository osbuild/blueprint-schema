package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/invopop/jsonschema"
	blueprint "github.com/osbuild/blueprint-schema"
	strcase "github.com/stoewer/go-strcase"
)

func rewrapText(text string) string {
	r := bufio.NewReader(bytes.NewBufferString(text))
	var wt bytes.Buffer
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, bufio.ErrFinalToken) || errors.Is(err, io.EOF) {
				wt.WriteString(line)
				break
			} else {
				panic(err)
			}
		}

		peek, err := r.Peek(1)
		if err != nil {
			if errors.Is(err, bufio.ErrFinalToken) || errors.Is(err, io.EOF) {
				wt.WriteString(line)
				break
			} else {
				panic(err)
			}
		}

		if strings.HasSuffix(line, "\n") && peek[0] == '\n' {
			wt.WriteString(line)
			wt.WriteByte('\n')
			r.ReadByte()
		} else if strings.HasSuffix(line, "\n") {
			wt.WriteString(line[:len(line)-1])
			wt.WriteByte(' ')
		} else {
			wt.WriteString(line)
		}
	}
	result := wt.String()
	if strings.HasSuffix(result, "\n") {
		return result
	} else if result != "" {
		return result + "\n"
	} else {
		return result
	}
}

func main() {
	pkgPath := flag.String("src-path", ".", "path to Go source package with structs to reflect")

	flag.Parse()

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase
	r.ExpandedStruct = true

	if _, err := os.Stat(filepath.Join(*pkgPath, "/blueprint.go")); errors.Is(err, os.ErrNotExist) {
		panic("must be run from the root of the project in order to load Go comments via Go AST parser")
	}
	if err := r.AddGoComments("github.com/osbuild/blueprint-schema", ".", jsonschema.WithFullComment()); err != nil {
		panic(err)
	}

	for k, v := range r.CommentMap {
		r.CommentMap[k] = rewrapText(v)
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
