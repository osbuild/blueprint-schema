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

	if tmp.TLSVerify == nil {
		tmp.TLSVerify = ptr.To(true)
	}

	if tmp.GPGCheck == nil {
		tmp.GPGCheck = ptr.To(true)
	}

	if tmp.GPGCheckRepo == nil {
		tmp.GPGCheckRepo = ptr.To(true)
	}

	*dr = DNFRepository(tmp)
	return nil
}

// MarshalJSON handles default values.
func (dr DNFRepository) MarshalJSON() ([]byte, error) {
	type tmpType DNFRepository
	tmp := tmpType(dr)

	if tmp.TLSVerify != nil && *tmp.TLSVerify {
		tmp.TLSVerify = nil
	}

	if tmp.GPGCheck != nil && *tmp.GPGCheck {
		tmp.GPGCheck = nil
	}

	if tmp.GPGCheckRepo != nil && *tmp.GPGCheckRepo {
		tmp.GPGCheckRepo = nil
	}

	return json.Marshal(tmp)
}
