package blueprint_test

import (
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	blueprint "github.com/lzap/common-blueprint-example"
)

func TestEmbeddedSchema(t *testing.T) {
	file, err := os.Open("blueprint-schema.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	schema, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(schema, blueprint.SchemaJSON); diff != "" {
		t.Errorf("Embedded schema differs: %s", diff)
	}
}
