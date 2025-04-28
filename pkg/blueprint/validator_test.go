package blueprint

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var writeFixtures = os.Getenv("WRITE_FIXTURES") != ""

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

			schema, err := CompileSchema()
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
					outFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
					if err != nil {
						t.Fatal(err)
					}
					defer outFile.Close()

					outFile.Write([]byte(validationOutput.Error()))
					t.Logf("Written %s", output)
				}
			} else {
				want := []byte{}

				// if file output exists
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
		validFile := baseFile + ".validator.out"

		validationTest(t, file, validFile)
	}
}
