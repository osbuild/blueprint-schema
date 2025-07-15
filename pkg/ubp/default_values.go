package ubp

import (
	"encoding/json"
	"fmt"
)

var ErrPopulateDefaults = "error during populating defaults"

// PopulateDefaults applies default values to the Blueprint fields that are not set.
// It marshals the Blueprint to JSON and unmarshals it back to apply defaults.
// The function is called after conversion from BP to ensure consistency, and also
// from unit tests to ensure consistency with schema documentation.
func PopulateDefaults(ubp *Blueprint) error {
	if ubp == nil {
		return nil
	}

	buf, err := json.Marshal(ubp)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrPopulateDefaults, err)
	}

	ubpDefs := &Blueprint{}
	if err := json.Unmarshal(buf, ubpDefs); err != nil {
		return fmt.Errorf("%s: %w", ErrPopulateDefaults, err)
	}

	*ubp = *ubpDefs
	return nil
}
