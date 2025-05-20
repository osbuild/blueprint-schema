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

// ToErr returns a pointer to the given value and error which is unchanged.
func ToErr[T any](value T, e error) (*T, error) {
	return &value, e
}

// From returns the value from a given pointer. If ref is nil, a zero
// value of type T will be returned.
func From[T any](ref *T) (value T) {
	if ref != nil {
		value = *ref
	}
	return
}

// FromOr return value of pointer if not nil, else return default value.
func FromOr[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

// Or returns the pointer if not nil, else return default value.
func Or[T any](value *T, defaultValue T) *T {
	if value == nil {
		return &defaultValue
	}
	return value
}

// FromOrEmpty returns the value or empty value in case the value is nil.
func FromOrEmpty[T any](ref *T) (value T) {
	if ref != nil {
		value = *ref
	} else {
		value = *new(T)
	}
	return
}
