package ubp

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEpochDays(t *testing.T) {
	tests := []struct {
		json     string
		expected int
	}{
		{`"1970-01-01"`, 0},
		{`"1970-01-02"`, 1},
		{`"1980-01-15"`, 3666},
		{`"2023-10-01"`, 19631},
		{`"2023-10-01T00:00:00Z"`, 19631},
		{`"2023-10-01T12:34:56Z"`, 19631},
	}

	for _, test := range tests {
		bsi := []byte(test.json)
		var ed EpochDays
		err := json.Unmarshal(bsi, &ed)
		if err != nil {
			t.Errorf("Expected no error for input %s, got %v", test.json, err)
		}

		expectedEd := EpochDays(test.expected)
		if ed != expectedEd {
			t.Errorf("Expected %d, got %d for input %s", expectedEd.Days(), ed.Days(), test.json)
		}

		bs, err := json.Marshal(expectedEd)
		if err != nil {
			t.Errorf("Expected no error for marshaling %d, got %v", expectedEd.Days(), err)
		}

		if string(bs) != test.json && !strings.Contains(test.json, "T") {
			t.Errorf("Expected %s, got %s for marshaling %d", test.json, string(bs), expectedEd.Days())
		}
	}
}
