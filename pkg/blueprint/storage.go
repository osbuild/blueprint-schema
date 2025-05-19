package blueprint

import "encoding/json"

func (s StorageType) String() string {
	return string(s)
}

func (s FSType) String() string {
	return string(s)
}

func (s StorageType) Size() (ByteSize, error) {
	return ParseSize(string(s))
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
