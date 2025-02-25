package validate_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/goccy/go-json"
	"github.com/google/go-cmp/cmp"
	"github.com/invopop/yaml"
	blueprint "github.com/osbuild/blueprint-schema"
	validate "github.com/osbuild/blueprint-schema/validate"
	"github.com/wI2L/jsondiff"
)

var writeFixtures = os.Getenv("WRITE_FIXTURES") != ""

func TestFix(t *testing.T) {
	marshalTest := func(t *testing.T, input, output string) {
		t.Run("Marshal/"+input, func(t *testing.T) {
			t.Parallel()

			var b *blueprint.Blueprint
			inputFile, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer inputFile.Close()

			if strings.HasSuffix(input, ".json") {
				b, err = blueprint.ReadJSON(inputFile)
				if err != nil {
					t.Fatal(err)
				}
			} else if strings.HasSuffix(input, ".yaml") {
				b, err = blueprint.ReadYAML(inputFile)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				t.Fatalf("Unknown fixture extension: %s", input)
			}

			if writeFixtures {
				outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					t.Fatal(err)
				}
				defer outputFile.Close()

				err = blueprint.WriteYAML(b, outputFile)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("Written %s", output)
			} else {
				outputFile, err := os.Open(output)
				if err != nil {
					t.Fatal(err)
				}
				defer outputFile.Close()
				want, err := io.ReadAll(outputFile)
				if err != nil {
					t.Fatal(err)
				}

				buf := bytes.Buffer{}
				err = blueprint.WriteYAML(b, &buf)
				if err != nil {
					t.Fatal(err)
				}

				if diff := cmp.Diff(want, buf.Bytes()); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}

	validJSONstring := "{ \"valid\": true }\n"
	validationTest := func(t *testing.T, input, output string) {
		t.Run("Valid/"+input, func(t *testing.T) {
			t.Parallel()

			inputFile, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer inputFile.Close()

			inputBuf := bytes.Buffer{}
			_, err = inputBuf.ReadFrom(inputFile)
			if err != nil {
				t.Fatal(err)
			}

			var data map[string]any

			if strings.HasSuffix(input, ".yaml") {
				jsonBuffer, err := yaml.YAMLToJSON(inputBuf.Bytes())
				if err != nil {
					t.Fatal(err)
				}

				err = json.Unmarshal(jsonBuffer, &data)
				if err != nil {
					t.Fatal(err)
				}

			} else if strings.HasSuffix(input, ".json") {
				err = json.Unmarshal(inputBuf.Bytes(), &data)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				t.Fatalf("Unknown fixture extension: %s", input)
			}

			schema, err := validate.CompileSchema()
			if err != nil {
				t.Fatal(err)
			}
			if writeFixtures {
				outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					t.Fatal(err)
				}
				defer outputFile.Close()

				valid, details := schema.ValidateMap(data)
				if valid {
					outputFile.WriteString(validJSONstring)
				} else {
					outputFile.WriteString(details)
				}
				t.Logf("Written %s", output)
			} else {
				outputFile, err := os.Open(output)
				if err != nil {
					t.Fatal(err)
				}
				defer outputFile.Close()
				want, err := io.ReadAll(outputFile)
				if err != nil {
					t.Fatal(err)
				}

				valid, details := schema.ValidateMap(data)
				if valid {
					details = validJSONstring
				}

				// The root 'errors` field must be ignored because it is marshaled in random order as a string:
				// "Property 'a', 'b' does not match the schema" where a and b can be in any order. This has no effect
				// on testing tho since errors are reported via the 'details'. Reported upstream:
				// https://github.com/kaptinlin/jsonschema/issues/28
				diff, err := jsondiff.CompareJSON([]byte(want), []byte(details), jsondiff.Ignores("/errors"))
				if err != nil {
					t.Fatal(err)
				}

				if len(diff) > 0 {
					b, err := json.MarshalIndent(diff, "", "    ")
					if err != nil {
						t.Fatal(err)
					}
					t.Errorf("JSON diff for %q:\n%s", output, string(b))
				}
			}
		})
	}

	files, err := filepath.Glob("fixtures/*.in.*")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if s, err := os.Stat(file); err != nil || s.IsDir() {
			continue
		}

		format := filepath.Ext(file)
		fileWithoutFormat := file[0 : len(file)-len(format)]
		direction := filepath.Ext(fileWithoutFormat)
		baseFile := file[0 : len(fileWithoutFormat)-len(direction)]
		outFile := baseFile + ".out.yaml"
		validFile := baseFile + ".valid.json"

		marshalTest(t, file, outFile)
		validationTest(t, file, validFile)
	}
}
