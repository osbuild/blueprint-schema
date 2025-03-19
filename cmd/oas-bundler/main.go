package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/bundler"
	"github.com/pb33f/libopenapi/datamodel"
)

func main() {
	input := flag.String("input", "./oas/blueprint-oas.yaml", "path to OAS YAML file")
	base := flag.String("base", "./oas", "base directory for resolving relative paths")
	oasYamlOutput := flag.String("oas-yaml-output", "", "bundled YAML file (none by default)")
	oasJsonOutput := flag.String("oas-json-output", "./blueprint-oas.json", "schema JSON file")
	schemaOutput := flag.String("schema-output", "./blueprint-schema.json", "output schema file")
	flag.Parse()

	specBytes, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	// load oas spec input
	doc, err := libopenapi.NewDocumentWithConfiguration([]byte(specBytes), &datamodel.DocumentConfiguration{
		BasePath:                *base,
		ExtractRefsSequentially: true,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})),
	})
	if err != nil {
		slog.Error("problem creating document", "err", err)
		panic(err)
	}

	v3Doc, errs := doc.BuildV3Model()
	if len(errs) > 0 {
		for _, e := range errs {
			slog.Error("problem building model", "err", e)
		}
	}

	// bundle the result
	bytes, err := bundler.BundleDocument(&v3Doc.Model)
	if err != nil {
		slog.Error("problem bundling document", "err", err)
		panic(err)
	}

	// write as oas yaml
	if oasYamlOutput != nil && *oasYamlOutput != "" {
		err = os.WriteFile(*oasYamlOutput, bytes, 0644)
		if err != nil {
			slog.Error("problem writing output file", "err", err)
			panic(err)
		}
	}

	// read the bundled oas yaml
	doc, err = libopenapi.NewDocumentWithConfiguration([]byte(bytes), &datamodel.DocumentConfiguration{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})),
	})
	if err != nil {
		slog.Error("problem re-opening document", "err", err)
		panic(err)
	}

	v3Doc, errs = doc.BuildV3Model()
	if len(errs) > 0 {
		for _, e := range errs {
			slog.Error("problem re-building model", "err", e)
		}
	}

	// write as oas json
	jsonBytes, err := v3Doc.Model.RenderJSON(" ")
	if err != nil {
		slog.Error("problem marshaling JSON", "err", err)
		panic(err)
	}

	err = os.WriteFile(*oasJsonOutput, jsonBytes, 0644)
	if err != nil {
		slog.Error("problem writing output file", "err", err)
		panic(err)
	}

	// read oas json into map
	var oasJson map[string]any
	err = json.Unmarshal(jsonBytes, &oasJson)
	if err != nil {
		slog.Error("problem unmarshaling oas JSON", "err", err)
		panic(err)
	}

	// write as schema json
	oasJsonBlueprint := oasJson["components"].(map[string]any)["schemas"].(map[string]any)["blueprint"].(map[string]any)
	oasJsonBlueprint["$schema"] = "http://json-schema.org/draft-04/schema#"
	schemaBytes, err := json.MarshalIndent(oasJsonBlueprint, "", " ")
	if err != nil {
		slog.Error("problem marshaling schema JSON", "err", err)
		panic(err)
	}

	err = os.WriteFile(*schemaOutput, schemaBytes, 0644)
	if err != nil {
		slog.Error("problem writing output file", "err", err)
		panic(err)
	}
}
