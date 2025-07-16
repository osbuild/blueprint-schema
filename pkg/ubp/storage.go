package ubp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type StorageSize string

func (s StorageType) String() string {
	return string(s)
}

func (s FSType) String() string {
	return string(s)
}

func (s StorageType) Size() (ByteSize, error) {
	return ParseSize(string(s))
}

var ErrInvalidStorageType = errors.New("invalid storage type")

// StorageTypeDefault is used when no storage type was specified. This can only happen
// for a converted blueprint, UBP schema requires a storage type to be specified.
const StorageTypeDefault StorageType = ""

func ParseStorageType(s string) (StorageType, error) {
	switch strings.ToLower(s) {
	case "":
		return StorageTypeDefault, nil
	case "gpt":
		return StorageTypeGPT, nil
	case "dos", "mbr": // "dos" used for BP, "mbr" for UBP
		return StorageTypeMBR, nil
	default:
		return "", fmt.Errorf("%w: %q", ErrInvalidStorageType, s)
	}
}

var ErrInvalidFSType = errors.New("invalid filesystem type")

const FSTypeDefault FSType = ""

func ParseFSType(s string) (FSType, error) {
	switch strings.ToLower(s) {
	case "":
		return FSTypeDefault, nil
	case "ext4":
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
