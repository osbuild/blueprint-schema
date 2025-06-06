package blueprint

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
	ib "github.com/osbuild/blueprint/pkg/blueprint"
)

var writeFixtures = os.Getenv("WRITE_FIXTURES") != ""

func writeFile(t *testing.T, output string, buffers ...*[]byte) {
	outFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()

	for _, buf := range buffers {
		if buf == nil || *buf == nil {
			continue
		}
		_, err = outFile.Write(*buf)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("Written %s", output)
}

func unwrapErrorsAsComments(es error) []byte {
	if es == nil {
		return nil
	}

	var buf bytes.Buffer

	var errs []error
	if unwrapped, ok := es.(interface{ Unwrap() []error }); ok {
		errs = unwrapped.Unwrap()
	} else {
		errs = []error{es}
	}
	for _, err := range errs {
		if err == nil {
			continue
		}
		buf.WriteString("# ")
		buf.WriteString(err.Error())
		buf.WriteString("\n")
	}

	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	return buf.Bytes()
}

func TestFix(t *testing.T) {
	validationTest := func(t *testing.T, input, output string) {
		t.Run("Valid/"+input, func(t *testing.T) {
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

			schema, err := CompileSourceSchema()
			if err != nil {
				t.Fatal(err)
			}
			if writeFixtures {
				var validationOutput error
				if strings.HasSuffix(input, ".json") {
					validationOutput = schema.ValidateJSON(context.Background(), inputBuf.Bytes())
				} else if strings.HasSuffix(input, ".yaml") {
					validationOutput = schema.ValidateYAML(context.Background(), inputBuf.Bytes())
				} else {
					t.Fatalf("Unknown fixture extension: %s", input)
				}

				if validationOutput != nil {
					writeFile(t, output, ptr.To([]byte(validationOutput.Error())))
				}
			} else {
				want := []byte{}

				if _, err := os.Stat(output); err == nil {
					inFile, err := os.Open(output)
					if err != nil {
						t.Fatal(err)
					}
					want, err = io.ReadAll(inFile)
					inFile.Close()
					if err != nil {
						t.Fatal(err)
					}
				}

				var validationOutput error
				if strings.HasSuffix(input, ".json") {
					validationOutput = schema.ValidateJSON(context.Background(), inputBuf.Bytes())
				} else if strings.HasSuffix(input, ".yaml") {
					validationOutput = schema.ValidateYAML(context.Background(), inputBuf.Bytes())
				} else {
					t.Fatalf("Unknown fixture extension: %s", input)
				}

				var got string
				if validationOutput != nil {
					got = validationOutput.Error()
				}
				if diff := cmp.Diff(string(want), got); diff != "" {
					t.Errorf("validity mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}

	conversionTest := func(t *testing.T, input, output string) {
		t.Run("Valid/"+input, func(t *testing.T) {
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

			var convErrs error
			var got []byte
			if strings.HasSuffix(input, ".json") || strings.HasSuffix(input, ".yaml") {
				var inputBlueprint *Blueprint
				if strings.HasSuffix(input, ".json") {
					b, err := unmarshalJSON(inputBuf.Bytes())
					if err != nil {
						t.Fatal(err)
					}
					inputBlueprint = b
				}
				if strings.HasSuffix(input, ".yaml") {
					b, err := UnmarshalYAML(inputBuf.Bytes())
					if err != nil {
						t.Fatal(err)
					}
					inputBlueprint = b
				}

				exporter := NewInternalExporter(inputBlueprint)
				convErrs = exporter.Export()
				result := exporter.Result()
				result.Version = "1.0.0"
				got, err = toml.Marshal(result)
				if err != nil {
					t.Fatal(err)
				}
			} else if strings.HasSuffix(input, ".toml") {
				inputBlueprint := &ib.Blueprint{}
				err := toml.Unmarshal(inputBuf.Bytes(), inputBlueprint)
				if err != nil {
					t.Fatal(err)
				}
				importer := NewInternalImporter(inputBlueprint)
				convErrs = importer.Import()
				result := importer.Result()
				var buf bytes.Buffer
				err = WriteYAML(result, &buf)
				if err != nil {
					t.Fatal(err)
				}
				got = buf.Bytes()
			} else {
				t.Fatalf("Unknown fixture extension: %s", input)
			}

			if writeFixtures {
				writeFile(t, output, ptr.To(unwrapErrorsAsComments(convErrs)), ptr.To(got))
			} else {
				want := []byte{}

				if _, err := os.Stat(output); err == nil {
					inFile, err := os.Open(output)
					if err != nil {
						t.Fatal(err)
					}
					want, err = io.ReadAll(inFile)
					inFile.Close()
					if err != nil {
						t.Fatal(err)
					}
				}

				if diff := cmp.Diff(want, append(unwrapErrorsAsComments(convErrs), got...)); diff != "" {
					t.Errorf("validity mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}

	files, err := filepath.Glob("../../testdata/*.in.*")
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

		if strings.HasSuffix(file, ".in.yaml") {
			validFile := baseFile + ".validator.out"
			validationTest(t, file, validFile)

			suffix := baseFile + ".out.toml"
			conversionTest(t, file, suffix)
		} else if strings.HasSuffix(file, ".in.toml") {
			suffix := baseFile + ".out.yaml"
			conversionTest(t, file, suffix)
		}
	}
}
