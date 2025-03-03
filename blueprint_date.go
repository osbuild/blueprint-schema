package blueprint

import (
	"fmt"
	"time"

	"github.com/invopop/jsonschema"
)

// Date type which accepts date (YYYY-MM-DD) or date-time (RFC3339) format and only
// marshals into date format. This is needed for JSON/YAML compatibility since YAML
// automatically converts strings which look like dates into time.Time.
type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	// try to unmarshal as date only
	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		// try to unmarshal as RFC3339
		t, err = time.Parse(`"2006-01-02T15:04:05Z07:00"`, string(data))
		if err != nil {
			return fmt.Errorf("cannot parse %s neither as YYYY-MM-DD nor as RFC3339: %w", string(data), err)
		}
	}
	d.Time = t
	return nil
}
func (d *Date) String() string {
	return d.Time.Format("2006-01-02")
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format(`"2006-01-02"`)), nil
}

// DaysFrom1970 returns the number of days since 1970-01-01 as required for useradd/usermod commands.
func (d *Date) DaysFrom1970() int {
	return int(d.Time.Sub(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)
}

func (d *Date) IsZero() bool {
	return d.Time.IsZero()
}

// JSONSchemaExtend can be used to extend the generated JSON schema from Go struct tags
func (Date) JSONSchemaExtend(s *jsonschema.Schema) {
	s.Type = "string"
	s.Pattern = `^\d{4}-\d{2}-\d{2}$T?[0-9:Z-]*`
}
