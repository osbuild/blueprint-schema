package blueprint

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
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

	if b.Name.Get() != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name.Get())
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

	if b.Name.Get() != "test" {
		t.Fatalf("Expected 'test', got '%s'", b.Name.Get())
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

func TestUnmarshalYAMLMultipleDocuments(t *testing.T) {
	tc := []struct {
		name        string
		in          string
		want        string
		wantErr     bool
		errContains string
	}{
		{
			name: "simple",
			in: `
name: test1
---
name: test2
`,
			want: `
name: test2
`,
		},
		{
			name: "three",
			in: `
name: test1
---
name: test2
---
name: test3
`,
			want: `
name: test3
`,
		},
		{
			name: "two unique",
			in: `
name: test1
description: "desc1"
---
name: test2
`,
			want: `
name: test2
description: "desc1"
`,
		},
		{
			name: "tail",
			in: `
name: test1
---
`,
			want: `
name: test1
`,
		},
		{
			name: "with header",
			in: `
---
name: test1
---
name: test2
`,
			want: `
name: test2
`,
		},
		{
			name: "with header and dots",
			in: `
---
name: test1
...
---
name: test2
`,
			want: `
name: test2
`,
		},
		{
			name: "empty string",
			in: `
name: test1
---
name: ""
`,
			want: `
name: "test1"
`,
		},
		{
			name: "slices",
			in: `
accounts:
  users:
    - name: "user1"
---
accounts:
  users:
    - name: "user2"
`,
			want: `
accounts:
  users:
    - name: "user1"
    - name: "user2"
`,
		},
		{
			name: "slices with empty",
			in: `
accounts:
  users:
    - name: "user1"
---
accounts:
  users:
`,
			want: `
accounts:
  users:
    - name: "user1"
`,
		},
		{
			name: "bool positive",
			in: `
fips:
  enabled: false
---
fips:
  enabled: true
`,
			want: `
fips:
  enabled: true
`,
		},
		{
			name: "bool negative",
			in: `
fips:
  enabled: true
---
fips:
  enabled: false
`,
			want: `
fips:
  enabled: true # TODO: this is confusing pointer must be used for all bools, strings and numbers
`,
		},
		{
			name: "bool with empty",
			in: `
fips:
  enabled: true
---
fips:
`,
			want: `
fips:
  enabled: true
`,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got, err := ReadYAML(bytes.NewBufferString(c.in))
			if err != nil {
				t.Fatal(err)
			}

			want, err := ReadYAML(bytes.NewBufferString(c.want))
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(*want, *got, cmp.AllowUnexported(Blueprint{}), cmpopts.EquateComparable(ptr.Ref[string]{})) {
				t.Fatalf("Unexpected data: %s", cmp.Diff(*want, *got, cmp.AllowUnexported(Blueprint{}), cmpopts.EquateComparable(ptr.Ref[string]{})))
			}
		})
	}
}
