package ubp

import (
	"encoding/json"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// UnmarshalJSON handles default values.
func (dr *DNFRepoUsage) UnmarshalJSON(data []byte) error {
	type tmpType DNFRepoUsage
	tmp := tmpType(*dr)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Configure == nil {
		tmp.Configure = ptr.To(true)
	}

	if tmp.Install == nil {
		tmp.Install = ptr.To(true)
	}

	*dr = DNFRepoUsage(tmp)
	return nil
}

// MarshalJSON handles default values.
func (dr DNFRepoUsage) MarshalJSON() ([]byte, error) {
	type tmpType DNFRepoUsage
	tmp := tmpType(dr)

	if tmp.Configure != nil && *tmp.Configure {
		tmp.Configure = nil
	}

	if tmp.Install != nil && *tmp.Install {
		tmp.Install = nil
	}

	return json.Marshal(tmp)
}

// IsZero checks if the DNFRepoUsage is empty.
func (dr DNFRepoUsage) IsZero() bool {
	return (dr.Configure == nil || *dr.Configure) && (dr.Install == nil || *dr.Install)
}
