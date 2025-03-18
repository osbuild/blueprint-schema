package blueprint

import (
	"embed"
	"io"

	"github.com/invopop/jsonschema"
)

//go:embed blueprint_*.yaml
var partialSchemaFS embed.FS

var partialSchemaMap map[string]*jsonschema.Schema = make(map[string]*jsonschema.Schema)

func init() {
	dir, err := partialSchemaFS.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		f, err := partialSchemaFS.Open(file.Name())
		if err != nil {
			panic(err)
		}
		contents, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Close()

		schema, err := buildFromYAML(contents)
		if err != nil {
			panic(err.Error() + " in file " + file.Name())
		}
		partialSchemaMap[file.Name()] = schema
	}
}

func buildFromYAML(yaml []byte) (*jsonschema.Schema, error) {
	jsonBuffer, err := ConvertYAMLtoJSON(yaml)
	if err != nil {
		return nil, err
	}

	schema := &jsonschema.Schema{}
	err = schema.UnmarshalJSON(jsonBuffer)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

// PartialSchema returns a JSON schema loaded from `blueprint_*.yaml` file. These files are embedded
// into the binary and can be accessed using this function. They provide an alternative method of
// defining JSON schema for a struct when Go struct tags are not sufficient.
func PartialSchema(name string) *jsonschema.Schema {
	return partialSchemaMap[name]
}
