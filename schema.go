package blueprint

import (
	"embed"
	"io"
)

//go:embed oas
var SchemaFS embed.FS

func Schema() []byte {
	buf, err := SchemaFS.Open("oas/blueprint-oas.yaml")
	if err != nil {
		panic(err)
	}
	defer buf.Close()

	schema, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}
	return schema
}
