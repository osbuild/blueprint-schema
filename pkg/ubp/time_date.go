package ubp

import (
	"encoding/json"
)

// UnmarshalJSON handles default values.
func (td *TimeDate) UnmarshalJSON(data []byte) error {
	type tmpType TimeDate
	tmp := tmpType(*td)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Timezone == "" {
		tmp.Timezone = "UTC"
	}

	*td = TimeDate(tmp)
	return nil
}

// MarshalJSON handles default values.
func (td TimeDate) MarshalJSON() ([]byte, error) {
	type tmpType TimeDate
	tmp := tmpType(td)

	if tmp.Timezone == "UTC" {
		tmp.Timezone = ""
	}

	return json.Marshal(tmp)
}
