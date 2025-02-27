//go:build !js
// +build !js

package blueprint_test

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	blueprint "github.com/osbuild/blueprint-schema"
)

func TestReadYAMLWriteJSON(t *testing.T) {
	in := "name: test\n"
	want := "{\n  \"name\": \"test\"\n}"

	b, err := blueprint.ReadYAML(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := blueprint.MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}

	out, err = blueprint.MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), in) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}
}

func TestReadJSONWriteYAML(t *testing.T) {
	in := "{\n  \"name\": \"test\"\n}"
	want := "name: test\n"

	b, err := blueprint.ReadJSON(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := blueprint.MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}

	out, err = blueprint.MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), in) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}
}

func TestConvertJSONtoYAML(t *testing.T) {
	in := "{\n  \"name\": \"test\"\n}"
	want := "name: test\n"

	out, err := blueprint.ConvertJSONtoYAML([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}
}

func TestConvertYAMLtoJSON(t *testing.T) {
	in := "name: test\n"
	want := "{\"name\":\"test\"}"

	out, err := blueprint.ConvertYAMLtoJSON([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}
}
