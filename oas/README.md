# Blueprint OAS Source

Source files for the blueprint OAS and JSON Schema.

The target draft for the schema is OpenAPI 3.0 Schema (OAS 3.0) so make sure to only use features available in this schema draft.

## Bundling

The bundling process is done via `make schema-oas` command.

* Loads source YAML files using `/cmd/oas-bundler` tool based on `libopenapi` Go library.
* Performs bundling and writes the result into `/blueprint-oas.json` as OAS3.
* Writes `/components/schemas` separately as `/blueprint-schema.json`.
* Transforms the `/blueprint-schema.json` via `/cmd/oas2js.py` script from OAS3 to JSON Schema draft 5.

Additional linting of the OAS schema can be performed by `make schema-oas-lint`.

## Limitations

Please keep in mind the OAS3 to JSON Schema draft 5 convertor is simple, avoid complex features. Notable workarounds are described below.

### Nullable

Null type (`null`) is not supported, use `nullable` instead:

```yaml
flag:
  type: boolean
  nullable: true
```

In OpenAPI 3.1 this can be written either as:

```yaml
flag:
  oneOf:
    - type: boolean
    - type: null
```

or as more simple:

```yaml
flag:
  type: [boolean, null]
```

This is how the converter handles this situation when transforming to JSON Schema draft 5.

### Implication

Conditionals (`if-then-else`) and constants (`const`) are supported in OpenAPI 3.1, use implication instead:

```yaml
anyOf:
- not:
    properties:
      type:
        enum: ["dir"]
    required:
      - type
- not:
    required:
    - contents   
```

This can be rewritten in OpenAPI 3.1 as:

```yaml
- if:
    required:
    - type
    properties:
      type:
        const: dir
  then:
    not:
      required:
      - contents
```

For more info: https://json-schema.org/understanding-json-schema/reference/conditionals#implication

## Plans

We are currently stuck with OpenAPI 3.0 for our services because there are no good Go code generators available for OpenAPI 3.1 but once we upgrade, this bundling process can vastly simplified since OpenAPI 3.1 and JSON Schema 2020-12 are fully compatible therefore no conversion will be necessary and components can be separated into individual file
