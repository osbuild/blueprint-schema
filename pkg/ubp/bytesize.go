package ubp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ByteSize uint64

// HumanFriendly returns a human-readable representation of the byte size.
// Returns the largest possible unit as integer number and it uses binary units (e.g., KiB, MiB, GiB).
// Never returns a float.
func (b ByteSize) HumanFriendly() string {
	switch {
	case b%(1024*1024*1024*1024) == 0 && b >= (1024*1024*1024*1024):
		return fmt.Sprintf("%d TiB", b.IntTiB())
	case b%(1000*1000*1000*1000) == 0 && b >= (1000*1000*1000*1000):
		return fmt.Sprintf("%d TB", b.IntTB())
	case b%(1024*1024*1024) == 0 && b >= (1024*1024*1024):
		return fmt.Sprintf("%d GiB", b.IntGiB())
	case b%(1000*1000*1000) == 0 && b >= (1000*1000*1000):
		return fmt.Sprintf("%d GB", b.IntGB())
	case b%(1024*1024) == 0 && b >= (1024*1024):
		return fmt.Sprintf("%d MiB", b.IntMiB())
	case b%(1000*1000) == 0 && b >= (1000*1000):
		return fmt.Sprintf("%d MB", b.IntMB())
	case b%(1024) == 0 && b >= (1024):
		return fmt.Sprintf("%d KiB", b.IntKiB())
	case b%(1000) == 0 && b >= (1000):
		return fmt.Sprintf("%d KB", b.IntKB())
	default:
		return fmt.Sprintf("%d", b)
	}
}

func (b ByteSize) Uint64() uint64 {
	return uint64(b)
}

func (b ByteSize) Bytes() uint64 {
	return uint64(b)
}

func (b ByteSize) IntKB() uint64 {
	return uint64(b) / 1000
}

func (b ByteSize) IntMB() uint64 {
	return uint64(b) / (1000 * 1000)
}

func (b ByteSize) IntGB() uint64 {
	return uint64(b) / (1000 * 1000 * 1000)
}

func (b ByteSize) IntTB() uint64 {
	return uint64(b) / (1000 * 1000 * 1000 * 1000)
}

func (b ByteSize) IntKiB() uint64 {
	return uint64(b) / 1024
}

func (b ByteSize) IntMiB() uint64 {
	return uint64(b) / (1024 * 1024)
}

func (b ByteSize) IntGiB() uint64 {
	return uint64(b) / (1024 * 1024 * 1024)
}

func (b ByteSize) IntTiB() uint64 {
	return uint64(b) / (1024 * 1024 * 1024 * 1024)
}

func (b ByteSize) KB() float64 {
	return float64(b) / 1000
}

func (b ByteSize) MB() float64 {
	return float64(b) / (1000 * 1000)
}

func (b ByteSize) GB() float64 {
	return float64(b) / (1000 * 1000 * 1000)
}

func (b ByteSize) TB() float64 {
	return float64(b) / (1000 * 1000 * 1000 * 1000)
}

func (b ByteSize) KiB() float64 {
	return float64(b) / 1024
}

func (b ByteSize) MiB() float64 {
	return float64(b) / (1024 * 1024)
}

func (b ByteSize) GiB() float64 {
	return float64(b) / (1024 * 1024 * 1024)
}

func (b ByteSize) TiB() float64 {
	return float64(b) / (1024 * 1024 * 1024 * 1024)
}

func NewSize(bytes uint64) ByteSize {
	return ByteSize(bytes)
}

func NewSizeFloat(bytes float64) ByteSize {
	return ByteSize(uint64(bytes))
}

func ToByteSize(size uint64) ByteSize {
	return ByteSize(size)
}

func ParseSize(size string) (ByteSize, error) {
	sizeStr := strings.ToUpper(strings.TrimSpace(size))
	var numStr string
	var unitStr string

	for i := range len(sizeStr) {
		if (sizeStr[i] >= '0' && sizeStr[i] <= '9') || sizeStr[i] == '.' {
			numStr += string(sizeStr[i])
		} else {
			unitStr = strings.TrimSpace(sizeStr[i:])
			break
		}
	}

	if numStr == "" {
		return 0, fmt.Errorf("expected number: %q", size)
	}

	numberFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number in size: %v", err)
	}

	var bytes uint64

	switch unitStr {
	case "B", "", "BYTES", "BYTE":
		bytes = uint64(numberFloat)
	case "KB":
		bytes = uint64(numberFloat * 1000)
	case "MB":
		bytes = uint64(numberFloat * 1000 * 1000)
	case "GB":
		bytes = uint64(numberFloat * 1000 * 1000 * 1000)
	case "TB":
		bytes = uint64(numberFloat * 1000 * 1000 * 1000 * 1000)
	case "KIB":
		bytes = uint64(numberFloat * 1024)
	case "MIB":
		bytes = uint64(numberFloat * 1024 * 1024)
	case "GIB":
		bytes = uint64(numberFloat * 1024 * 1024 * 1024)
	case "TIB":
		bytes = uint64(numberFloat * 1024 * 1024 * 1024 * 1024)
	default:
		return 0, fmt.Errorf("unsupported unit: %s", unitStr)
	}

	return ByteSize(bytes), nil
}

func (bs *ByteSize) UnmarshalJSON(data []byte) error {
	var sizeStr string
	if err := json.Unmarshal(data, &sizeStr); err != nil {
		return fmt.Errorf("unmarshalling bytesize: %w", err)
	}

	size, err := ParseSize(sizeStr)
	if err != nil {
		return fmt.Errorf("parsing bytesize: %w", err)
	}

	*bs = size
	return nil
}

func (bs ByteSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(bs.HumanFriendly())
}
