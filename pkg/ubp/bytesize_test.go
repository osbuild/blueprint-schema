package ubp

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestByteSize(t *testing.T) {
	tests := []struct {
		input    string
		expected ByteSize
		err      string
	}{
		{"", 0, `expected number: ""`},
		{"xxx", 0, `expected number: "xxx"`},
		{"0", 0, ""},
		{"0B", 0, ""},
		{"1B", 1, ""},
		{"1 B", 1, ""},
		{"1 Byte", 1, ""},
		{"2 Bytes", 2, ""},
		{"1.5B", 1, ""},
		{"1 KB", 1000, ""},
		{"1.5 KB", 1500, ""},
		{"1.5 KiB", 1536, ""},
		{"1.5 MB", 1500000, ""},
		{"1.5 MiB", 1572864, ""},
		{"1.5 GB", 1500000000, ""},
		{"1.5 GiB", 1610612736, ""},
		{"1.5 TB", 1500000000000, ""},
		{"1.5 TiB", 1649267441664, ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseSize(test.input)
			if err == nil && test.err != "" {
				t.Errorf("expected error: %s, got nil", test.err)
			} else if err != nil && test.err == "" {
				t.Errorf("expected nil, got error: %v", err)
			} else if err != nil && test.err != "" && err.Error() != test.err {
				t.Errorf("expected error: %s, got: %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestByteSizeHuman(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{1000, "1 KB"},
		{1024, "1 KiB"},
		{1025, "1025"},
		{4096, "4 KiB"},
		{10000, "10 KB"},
		{100000000, "100 MB"},
		{100000000000, "100 GB"},
		{53687091200, "50 GiB"},
		{1099511627776, "1 TiB"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			bs := ByteSize(test.input)
			result := bs.HumanFriendly()

			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		input    ByteSize
		expected string
	}{
		{0, `"0"`},
		{1, `"1"`},
		{1000, `"1 KB"`},
		{1500, `"1500"`},
		{10240, `"10 KiB"`},
		{1048576, `"1 MiB"`},
		{1073741824, `"1 GiB"`},
		{1099511627776, `"1 TiB"`},
		{100000000000, `"100 GB"`},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			data, err := json.Marshal(test.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if string(data) != test.expected {
				t.Errorf("expected %s, got %s", test.expected, data)
			}
		})
	}
}

func TestUnarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected ByteSize
	}{
		{`"0B"`, 0},
		{`"1B"`, 1},
		{`"1500B"`, 1500},
		{`"10kB"`, 10000},
		{`"10KiB"`, 10240},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var result ByteSize
			if err := json.Unmarshal([]byte(test.input), &result); err != nil {
				t.Errorf("unexpected error on unmarshal: %v", err)
			}
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.input, result)
			}
		})
	}
}
