package ubp

import (
	"encoding/json"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// UnmarshalJSON handles default values.
func (c *Container) UnmarshalJSON(data []byte) error {
	type tmpType Container
	tmp := tmpType(*c)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.TLSVerify == nil {
		tmp.TLSVerify = ptr.To(true)
	}

	*c = Container(tmp)
	return nil
}

// MarshalJSON handles default values.
func (c Container) MarshalJSON() ([]byte, error) {
	type tmpType Container
	tmp := tmpType(c)

	if tmp.TLSVerify != nil && *tmp.TLSVerify {
		tmp.TLSVerify = nil
	}

	return json.Marshal(tmp)
}
