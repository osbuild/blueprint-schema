package ubp

import (
	"encoding/json"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// UnmarshalJSON handles default values.
func (fw *FirewallFromTo) UnmarshalJSON(data []byte) error {
	type tmpType FirewallFromTo
	tmp := tmpType(*fw)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Enabled == nil {
		tmp.Enabled = ptr.To(true)
	}

	*fw = FirewallFromTo(tmp)
	return nil
}

// MarshalJSON handles default values.
func (fw FirewallFromTo) MarshalJSON() ([]byte, error) {
	type tmpType FirewallFromTo
	tmp := tmpType(fw)

	if tmp.Enabled != nil && *tmp.Enabled {
		tmp.Enabled = nil
	}

	return json.Marshal(tmp)
}
