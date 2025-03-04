package blueprint

import (
	"fmt"

	"github.com/invopop/jsonschema"
)

// BoolDefaultTrue is a boolean type that defaults to true when unmarshaling from JSON.
// If the value is not set, it defaults to true. This is safer than using pointer to bool.
type BoolDefaultTrue struct {
	v   bool
	set bool
}

func (b *BoolDefaultTrue) UnmarshalJSON(data []byte) error {
	if data == nil || string(data) == "null" {
		b.set = false
		return nil
	}

	if data[0] != 't' && data[0] != 'f' {
		return fmt.Errorf("invalid bool value: %q", data)
	}

	b.v = data[0] == 't'
	b.set = true
	return nil
}

func (b *BoolDefaultTrue) MarshalJSON() ([]byte, error) {
	if !b.set {
		return []byte("null"), nil
	}

	if b.v {
		return []byte("true"), nil
	}

	return []byte("false"), nil
}

// IsSet returns true if the value is set.
func (b *BoolDefaultTrue) IsSet() bool {
	return b.set
}

// Bool returns the boolean value of the BoolDefaultTrue. If the value is not set, it returns true.
func (b *BoolDefaultTrue) Bool() bool {
	if !b.set {
		return true
	}
	return b.v
}

func (BoolDefaultTrue) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		OneOf: []*jsonschema.Schema{
			{
				Type: "boolean",
				Default: true,
			},
			{
				Type: "null",
			},
		},
	}
}
