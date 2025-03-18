package blueprint

type DiskCustomization struct {
	// Type of the partition table: gpt or dos.
	// Optional, the default depends on the distro and image type.
	Type       string
	MinSize    uint64
	Partitions []PartitionCustomization
}

// PartitionCustomization defines a single partition on a disk. The Type
// defines the kind of "payload" for the partition: plain, lvm, or btrfs.
//   - plain: the payload will be a filesystem on a partition (e.g. xfs, ext4).
//     See [FilesystemTypedCustomization] for extra fields.
//   - lvm: the payload will be an LVM volume group. See [VGCustomization] for
//     extra fields
//   - btrfs: the payload will be a btrfs volume. See
//     [BtrfsVolumeCustomization] for extra fields.
type PartitionCustomization struct {
	// The type of payload for the partition (optional, defaults to "plain").
	Type string `json:"type" toml:"type"`

	// Minimum size of the partition that contains the filesystem (for "plain"
	// filesystem), volume group ("lvm"), or btrfs volume ("btrfs"). The final
	// size of the partition will be larger than the minsize if the sum of the
	// contained volumes (logical volumes or subvolumes) is larger. In
	// addition, certain mountpoints have required minimum sizes. See
	// https://osbuild.org/docs/user-guide/partitioning for more details.
	// (optional, defaults depend on payload and mountpoints).
	MinSize string `json:"minsize" toml:"minsize"`

	BtrfsVolumeCustomization

	VGCustomization

	FilesystemTypedCustomization
}

// A filesystem on a plain partition or LVM logical volume.
// Note the differences from [FilesystemCustomization]:
//   - Adds a label.
//   - Adds a filesystem type (fs_type).
//   - Does not define a size. The size is defined by its container: a
//     partition ([PartitionCustomization]) or LVM logical volume
//     ([LVCustomization]).
//
// Setting the FSType to "swap" creates a swap area (and the Mountpoint must be
// empty).
type FilesystemTypedCustomization struct {
	Mountpoint string `json:"mountpoint" toml:"mountpoint"`
	Label      string `json:"label,omitempty" toml:"label,omitempty"`
	FSType     string `json:"fs_type,omitempty" toml:"fs_type,omitempty"`
}

// An LVM volume group with one or more logical volumes.
type VGCustomization struct {
	// Volume group name (optional, default will be automatically generated).
	Name           string            `json:"name" toml:"name"`
	LogicalVolumes []LVCustomization `json:"logical_volumes,omitempty" toml:"logical_volumes,omitempty"`
}

type LVCustomization struct {
	// Logical volume name
	Name string `json:"name,omitempty" toml:"name,omitempty"`

	// Minimum size of the logical volume
	MinSize string `json:"minsize,omitempty" toml:"minsize,omitempty"`

	FilesystemTypedCustomization
}

// A btrfs volume consisting of one or more subvolumes.
type BtrfsVolumeCustomization struct {
	Subvolumes []BtrfsSubvolumeCustomization
}

type BtrfsSubvolumeCustomization struct {
	// The name of the subvolume, which defines the location (path) on the
	// root volume (required).
	// See https://btrfs.readthedocs.io/en/latest/Subvolumes.html
	Name string `json:"name" toml:"name"`

	// Mountpoint for the subvolume.
	Mountpoint string `json:"mountpoint" toml:"mountpoint"`
}
