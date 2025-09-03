package parse

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadYAMLWriteJSON(t *testing.T) {
	in := "name: test\n"
	wantIndent := "{\n\t\"name\": \"test\"\n}"
	want := `{"name":"test"}`

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

	out, err := MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), wantIndent); diff != "" {
		t.Fatalf("Unexpected output: %s", diff)
	}

	out, err = MarshalJSON(b, false)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), want); diff != "" {
		t.Fatalf("Unexpected output: %s", diff)
	}

	out, err = MarshalYAML(b)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), in); diff != "" {
		t.Fatalf("Unexpected output: %s", diff)
	}
}

func TestReadJSONWriteYAML(t *testing.T) {
	in := "{\n\t\"name\": \"test\"\n}"
	want := "name: test\n"

	b, err := ReadJSON(bytes.NewBufferString(in))
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

	if diff := cmp.Diff(string(out), want); diff != "" {
		t.Fatalf("Unexpected YAML output: %s", diff)
	}

	out, err = MarshalJSON(b, true)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(out), in); diff != "" {
		t.Fatalf("Unexpected JSON output: %s", diff)
	}
}

func TestConvertJSONtoYAML(t *testing.T) {
	in := "{\n  \"name_test\": \"test\"\n}"
	want := "name_test: test\n"

	out, err := ConvertJSONtoYAML([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected YAML output: %s", out)
	}
}

func TestConvertYAMLtoJSON(t *testing.T) {
	in := "name_test: test\n"
	want := "{\"name_test\":\"test\"}"

	out, err := ConvertYAMLtoJSON([]byte(in))
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Diff(string(out), want) != "" {
		t.Fatalf("Unexpected JSON output: %s", out)
	}
}
