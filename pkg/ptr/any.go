// Package ptr provides value to pointer and pointer to value functions.
//
// In most cases, simply use ptr.To or ptr.From functions. There are some
// extra functions for types which would normally need type conversions
// from constants, such as ptr.ToInt64.
package ptr

// To returns a pointer to the given value.
func To[T any](value T) *T {
	return &value
}

// ToNilIfEmpty returns a pointer to the given value if it is not equal to
// the zero value of type T. If it is equal to the zero value, it returns nil.
func ToNilIfEmpty[T comparable](value T) *T {
	if value == *new(T) {
		return nil
	}
	return &value
}

// ValueOrEmpty returns the value from a given pointer. If ref is nil,
// a zero value of type T will be returned.
func ValueOrEmpty[T any](ref *T) (value T) {
	if ref != nil {
		value = *ref
	}
	return
}

// ValueOr return value of pointer if not nil, else return default value.
func ValueOr[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

// OrTo returns the pointer if not nil, else return a pointer to the
// default value.
func OrTo[T any](value *T, defaultValue T) *T {
	if value == nil {
		return &defaultValue
	}
	return value
}

// EmptyToNil checks if the pointer is nil or if it points to a zero value of type T.
// If it is nil or points to a zero value, it returns nil. Otherwise, it returns
// a pointer to the value it points to.
func EmptyToNil[T comparable](value *T) *T {
	if value == nil || *value == *new(T) {
		return nil
	}
	return value
}
