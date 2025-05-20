package blueprint

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	int "github.com/osbuild/blueprint/pkg/blueprint"
)

// InternalExporter is used to convert a blueprint to the internal representation.
//
// TODO we need a place for distro, registration and repos, likely BuildOptions:
// https://github.com/osbuild/blueprint-schema/issues/23
type InternalExporter struct {
	from *Blueprint
	to   *int.Blueprint
	log  *logs
}

func NewInternalExporter(inputBlueprint *Blueprint) *InternalExporter {
	return &InternalExporter{
		from: inputBlueprint,
		log:  newCollector(),
	}
}

// ExportInternal converts the blueprint to the internal representation.
func (e *InternalExporter) Export(ed ComposeRequest) error {
	to := &int.Blueprint{}

	to.Name = e.from.Name
	to.Description = e.from.Description
	if ed.Version == "" {
		// Create monotonic incremental version number based on miliseconds
		to.Version = int64ToVersion(uint64(time.Now().UTC().UnixMilli()))
	} else {
		to.Version = ed.Version
	}

	to.Packages = e.exportPackages()
	to.EnabledModules = e.exportModules()
	to.Groups = e.exportGroups()
	to.Containers = e.exportContainers()
	to.Customizations = e.exportCustomizations()
	to.Distro = ed.Distro
	to.Arch = ed.Arch

	e.to = to
	return e.log.Errors()
}

func (e *InternalExporter) Result() *int.Blueprint {
	return e.to
}

func (e *InternalExporter) exportPackages() []int.Package {
	var s []int.Package

	for _, pkg := range e.from.DNF.Packages {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, int.Package{
			Name:    p[0],
			Version: p[1],
		})
	}

	return s
}

func (e *InternalExporter) exportGroups() []int.Group {
	var s []int.Group

	for _, pkg := range e.from.DNF.Groups {
		s = append(s, int.Group{
			Name: pkg,
		})
	}

	return s
}

func (e *InternalExporter) exportModules() []int.EnabledModule {
	var s []int.EnabledModule

	for _, pkg := range e.from.DNF.Modules {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, int.EnabledModule{
			Name:   p[0],
			Stream: p[1],
		})
	}

	return s
}

func (e *InternalExporter) exportContainers() []int.Container {
	var s []int.Container

	for _, container := range e.from.Containers {
		s = append(s, int.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    &container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return s
}

func (e *InternalExporter) exportCustomizations() *int.Customizations {
	to := &int.Customizations{}

	to.Hostname = &e.from.Hostname
	to.Kernel = e.exportKernel()
	to.User = e.exportUserCustomization()
	to.Group = e.exportGroupCustomization()
	to.Timezone = e.exportTimezoneCustomization()
	to.Locale = e.exportLocaleCustomization()
	to.Firewall = e.exportFirewallCustomization()
	to.Services = e.exportSystemdCustomization()
	to.Disk = e.exportStorage()
	to.InstallationDevice, to.Installer = e.exportInstaller()
	to.RHSM, to.FDO = e.exportRegistration()

	return to
}

func (e *InternalExporter) exportKernel() *int.KernelCustomization {
	if e.from.Kernel == nil {
		return nil
	}

	to := &int.KernelCustomization{}
	to.Name = e.from.Kernel.Package
	to.Append = strings.Join(e.from.Kernel.CmdlineAppend, " ")

	return to
}

func (e *InternalExporter) exportUserCustomization() []int.UserCustomization {
	if e.from.Accounts.Users == nil {
		return nil
	}

	var s []int.UserCustomization
	for _, u := range e.from.Accounts.Users {
		uc := int.UserCustomization{}
		uc.Name = u.Name
		uc.Description = &u.Description
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = &u.SSHKeys[0]
		} else if len(u.SSHKeys) > 1 {
			uc.Key = &u.SSHKeys[0]
			e.log.Printf("only one ssh key supported for user: %s", u.Name)
		}
		uc.Home = &u.Home
		uc.Shell = &u.Shell
		uc.Groups = u.Groups
		if u.UID != 0 {
			uc.UID = &u.UID
		}
		if u.GID != 0 {
			uc.GID = &u.GID
		}
		if u.Expires != nil {
			var err error
			uc.ExpireDate, err = ptr.ToErr(ExpireDateToEpochDays(*u.Expires))
			if err != nil {
				e.log.Printf("error converting expire date for user %s: %v", u.Name, err)
			}
		}
		if u.ForcePasswordChange != nil {
			uc.ForcePasswordReset = u.ForcePasswordChange
		}

		s = append(s, uc)
	}

	return s
}

func (e *InternalExporter) exportGroupCustomization() []int.GroupCustomization {
	if e.from.Accounts.Groups == nil {
		return nil
	}

	var s []int.GroupCustomization
	for _, g := range e.from.Accounts.Groups {
		gc := int.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = &g.GID
		}
		s = append(s, gc)
	}

	return s
}

func (e *InternalExporter) exportTimezoneCustomization() *int.TimezoneCustomization {
	if e.from.Timedate == nil {
		return nil
	}

	to := &int.TimezoneCustomization{}
	to.Timezone = &e.from.Timedate.Timezone
	to.NTPServers = e.from.Timedate.NTPServers

	return to
}

func (e *InternalExporter) exportLocaleCustomization() *int.LocaleCustomization {
	if e.from.Locale == nil {
		return nil
	}

	to := &int.LocaleCustomization{}
	if len(e.from.Locale.Keyboards) > 0 {
		to.Keyboard = &e.from.Locale.Keyboards[0]
		if len(e.from.Locale.Keyboards) > 1 {
			e.log.Println("only one keyboard layout supported, selecting first one")
		}
	}
	to.Languages = e.from.Locale.Languages

	return to
}

func (e *InternalExporter) exportFirewallCustomization() *int.FirewallCustomization {
	if e.from.Network.Firewall == nil || len(e.from.Network.Firewall.Services) == 0 {
		return nil
	}

	to := &int.FirewallCustomization{
		Ports: make([]string, 0),
		Services: &int.FirewallServicesCustomization{
			Enabled:  make([]string, 0),
			Disabled: make([]string, 0),
		},
	}
	for i, s := range e.from.Network.Firewall.Services {
		fs, fp, fft, err := s.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse network service %d: %v", i, err)
			continue
		}

		proto := "tcp"
		if fs.Service != "" {
			if fs.Enabled == nil || *fs.Enabled {
				to.Services.Enabled = append(to.Services.Enabled, fs.Service)
			} else {
				to.Services.Disabled = append(to.Services.Disabled, fs.Service)
			}
		} else if fp.Port != 0 {
			if fp.Protocol != "" {
				proto = fp.Protocol.String()
			}
			srv := fmt.Sprintf("%d/%s", fp.Port, proto)

			if fp.Enabled == nil || *fp.Enabled {
				to.Ports = append(to.Ports, srv)
			} else {
				e.log.Printf("network service %d error: port number %d cannot be disabled", i, fp.Port)
				continue
			}
		} else if fft.From != 0 && fft.To != 0 {
			if fft.Protocol != "" {
				proto = fft.Protocol.String()
			}
			srv := fmt.Sprintf("%d-%d/%s", fft.From, fft.To, proto)

			if fft.Enabled == nil || *fft.Enabled {
				to.Ports = append(to.Ports, srv)
			} else {
				e.log.Printf("network service %d error: port number %d cannot be disabled", i, fp.Port)
				continue
			}
		} else {
			e.log.Printf("network service %d error: one of service, port or from and to present", i)
		}
	}

	if len(to.Ports) == 0 {
		to.Ports = nil
	}
	if len(to.Services.Enabled) == 0 {
		to.Services.Enabled = nil
	}
	if len(to.Services.Disabled) == 0 {
		to.Services.Disabled = nil
	}
	if to.Services.Enabled == nil && to.Services.Disabled == nil {
		to.Services = nil
	}

	return to
}

func (e *InternalExporter) exportSystemdCustomization() *int.ServicesCustomization {
	if e.from.Systemd == nil {
		return nil
	}

	to := &int.ServicesCustomization{}

	if len(e.from.Systemd.Enabled) > 0 {
		to.Enabled = make([]string, len(e.from.Systemd.Enabled))
		copy(to.Enabled, e.from.Systemd.Enabled)
	}
	if len(e.from.Systemd.Disabled) > 0 {
		to.Disabled = make([]string, len(e.from.Systemd.Disabled))
		copy(to.Disabled, e.from.Systemd.Disabled)
	}
	if len(e.from.Systemd.Masked) > 0 {
		to.Masked = make([]string, len(e.from.Systemd.Masked))
		copy(to.Masked, e.from.Systemd.Masked)
	}

	return to
}

func (e *InternalExporter) exportStorage() *int.DiskCustomization {
	if e.from.Storage == nil {
		return nil
	}

	to := &int.DiskCustomization{}
	to.Type = e.from.Storage.Type.String()
	size, err := ParseSize(e.from.Storage.Minsize)
	if err != nil {
		e.log.Printf("error parsing size %s: %v", e.from.Storage.Minsize, err)
	} else {
		to.MinSize = size.Bytes()
	}

	if len(e.from.Storage.Partitions) == 0 {
		return to
	}

	to.Partitions = make([]int.PartitionCustomization, 0, len(e.from.Storage.Partitions))
	for i, p := range e.from.Storage.Partitions {
		pp, pl, pb, err := p.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse partition %d: %v", i, err)
			continue
		}

		if pp.Type == PartTypePlain {
			size, err := ParseSize(pp.Minsize)
			if err != nil {
				e.log.Printf("error parsing size %q: %v", pp.Minsize, err)
				continue
			}

			part := &int.PartitionCustomization{
				Type:    "plain",
				MinSize: size.Bytes(),
				FilesystemTypedCustomization: int.FilesystemTypedCustomization{
					Label:      pp.Label,
					Mountpoint: pp.Mountpoint,
					FSType:     pp.FSType.String(),
				},
			}

			to.Partitions = append(to.Partitions, *part)
		} else if pl.Type == PartTypeLVM {
			size, err := ParseSize(pl.Minsize)
			if err != nil {
				e.log.Printf("error parsing size %q: %v", pl.Minsize, err)
				continue
			}

			part := &int.PartitionCustomization{
				Type:    "lvm",
				MinSize: size.Bytes(),
				VGCustomization: int.VGCustomization{
					Name:           pl.Name,
					LogicalVolumes: make([]int.LVCustomization, 0, len(pl.LogicalVolumes)),
				},
			}

			for _, lv := range pl.LogicalVolumes {
				lvSize, err := ParseSize(lv.Minsize)
				if err != nil {
					e.log.Printf("error parsing size %q: %v", lv.Minsize, err)
					continue
				}
				lvc := int.LVCustomization{
					Name:    lv.Name,
					MinSize: lvSize.Bytes(),
					FilesystemTypedCustomization: int.FilesystemTypedCustomization{
						Label:      lv.Label,
						Mountpoint: lv.Mountpoint,
						FSType:     lv.FSType.String(),
					},
				}
				part.VGCustomization.LogicalVolumes = append(part.VGCustomization.LogicalVolumes, lvc)
			}

			to.Partitions = append(to.Partitions, *part)
		} else if pb.Type == PartTypeBTRFS {
			size, err := ParseSize(pb.Minsize)
			if err != nil {
				e.log.Printf("error parsing size %q: %v", pb.Minsize, err)
				continue
			}

			part := &int.PartitionCustomization{
				Type:    "btrfs",
				MinSize: size.Bytes(),
				BtrfsVolumeCustomization: int.BtrfsVolumeCustomization{
					Subvolumes: make([]int.BtrfsSubvolumeCustomization, 0, len(pb.Subvolumes)),
				},
			}

			for _, sv := range pb.Subvolumes {
				svc := int.BtrfsSubvolumeCustomization{
					Name:       sv.Name,
					Mountpoint: sv.Mountpoint,
				}
				part.BtrfsVolumeCustomization.Subvolumes = append(part.BtrfsVolumeCustomization.Subvolumes, svc)
			}

			to.Partitions = append(to.Partitions, *part)
		} else {
			e.log.Printf("unknown partition type %q", pl.Type)
		}
	}

	return to
}

func (e *InternalExporter) exportInstaller() (string, *int.InstallerCustomization) {
	var installationDevice string
	if e.from.Installer == nil {
		return installationDevice, nil
	}

	var to *int.InstallerCustomization
	if e.from.Installer.Anaconda != nil {
		to = &int.InstallerCustomization{
			Modules: &int.AnacondaModules{},
		}
		to.Unattended = e.from.Installer.Anaconda.Unattended
		to.SudoNopasswd = e.from.Installer.Anaconda.SudoNOPASSWD
		if e.from.Installer.Anaconda.Kickstart != "" {
			to.Kickstart = &int.Kickstart{Contents: e.from.Installer.Anaconda.Kickstart}
		}

		if len(e.from.Installer.Anaconda.EnabledModules) > 0 {
			to.Modules.Enable = make([]string, len(e.from.Installer.Anaconda.EnabledModules))

			for i, module := range e.from.Installer.Anaconda.EnabledModules {
				to.Modules.Enable[i] = string(module)
			}
		}

		if len(e.from.Installer.Anaconda.DisabledModules) > 0 {
			to.Modules.Disable = make([]string, len(e.from.Installer.Anaconda.DisabledModules))
			for i, module := range e.from.Installer.Anaconda.DisabledModules {
				to.Modules.Disable[i] = string(module)
			}
		}

		if to.Modules.Enable == nil && to.Modules.Disable == nil {
			to.Modules = nil
		}
	}

	if e.from.Installer.CoreOS != nil {
		installationDevice = e.from.Installer.CoreOS.InstallationDevice
	}

	return installationDevice, to
}

func (e *InternalExporter) exportRegistration() (*int.RHSMCustomization, *int.FDOCustomization) {
	if e.from.Registration == nil {
		return nil, nil
	}

	var fdo *int.FDOCustomization
	if e.from.Registration.RegistrationFDO != nil {
		fdo = &int.FDOCustomization{}
		fdo.DiMfgStringTypeMacIface = e.from.Registration.RegistrationFDO.DiMfgStringTypeMacIface
		fdo.DiunPubKeyHash = e.from.Registration.RegistrationFDO.DiunPubKeyHash
		fdo.DiunPubKeyInsecure = strconv.FormatBool(e.from.Registration.RegistrationFDO.DiunPubKeyInsecure)
		fdo.DiunPubKeyRootCerts = e.from.Registration.RegistrationFDO.DiunPubKeyRootCerts
		fdo.ManufacturingServerURL = e.from.Registration.RegistrationFDO.ManufacturingServerURL
	}

	var rhsm *int.RHSMCustomization
	if e.from.Registration.RegistrationRedHat != nil {
		rhsm = &int.RHSMCustomization{
			Config: &int.RHSMConfig{
				SubscriptionManager: &int.SubManConfig{
					RHSMConfig:      &int.SubManRHSMConfig{},
					RHSMCertdConfig: &int.SubManRHSMCertdConfig{},
				},
				DNFPlugins: &int.SubManDNFPluginsConfig{
					ProductID:           &int.DNFPluginConfig{},
					SubscriptionManager: &int.DNFPluginConfig{},
				},
			},
		}
		e.log.Println("TODO: RHSM customization not yet implemented")
	}

	return rhsm, fdo
}
