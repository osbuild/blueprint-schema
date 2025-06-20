package parse

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadYAMLWriteJSON(t *testing.T) {
	in := "name: test\n"
	want := "{\n  \"name\": \"test\"\n}"

	b, err := ReadYAML(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := marshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}

	out, err = MarshalYAML(b)
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

	b, err := readJSON(bytes.NewBufferString(in))
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatal("Expected data, got nil")
	}

	if b.Name != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name)
	}

	out, err := MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}

	out, err = marshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), in) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}

	err = writeJSON(b, bytes.NewBufferString(""), true)
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

	out, err := ConvertJSONtoYAML([]byte(in))
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

	out, err := ConvertYAMLtoJSON([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}
}
