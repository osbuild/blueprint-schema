package blueprint

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// EpochDays represents a date as the number of days since the Unix epoch (1970-01-01).
type EpochDays int

// Days returns the number of days since the Unix epoch.
func (e EpochDays) Days() int {
	return int(e)
}

// NewStringEpochDays converts date in format YYYY-MM-DD or RFC3339 date to amount of days since epoch.
func NewStringEpochDays(date string) (*EpochDays, error) {
	if date == "" {
		return nil, nil
	}

	// Convert to RFC3339 format if not already in that format
	if !strings.Contains(date, "T") {
		date = date + "T00:00:00Z"
	}

	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}
	return ptr.To(EpochDays(t.UTC().Unix() / (24 * 60 * 60))), nil
}

// NewIntEpochDays creates a new EpochDays from an integer representing the number of days since the epoch.
func NewIntEpochDays(epochDays int) *EpochDays {
	return ptr.To(EpochDays(epochDays))
}

func (e *EpochDays) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == `""` {
		return nil
	}

	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	ed, err := NewStringEpochDays(dateStr)
	if err != nil {
		return err
	}

	*e = *ed
	return nil
}

func (e EpochDays) MarshalJSON() ([]byte, error) {
	t := time.Unix(int64(e.Days())*(24*60*60), 0).UTC()
	return json.Marshal(t.Format("2006-01-02"))
}
