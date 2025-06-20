package ubp

import (
	"encoding/json"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// UnmarshalJSON handles default values.
func (dr *DNFRepository) UnmarshalJSON(data []byte) error {
	type tmpType DNFRepository
	tmp := tmpType(*dr)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.SSLVerify == nil {
		tmp.SSLVerify = ptr.To(true)
	}

	*dr = DNFRepository(tmp)
	return nil
}

// MarshalJSON handles default values.
func (dr DNFRepository) MarshalJSON() ([]byte, error) {
	type tmpType DNFRepository

	if dr.SSLVerify != nil && *dr.SSLVerify {
		dr.SSLVerify = nil
	}

	tmp := tmpType(dr)
	return json.Marshal(tmp)
}
