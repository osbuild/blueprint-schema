package parse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
	"github.com/osbuild/blueprint-schema/pkg/ubp"
	bp "github.com/osbuild/blueprint/pkg/blueprint"
)

var writeFixtures = os.Getenv("WRITE_FIXTURES") != ""

func writeFile(t *testing.T, output string, buffers ...*[]byte) {
	outFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = outFile.Close()
	}()

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

// cmpTransformerForRawMessage is a cmp.Option that transforms json.RawMessage
// into a map[string]interface{} before comparison. This makes the comparison
// independent of key ordering in the JSON string.
var cmpTransformerForRawMessage = cmp.Transformer("RawMessage", func(in json.RawMessage) map[string]interface{} {
	// If the raw message is nil or empty, return a nil map.
	if len(in) == 0 {
		return nil
	}

	var out map[string]interface{}
	// Unmarshal the raw bytes into a generic map.
	// We panic on error because in a test context, malformed JSON in test data
	// is a test setup error that should be fixed.
	if err := json.Unmarshal(in, &out); err != nil {
		panic(fmt.Sprintf("cmp.Transformer: cannot unmarshal json.RawMessage: %v", err))
	}
	return out
})

// cleanDiff removes intentional non-printable characters from the diff
// output: https://github.com/google/go-cmp/issues/344
func cleanDiff(diff string) string {
	return strings.Map(func(r rune) rune {
		if r == 0x00a0 {
			return ' '
		}
		return r
	}, diff)
}

func TestFix(t *testing.T) {
	convert := func(t *testing.T, input, output string) (*ubp.Blueprint, *bp.Blueprint) {
		var resultUBP *ubp.Blueprint
		var resultBP *bp.Blueprint

		t.Run(fmt.Sprintf("Convert/%s/%s", filepath.Base(input), filepath.Base(output)), func(t *testing.T) {
			t.Logf("Converting %s", input)
			inputFile, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				_ = inputFile.Close()
			}()

			inputBuf := bytes.Buffer{}
			_, err = inputBuf.ReadFrom(inputFile)
			if err != nil {
				t.Fatal(err)
			}

			var convErrs error
			var got []byte
			details := AnyDetails{}
			ubpBP, err := UnmarshalAny(inputBuf.Bytes(), &details)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("UnmarshalAny details: %v", details)
			if details.Warnings != nil {
				convErrs = details.Warnings
			}
			resultUBP = ubpBP
			resultBP = details.Intermediate

			if details.Format == AnyFormatBPJSON || details.Format == AnyFormatBPTOML {
				// conversion was done during loading
				got, err = MarshalYAML(ubpBP)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				// no conversion was done during loading
				var result *bp.Blueprint
				exporter := conv.NewInternalExporter(ubpBP)
				result, convErrs = exporter.Export()

				got, err = toml.Marshal(result)
				if err != nil {
					t.Fatal(err)
				}
				resultBP = result
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
					_ = inFile.Close()
					if err != nil {
						t.Fatal(err)
					}
				}

				if diff := cmp.Diff(string(want), string(append(unwrapErrorsAsComments(convErrs), got...))); diff != "" {
					t.Errorf("validity mismatch (-want +got):\n%s", diff)
				}
			}
		})

		return resultUBP, resultBP
	}

	diffObjects := func(t *testing.T, s1, s2 any, diffFile string) {
		t.Run("DiffObj", func(t *testing.T) {
			if s1 == nil || s2 == nil {
				t.Fatal("objects are nil, cannot diff")
			}

			// Transform raw messages and allow unexported fields because of "union" field
			diffBuf := bytes.Buffer{}
			diff := cmp.Diff(s1, s2, cmpTransformerForRawMessage, cmp.AllowUnexported(
				ubp.Ignition{},
				ubp.DNFSource{},
				ubp.FSNodeContents{},
				ubp.NetworkService{},
				ubp.OpenSCAPTailoring{},
				ubp.StoragePartition{},
			))
			if diff != "" {
				diffBuf.WriteString(cleanDiff(diff))
			}

			if writeFixtures {
				if _, err := os.Stat(diffFile); err == nil {
					_ = os.Remove(diffFile)
				}

				if diffBuf.Len() > 0 {
					err := os.WriteFile(diffFile, diffBuf.Bytes(), 0644)
					if err != nil {
						t.Fatal(err)
					}
					t.Logf("Written diff to %s", diffFile)
				}
			} else {
				want := []byte{}

				if _, err := os.Stat(diffFile); err == nil {
					inFile, err := os.Open(diffFile)
					if err != nil {
						t.Fatal(err)
					}
					want, err = io.ReadAll(inFile)
					_ = inFile.Close()
					if err != nil {
						t.Fatal(err)
					}
				}

				// diff or diffs are not too readable, just print the diff
				if diffBuf.Len() > 0 && diffBuf.String() != string(want) {
					t.Logf("diff mismatch (-want +got):\n%s", cmp.Diff(string(want), diffBuf.String()))
				}
			}
		})
	}

	validate := func(t *testing.T, input, output string) {
		t.Run("Validate/"+input, func(t *testing.T) {
			inputFile, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				_ = inputFile.Close()
			}()

			inputBuf := bytes.Buffer{}
			_, err = inputBuf.ReadFrom(inputFile)
			if err != nil {
				t.Fatal(err)
			}

			schema, err := CompileSourceSchema()
			if err != nil {
				t.Fatal(err)
			}

			validationOutput := schema.ValidateAny(context.Background(), inputBuf.Bytes())

			if writeFixtures {
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
					_ = inFile.Close()
					if err != nil {
						t.Fatal(err)
					}
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

	extRegesp := regexp.MustCompile(`\.in\.(yaml|json|toml)$`)
	processFile := func(file string) bool {
		if s, err := os.Stat(file); err != nil || s.IsDir() {
			return false
		}

		if !extRegesp.MatchString(file) {
			t.Logf("Skipping file %q, does not match filename test pattern", file)
			return false
		}

		if strings.HasSuffix(file, ".in.yaml") {
			valid := extRegesp.ReplaceAllString(file, ".validator.out.txt")
			validate(t, file, valid)
		}

		out1 := extRegesp.ReplaceAllString(file, ".out1.txt")
		out2 := extRegesp.ReplaceAllString(file, ".out2.txt")
		inout2diff := extRegesp.ReplaceAllString(file, ".out.diff")

		ubp1, bp1 := convert(t, file, out1)
		ubp2, bp2 := convert(t, out1, out2)
		if bp1 != nil && bp2 != nil {
			diffObjects(t, bp1, bp2, inout2diff)
		} else if ubp1 != nil && ubp2 != nil {
			diffObjects(t, ubp1, ubp2, inout2diff)
		} else {
			t.Errorf("Both UBP and BP are nil for file %q", file)
			return false
		}

		return true
	}

	for _, file := range files {
		if !processFile(file) {
			t.Errorf("Failed to process file: %s", file)
		}
	}
}

func TestImportOneFile(t *testing.T) {
	inputFile, err := os.Open("../../testdata/bp-minimal-environment.in.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = inputFile.Close()
	}()

	inputBuf := bytes.Buffer{}
	_, err = inputBuf.ReadFrom(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	details := AnyDetails{}
	_, err = UnmarshalAny(inputBuf.Bytes(), &details)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Unmarshal details: %+v", details)
}
