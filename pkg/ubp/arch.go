package ubp

import (
	"errors"
	"strings"
)

const ArchUnset Arch = ""

func (boa Arch) String() string {
	return string(boa)
}

var ErrInvalidArch = errors.New("invalid architecture")

func ParseArch(arch string) (Arch, error) {
	switch strings.ToLower(arch) {
	case "x86_64":
		return ArchX8664, nil
	case "aarch64":
		return ArchAarch64, nil
	case "ppc64le":
		return ArchPPC64le, nil
	case "s390x":
		return ArchS390x, nil
	case "riscv64":
		return ArchRISCV64, nil
	default:
		return "", ErrInvalidArch
	}
}
