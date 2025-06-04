package blueprint

import (
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
)

// Convert date in format YYYY-MM-DD or RFC3339 date to amount of days since epoch.
// Can be refactored to ExpireDate method once https://github.com/oapi-codegen/oapi-codegen/pull/1987 is merged.
func ExpireDateToEpochDays(date ExpireDate) (int, error) {
	if date == "" {
		return 0, nil
	}

	// Convert to RFC3339 format if not already in that format
	if !strings.Contains(string(date), "T") {
		date = ExpireDate(date + "T00:00:00Z")
	}

	t, err := time.Parse(time.RFC3339, string(date))
	if err != nil {
		return 0, err
	}
	return int(t.Unix() / (24 * 60 * 60)), nil
}

func ParseExpireDate(epochDays *int) *ExpireDate {
	if epochDays == nil || *epochDays < 0 {
		return nil
	}

	t := time.Unix(int64(*epochDays)*24*60*60, 0).UTC()
	return ptr.To(ExpireDate(t.Format(time.RFC3339)))
}
