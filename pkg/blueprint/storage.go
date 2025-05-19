package blueprint

func (s StorageType) String() string {
	return string(s)
}

func (s StorageType) Size() (ByteSize, error) {
	return ParseSize(string(s))
}
