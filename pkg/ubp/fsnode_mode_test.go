package ubp

import (
	"testing"
)

func TestFSNodeMode(t *testing.T) {
	tests := []struct {
		mode      FSNodeMode
		marshal   string
		unmarshal string
	}{
		{mode: 0777, marshal: `"0777"`, unmarshal: `"0o777"`},
		{mode: 0644, marshal: `"0644"`, unmarshal: `"0644"`},
		{mode: 0755, marshal: `"0755"`, unmarshal: `"755"`},
	}

	for _, tt := range tests {
		t.Run(tt.marshal, func(t *testing.T) {
			// Test MarshalJSON
			data, err := tt.mode.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON failed: %v", err)
			}
			if string(data) != tt.marshal {
				t.Errorf("expected %s, got %s", tt.marshal, data)
			}

			// Test UnmarshalJSON
			var mode FSNodeMode
			err = mode.UnmarshalJSON([]byte(tt.unmarshal))
			if err != nil {
				t.Fatalf("UnmarshalJSON failed: %v", err)
			}
			if mode != tt.mode {
				t.Errorf("expected %d, got %d", tt.mode, mode)
			}
		})
	}
}
