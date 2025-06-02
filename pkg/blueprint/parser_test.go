package blueprint

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestUnmarshalYAMLMultipleDocuments(t *testing.T) {
	tc := []struct {
		name        string
		in          string
		want        string
		errContains string
	}{
		{
			name: "Unsupported field",
			in: `
name: "Unsupported"
description: "Document one"
---
description: "Document one"
registration:
`,
			errContains: `cannot merge field into blueprint: ["Description"]`,
		},
		{
			name: "Registration",
			in: `
name: "Registration"
---
registration:
  redhat:
    activation_key: "123456789"
    organization: "123456"
    subscription_manager:
      enabled: true
`,
			want: `
name: "Registration"
registration:
  redhat:
    activation_key: "123456789"
    organization: "123456"
    subscription_manager:
      enabled: true
`,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got, err := ReadYAML(bytes.NewBufferString(c.in))

			if c.errContains != "" {
				if err == nil {
					t.Fatalf("Expected error containing %q, got nil", c.errContains)
				}

				if !bytes.Contains([]byte(err.Error()), []byte(c.errContains)) {
					t.Fatalf("Expected error containing %q, got %q", c.errContains, err.Error())
				}

				return
			} else if err != nil {
				t.Fatal(err)
			}

			want, err := ReadYAML(bytes.NewBufferString(c.want))
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(*want, *got) {
				t.Fatalf("Unexpected data: %s", cmp.Diff(*want, *got))
			}
		})
	}
}

func TestNonZeroFields(t *testing.T) {
	tests := []struct {
		name          string
		input         *Blueprint
		expectedNames []string
	}{
		{
			name: "Name",
			input: &Blueprint{
				Name: "name",
			},
			expectedNames: []string{"Name"},
		},
		{
			name: "Name and Description",
			input: &Blueprint{
				Name:        "name",
				Description: "description",
			},
			expectedNames: []string{"Name", "Description"},
		},
		{
			name: "Slice",
			input: &Blueprint{
				Containers: []Container{
					{
						Name: "cont",
					},
				},
			},
			expectedNames: []string{"Containers"},
		},
		{
			name: "Slice",
			input: &Blueprint{
				FIPS: &FIPS{
					Enabled: true,
				},
			},
			expectedNames: []string{"FIPS"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames := nonZeroFields(tt.input)

			// Use cmpopts.SortSlices to handle different orderings if field iteration order is not guaranteed
			// (though for struct fields it's typically definition order).
			sorter := cmpopts.SortSlices(func(x, y string) bool { return x < y })
			if !cmp.Equal(tt.expectedNames, gotNames, sorter) {
				t.Errorf("nonZeroFields() got = %v, want %v, diff: %s", gotNames, tt.expectedNames, cmp.Diff(tt.expectedNames, gotNames, sorter))
			}
		})
	}
}
