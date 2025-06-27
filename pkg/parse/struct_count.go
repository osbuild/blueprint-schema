package parse

import (
	"reflect"
)

// isFieldSet determines if a reflect.Value is "set".
// "Set" means it is not its type's natural zero value.
func isFieldSet(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return !v.IsNil()
	default:
		return !v.IsZero()
	}
}

// countRecursive is the helper function that traverses the struct.
// It takes a reflect.Value of a struct.
func countRecursive(val reflect.Value) int {
	// If the value is not a struct, there are no fields to count.
	if val.Kind() != reflect.Struct {
		return 0
	}

	count := 0
	structType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		structField := structType.Field(i)

		// Skip unexported fields. PkgPath is non-empty for unexported fields.
		if structField.PkgPath != "" {
			continue
		}

		// Check if the field is set (not its zero value).
		if isFieldSet(fieldVal) {
			// This is the recursive part. If the field is a struct or a pointer to one,
			// we recurse into it. Otherwise, we just count it as 1.
			kind := fieldVal.Kind()

			if kind == reflect.Struct {
				// If the field is a struct, add the count of its own set fields.
				count += countRecursive(fieldVal)
			} else if kind == reflect.Ptr && fieldVal.Elem().Kind() == reflect.Struct {
				// If the field is a non-nil pointer to a struct, recurse into the element it points to.
				count += countRecursive(fieldVal.Elem())
			} else {
				// For any other "set" type (int, string, slice, map, etc.), count it as 1.
				count++
			}
		}
	}
	return count
}

// countSetFieldsRecursive takes a struct (or a pointer to a struct) s and
// returns a total count of its exported fields that are not their zero value.
// It recursively traverses into nested structs and pointers to structs.
func countSetFieldsRecursive(s any) int {
	if s == nil {
		return 0
	}

	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return 0
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("input must be a struct or a pointer to a struct, got " + val.Kind().String())
	}

	return countRecursive(val)
}
