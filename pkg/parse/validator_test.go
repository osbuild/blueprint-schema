package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/google/go-cmp/cmp"
	"github.com/osbuild/blueprint-schema/pkg/conv"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
	"github.com/osbuild/blueprint-schema/pkg/ubp"
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

func TestDetectionCounts(t *testing.T) {
	tests := []struct {
		filename string
		ubpCount int
		bpCount  int
	}{
		{"../../testdata/all-fields.in.yaml", 43, 0},
		{"../../testdata/invalid-all-empty.in.yaml", 2, 0},
		{"../../testdata/valid-empty.in.yaml", 0, 0},
		{"../../testdata/valid-empty-j.in.json", 0, 0},
		{"../../testdata/small.json", 2, 0},
		{"../../testdata/legacy-small.json", 2, 3},
		{"../../testdata/bp-oscap-generic.in.json", 12, 20},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			buf, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", tt.filename, err)
			}

			ubp, bp, err, warn := UnmarshalAny(buf)
			if err != nil {
				t.Fatalf("failed to unmarshal file %s: %v", tt.filename, err)
			}

			ubpCount := countSetFieldsRecursive(ubp)
			bpCount := countSetFieldsRecursive(bp)

			if ubpCount != tt.ubpCount {
				t.Errorf("expected UBP count %d, got %d", tt.ubpCount, ubpCount)
			}
			if bpCount != tt.bpCount {
				t.Errorf("expected BP count %d, got %d", tt.bpCount, bpCount)
			}
			if warn != nil {
				t.Logf("Unmarshal warnings: %v", warn)
			}
		})
	}
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

func TestFix(t *testing.T) {
	convert := func(t *testing.T, input, output string) *ubp.Blueprint {
		var result *ubp.Blueprint

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
			ubpBP, bpBP, err, warn := UnmarshalAny(inputBuf.Bytes())
			if err != nil {
				t.Fatal(err)
			}
			if warn != nil {
				convErrs = warn
			}
			if bpBP == nil {
				// no conversion was done during loading
				exporter := conv.NewInternalExporter(ubpBP)
				convErrs = exporter.Export()
				result := exporter.Result()
				result.Version = "1.0.0"

				got, err = toml.Marshal(result)
				if err != nil {
					t.Fatal(err)
				}
			} else {
				// conversion was done during loading
				got, err = MarshalYAML(ubpBP)
				if err != nil {
					t.Fatal(err)
				}
			}

			t.Logf("Conversion warnings: %v", convErrs)

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

			result = ubpBP
		})

		return result
	}

	diffUBP := func(t *testing.T, ubp1, ubp2 *ubp.Blueprint, diffFile string) {
		t.Run("DiffUBP", func(t *testing.T) {
			t.Logf("Diffing UBP objects")
			if ubp1 == nil || ubp2 == nil {
				t.Fatal("UBP objects are nil, cannot diff")
			}

			diffBuf := bytes.Buffer{}
			diff := cmp.Diff(ubp1, ubp2, cmpTransformerForRawMessage, cmp.AllowUnexported(
				ubp.DNFSource{},
				ubp.FSNodeContents{},
				ubp.Ignition{},
				ubp.NetworkService{},
				ubp.OpenSCAPTailoring{},
				ubp.StoragePartition{},
			))
			if diff != "" {
				diffBuf.WriteString(diff)
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
					t.Logf("Written UBP diff to %s", diffFile)
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
					t.Logf("UBP diff mismatch (-want +got):\n%s", cmp.Diff(string(want), diffBuf.String()))
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

		out1 := extRegesp.ReplaceAllString(file, ".out1.txt")
		out2 := extRegesp.ReplaceAllString(file, ".out2.txt")
		inout2diff := extRegesp.ReplaceAllString(file, ".out.diff")

		ubp1 := convert(t, file, out1)
		ubp2 := convert(t, out1, out2)
		diffUBP(t, ubp1, ubp2, inout2diff)

		return true
	}

	for _, file := range files {
		if !processFile(file) {
			t.Errorf("Failed to process file: %s", file)
		}
	}
}
