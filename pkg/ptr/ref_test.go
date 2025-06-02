package ptr_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

func TestIntRef(t *testing.T) {
	var zeroint ptr.Ref[int]

	if zeroint.NotZero() {
		t.Errorf("Zero[int] should not be zero")
	}

	if ptr.Zero[int]().NotZero() {
		t.Errorf("Zero[int]() should be zero")
	}

	if ptr.New(13).IsZero() {
		t.Errorf("New(13) should not be zero")
	}

	if ptr.New(13).Get() != 13 {
		t.Errorf("New(13).Get() should return 13")
	}

	if ptr.New(13).GetOr(42) != 13 {
		t.Errorf("New(13).GetOr(42) should return 13")
	}

	if ptr.Zero[int]().GetOr(42) != 42 {
		t.Errorf("Zero[int]().GetOr(42) should return 42")
	}
}

func TestBoolRef(t *testing.T) {
	if ptr.New(true).Get() != true {
		t.Errorf("New[bool] should return true")
	}

	if ptr.New(false).Get() != false {
		t.Errorf("New[bool] should return false")
	}

	if ptr.Zero[bool]().GetOr(true) != true {
		t.Errorf("Zero[bool]().GetOr(true) should return true")
	}

	if ptr.ZeroBool().GetOr(true) != true {
		t.Errorf("ZeroBool().GetOr(true) should return true")
	}
}

func TestStringRef(t *testing.T) {
	if ptr.New("test").Get() != "test" {
		t.Errorf("New[string] should return 'test'")
	}

	if ptr.Zero[string]().GetOr("default") != "default" {
		t.Errorf("Zero[string]().GetOr('default') should return 'default'")
	}

	if ptr.ZeroString().GetOr("default") != "default" {
		t.Errorf("ZeroString().GetOr('default') should return 'default'")
	}
}

type testType struct {
	BoolTrue    ptr.Ref[bool] `json:"true" toml:"true"`
	BoolFalse   ptr.Ref[bool] `json:"false" toml:"false"`
	BoolUnset   ptr.Ref[bool] `json:"unset" toml:"unset"`
	StructSet   *allTypes     `json:"struct_set,omitempty" toml:"struct_set,omitempty"`
	StructUnset *allTypes     `json:"struct_unset,omitempty" toml:"struct_unset,omitempty"`
}

type allTypes struct {
	String ptr.Ref[string] `json:"string" toml:"string"`
	Bool   ptr.Ref[bool]   `json:"bool" toml:"bool"`
	Int    ptr.Ref[int64]  `json:"int" toml:"int"`
}

func TestUnmarshalJSON(t *testing.T) {
	testTomlStr := `
{
	"true": true,
	"false": false,
	"unset": null,
	"struct_set": {
		"string": "string",
		"bool": true,
		"int": 42
	}
}
`
	var got testType
	err := json.Unmarshal([]byte(testTomlStr), &got)
	if err != nil {
		t.Fatalf("Failed to unmarshal TOML: %v", err)
	}

	want := testType{
		BoolTrue:  ptr.New(true),
		BoolFalse: ptr.New(false),
		BoolUnset: ptr.ZeroBool(),
		StructSet: &allTypes{
			String: ptr.New("string"),
			Bool:   ptr.New(true),
			Int:    ptr.NewInt64(42),
		},
	}

	if cmp.Diff(want, got) != "" {
		t.Errorf("Unmarshaled struct does not match expected value:\n%s", cmp.Diff(want, got))
	}
}

func TestUnmarshalJSONError(t *testing.T) {
	testTomlStr := `
{
	"struct_set": {
		"string": 42,
		"bool": "string",
		"int": true
	}
}
`

	var got testType
	err := json.Unmarshal([]byte(testTomlStr), &got)
	if err == nil {
		t.Fatalf("Expected error when unmarshaling, but got none")
	}

	if !strings.Contains(err.Error(), "cannot unmarshal") {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestMarshalJSON(t *testing.T) {
	testData := testType{
		BoolTrue:  ptr.New(true),
		BoolFalse: ptr.New(false),
		BoolUnset: ptr.ZeroBool(),
	}

	want := `{"true":true,"false":false,"unset":null}`

	data, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	if string(data) != want {
		t.Errorf("Marshaled JSON does not match expected value:\nGot: %s\nWant: %s", data, want)
	}
}

func TestUnmarshalTOML(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  testType
	}{
		{
			name:  "Boolean",
			input: "true = true\nfalse = false\n",
			want: testType{
				BoolTrue:  ptr.New(true),
				BoolFalse: ptr.New(false),
			},
		},
		{
			name: "Boolean with null",
			input: `
[struct_set]
string = "string"
bool = true
int = 42
`,
			want: testType{
				StructSet: &allTypes{
					String: ptr.New("string"),
					Bool:   ptr.New(true),
					Int:    ptr.NewInt64(42),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got testType
			err := toml.Unmarshal([]byte(tt.input), &got)
			if err != nil {
				t.Fatalf("Failed to unmarshal TOML: %v", err)
			}

			if cmp.Diff(tt.want, got) != "" {
				t.Errorf("Unmarshaled struct does not match expected value:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestUnmarshalTOMLError(t *testing.T) {
	testTomlStr := `
[struct_set]
int = "string"
string = true
bool = 42
`

	var got testType
	err := toml.Unmarshal([]byte(testTomlStr), &got)
	if err == nil {
		t.Fatalf("Expected error when unmarshaling, but got none")
	}

	if !strings.Contains(err.Error(), "cannot use") {
		t.Fatalf("Expected specific error message, got: %v", err)
	}
}

func TestMarshalTOML(t *testing.T) {
	testData := testType{
		BoolTrue:  ptr.New(true),
		BoolFalse: ptr.New(false),
		BoolUnset: ptr.ZeroBool(),
	}

	want := "true = true\nfalse = false\nunset = \n"
	data, err := toml.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	if string(data) != want {
		t.Errorf("Marshaled JSON does not match expected value:\nGot: %q\nWant: %q", data, want)
	}
}
