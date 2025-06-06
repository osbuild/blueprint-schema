package blueprint

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StorageSize string

func (s StorageType) String() string {
	if s == "" {
		return StorageTypeGPT.String()
	}

	return string(s)
}

func (s FSType) String() string {
	if s == "" {
		return FSTypeExt4.String()
	}

	return string(s)
}

func (s StorageType) Size() (ByteSize, error) {
	return ParseSize(string(s))
}

var ErrInvalidStorageType = errors.New("invalid storage type")

func ParseStorageType(s string) (StorageType, error) {
	switch strings.ToLower(s) {
	case "gpt", "":
		return StorageTypeGPT, nil
	case "mbr":
		return StorageTypeMBR, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidStorageType, s)
	}
}

var ErrInvalidFSType = errors.New("invalid filesystem type")

func ParseFSType(s string) (FSType, error) {
	switch strings.ToLower(s) {
	case "ext4", "":
		return FSTypeExt4, nil
	case "vfat":
		return FSTypeVFAT, nil
	case "xfs":
		return FSTypeXFS, nil
	case "swap":
		return FSTypeSwap, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidFSType, s)
	}
}

func (sp StoragePartition) SelectUnion() (PartitionPlain, PartitionLVM, PartitionBTRFS, error) {
	var pp PartitionPlain
	err := json.Unmarshal(sp.union, &pp)
	if err != nil {
		return PartitionPlain{}, PartitionLVM{}, PartitionBTRFS{}, err
	}

	var pl PartitionLVM
	err = json.Unmarshal(sp.union, &pl)
	if err != nil {
		return PartitionPlain{}, PartitionLVM{}, PartitionBTRFS{}, err
	}

	var pb PartitionBTRFS
	err = json.Unmarshal(sp.union, &pb)
	if err != nil {
		return PartitionPlain{}, PartitionLVM{}, PartitionBTRFS{}, err
	}

	return pp, pl, pb, nil
}

func StoragePartitionFromPlain(node PartitionPlain) StoragePartition {
	u, _ := json.Marshal(node)
	return StoragePartition{union: u}
}

func StoragePartitionFromLVM(node PartitionLVM) StoragePartition {
	u, _ := json.Marshal(node)
	return StoragePartition{union: u}
}

func StoragePartitionFromBTRFS(node PartitionBTRFS) StoragePartition {
	u, _ := json.Marshal(node)
	return StoragePartition{union: u}
}
