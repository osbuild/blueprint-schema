package ptr

import (
	"cmp"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
)

// Zeroable is a type constraint with types which supports zero references.
// This includes int, bool, and string types.
type Zeroable interface {
	~bool | cmp.Ordered
}

// Ref is a reference type that holds a value of type T.
// It provides methods to check if the reference is zero, get the value.
// It is useful for working with optional values in a type-safe manner.
type Ref[T Zeroable] struct {
	v *T
}

func New[T Zeroable](v T) Ref[T] {
	return Ref[T]{&v}
}

func NewInt64(v int64) Ref[int64] {
	return New(v)
}

// Zero returns a reference to a zero value of type T.
func Zero[T Zeroable]() Ref[T] {
	return Ref[T]{nil}
}

func ZeroBool() Ref[bool] {
	return Zero[bool]()
}

func ZeroInt() Ref[int] {
	return Zero[int]()
}

func ZeroString() Ref[string] {
	return Zero[string]()
}

// IsZero returns true for references that are zero value.
func (o Ref[T]) IsZero() bool {
	return o.v == nil
}

// NotZero returns true for references that are not zero value.
func (o Ref[T]) NotZero() bool {
	return o.v != nil
}

// Get returns the value or returns a zero value if no value was set.
func (o Ref[T]) Get() T {
	if o.v == nil {
		var defaultValue T
		return defaultValue
	}
	return *o.v
}

// GetOr returns the value if it is not zero, otherwise it returns the provided default value.
func (o Ref[T]) GetOr(defaultValue T) T {
	if o.v == nil {
		return defaultValue
	}
	return *o.v
}

func (o Ref[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.v)
}

func (o *Ref[T]) UnmarshalJSON(data []byte) error {
	if len(data) <= 0 || string(data) == "null" {
		*o = Zero[T]()
		return nil
	}

	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	*o = New(v)
	return nil
}

func (o Ref[T]) MarshalTOML() ([]byte, error) {
	if o.v == nil {
		return []byte(""), nil
	}

	return toml.Marshal(o.v)
}

var _ toml.Marshaler = (*Ref[int])(nil)
var _ toml.Unmarshaler = (*Ref[int])(nil)

func (o *Ref[T]) UnmarshalTOML(data any) error {
	b, ok := data.(T)
	if !ok {
		return fmt.Errorf("cannot use %[1]v (%[1]T) as type %[2]T", data, *new(T))
	}

	*o = New(b)
	return nil
}

// Scan implements the database/sql.Scanner interface.
func (o *Ref[T]) Scan(src any) error {
	return nil
}

var _ sql.Scanner = (*Ref[int])(nil)

// Value implements the database/sql/driver.Valuer interface.
//
// DO NOT USE THIS METHOD unless you are working with a database driver.
// Use Get() and GetOr() instead to retrieve the value.
func (o Ref[T]) Value() (driver.Value, error) {
	if o.IsZero() {
		return nil, nil
	}
	return o.Get(), nil
}

var _ driver.Valuer = (*Ref[int])(nil)
