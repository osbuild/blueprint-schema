## Schema source

The target draft for the schema is: ???

### Use enum instead of const

### Use implication instead of if-then-else

https://json-schema.org/understanding-json-schema/reference/conditionals#implication

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
