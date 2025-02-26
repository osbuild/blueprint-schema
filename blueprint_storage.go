package blueprint

type Storage struct {
	// Device partitioning type: gpt (default) or mbr.
	Type string `json:"type" jsonschema:"required,default=gpt,enum=gpt,enum=mbr"`

	// Minimum size of the storage device. When not set, the image size is acquired from image request.
	//
	// Size must be formatted as an integer followed by whitespace and then either a decimal unit
	// (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB).
	MinSize string `json:"minsize" jsonschema:"pattern=^\\d+\\s*[BKkMGTPE]i?[BKMGTPE]?$"`

	// Partitions of the following types: plain (default), lvm, or btrfs.
	Partitions []PartitionsStorage `json:"partitions,omitempty" jsonschema:"required"`
}

type PartitionsStorage struct {
	// Partition type: plain (default), lvm, or btrfs.
	Type string `json:"type" jsonschema:"required,default=plain,enum=plain,enum=lvm,enum=btrfs"`

	// Label of the partition.
	//
	// Relevant for partition types: plain.
	Label string `json:"label,omitempty"`

	// Mount point of the partition. Required except for swap fs_type.
	//
	// Relevant for partition types: plain.
	MountPoint string `json:"mountpoint" jsonschema:"pattern=^/"`

	// File system type: ext4 (default), xfs, swap, or vfat.
	//
	// Relevant for partition types: plain.
	FSType string `json:"fs_type" jsonschema:"required,default=ext4,enum=ext4,enum=xfs,enum=swap,enum=vfat"`

	// Minimum size of the volume.
	//
	// Size must be formatted as an integer followed by whitespace and then either a decimal unit
	// (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB).
	//
	// Relevant for partition types: plain, lvm, btrfs.
	MinSize string `json:"minsize" jsonschema:"pattern=^\\d+\\s*[BKkMGTPE]i?[BKMGTPE]?$"`

	// LVM volume group name. When not set, will be generated automatically.
	//
	// Relevant for partition types: lvm.
	Name string `json:"name,omitempty"`

	// LVM logical volumes to create within the volume group.
	//
	// Relevant for partition types: lvm.
	LogicalVolumes []LogicalVolumesStorage `json:"logical_volumes,omitempty"`

	// BTRFS subvolumes to create.
	//
	// Relevant for partition types: btrfs.
	Subvolumes []SubvolumesStorage `json:"subvolumes,omitempty"`
}

type LogicalVolumesStorage struct {
	// Logical volume name. When not set, will be generated automatically.
	Name string `json:"name"`

	// Label of the logical volume.
	Label string `json:"label,omitempty"`

	// Mount point of the logical volume. Required except for swap fs_type.
	MountPoint string `json:"mountpoint" jsonschema:"pattern=^/"`

	// File system type: ext4 (default), xfs, swap, or vfat.
	FSType string `json:"fs_type" jsonschema:"required,default=ext4,enum=ext4,enum=xfs,enum=swap,enum=vfat"`

	// Minimum size of the logical volume.
	//
	// Size must be formatted as an integer followed by whitespace and then either a decimal unit
	// (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB).
	MinSize string `json:"minsize" jsonschema:"pattern=^\\d+\\s*[BKkMGTPE]i?[BKMGTPE]?$"`
}

type SubvolumesStorage struct {
	// Subvolume name, must also define its parent volume.
	Name string `json:"name" jsonschema:"required"`

	// Mount point of the subvolume. Required. Swap filesystem type is not supported on BTRFS volumes.
	MountPoint string `json:"mountpoint" jsonschema:"required,pattern=^/"`
}
