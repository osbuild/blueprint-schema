package blueprint

import (
	"embed"
	"io"
)

//go:embed oas
var SchemaFS embed.FS

//go:embed blueprint-oas3*
var BundledSchemaFS embed.FS

// SchemaSource returns the schema source schema as defined in oas/ directory.
func SchemaSource() []byte {
	buf, err := SchemaFS.Open("oas/blueprint-oas.yaml")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = buf.Close()
	}()

	schema, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}

	return schema
}

func BundledSchema() []byte {
	buf, err := BundledSchemaFS.Open("blueprint-oas3-ext.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = buf.Close()
	}()

	schema, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}

	return schema
}
