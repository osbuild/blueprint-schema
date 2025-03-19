## Schema source

The target draft for the schema is OpenAPI 3.0 Schema (OAS 3.0)

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

For more info:
https://json-schema.org/understanding-json-schema/reference/conditionals#implication
