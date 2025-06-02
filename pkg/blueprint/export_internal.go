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
func (e *InternalExporter) Export() error {
	to := &int.Blueprint{}

	// Create monotonic incremental version number based on miliseconds
	to.Version = int64ToVersion(uint64(time.Now().UTC().UnixMilli()))

	to.Name = e.from.Name.Get()
	to.Description = e.from.Description
	to.Packages = e.exportPackages()
	to.EnabledModules = e.exportModules()
	to.Groups = e.exportGroups()
	to.Containers = e.exportContainers()
	to.Customizations = e.exportCustomizations()
	to.Distro = e.from.Distribution
	to.Arch = e.from.Architecture.String()

	e.to = to
	return e.log.Errors()
}

func (e *InternalExporter) Result() *int.Blueprint {
	return e.to
}

func (e *InternalExporter) exportPackages() []int.Package {
	if e.from.DNF == nil || e.from.DNF.Packages == nil {
		return nil
	}

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
	if e.from.DNF == nil || e.from.DNF.Groups == nil {
		return nil
	}

	var s []int.Group
	for _, pkg := range e.from.DNF.Groups {
		s = append(s, int.Group{
			Name: pkg,
		})
	}

	return s
}

func (e *InternalExporter) exportModules() []int.EnabledModule {
	if e.from.DNF == nil || e.from.DNF.Modules == nil {
		return nil
	}

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
			TLSVerify:    ptr.ToNilEmpty(container.TLSVerify),
			LocalStorage: container.LocalStorage,
		})
	}

	return s
}

func (e *InternalExporter) exportCustomizations() *int.Customizations {
	to := &int.Customizations{}

	to.Hostname = ptr.ToNilEmpty(e.from.Hostname)
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
	to.OpenSCAP = e.exportOpenSCAP()
	to.Ignition = e.exportIgnition()
	to.Files, to.Directories = e.exportFSNodes()
	to.Repositories, to.RPM = e.exportRepositories()
	if e.from.FIPS != nil {
		to.FIPS = ptr.ToNilEmpty(e.from.FIPS.Enabled)
	}
	to.CACerts = e.exportCACerts()

	return to
}

func (e *InternalExporter) exportKernel() *int.KernelCustomization {
	if e.from.Kernel == nil {
		return nil
	}

	to := &int.KernelCustomization{}
	to.Name = e.from.Kernel.Package
	if len(e.from.Kernel.CmdlineAppend) > 0 {
		to.Append = strings.Join(e.from.Kernel.CmdlineAppend, " ")
	}

	return ptr.EmptyToNil(to)
}

func (e *InternalExporter) exportUserCustomization() []int.UserCustomization {
	if e.from.Accounts == nil || e.from.Accounts.Users == nil {
		return nil
	}

	var s []int.UserCustomization
	for _, u := range e.from.Accounts.Users {
		uc := int.UserCustomization{}
		uc.Name = u.Name
		uc.Description = ptr.ToNilEmpty(u.Description)
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = ptr.ToNilEmpty(u.SSHKeys[0])
		} else if len(u.SSHKeys) > 1 {
			uc.Key = ptr.ToNilEmpty(u.SSHKeys[0])
			e.log.Printf("only one ssh key supported for user: %s", u.Name)
		}
		uc.Home = ptr.ToNilEmpty(u.Home)
		uc.Shell = ptr.ToNilEmpty(u.Shell)
		uc.Groups = u.Groups
		if u.UID != 0 {
			uc.UID = ptr.ToNilEmpty(u.UID)
		}
		if u.GID != 0 {
			uc.GID = ptr.ToNilEmpty(u.GID)
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
	if e.from.Accounts == nil || e.from.Accounts.Groups == nil {
		return nil
	}

	var s []int.GroupCustomization
	for _, g := range e.from.Accounts.Groups {
		gc := int.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = ptr.ToNilEmpty(g.GID)
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
	to.Timezone = ptr.ToNilEmpty(e.from.Timedate.Timezone)
	to.NTPServers = e.from.Timedate.NTPServers

	return to
}

func (e *InternalExporter) exportLocaleCustomization() *int.LocaleCustomization {
	if e.from.Locale == nil {
		return nil
	}

	to := &int.LocaleCustomization{}
	if len(e.from.Locale.Keyboards) > 0 {
		to.Keyboard = ptr.ToNilEmpty(e.from.Locale.Keyboards[0])
		if len(e.from.Locale.Keyboards) > 1 {
			e.log.Println("only one keyboard layout supported, selecting first one")
		}
	}
	to.Languages = e.from.Locale.Languages

	return to
}

func (e *InternalExporter) exportFirewallCustomization() *int.FirewallCustomization {
	if e.from.Network == nil || e.from.Network.Firewall == nil || len(e.from.Network.Firewall.Services) == 0 {
		return nil
	}

	to := &int.FirewallCustomization{
		Services: &int.FirewallServicesCustomization{},
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

	return to
}

func (e *InternalExporter) exportSystemdCustomization() *int.ServicesCustomization {
	if e.from.Systemd == nil {
		return nil
	}

	to := &int.ServicesCustomization{}
	to.Enabled = e.from.Systemd.Enabled
	to.Disabled = e.from.Systemd.Disabled
	to.Masked = e.from.Systemd.Masked

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
		e.log.Printf("error parsing device size %s: %v", e.from.Storage.Minsize, err)
	} else {
		to.MinSize = size.Bytes()
	}

	for i, p := range e.from.Storage.Partitions {
		pp, pl, pb, err := p.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse partition %d: %v", i, err)
			continue
		}

		if pp.Type == PartTypePlain {
			size, err := ParseSize(pp.Minsize)
			if err != nil {
				e.log.Printf("error parsing parition size %q: %v", pp.Minsize, err)
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
				e.log.Printf("error parsing volume size %q: %v", pl.Minsize, err)
				continue
			}

			part := &int.PartitionCustomization{
				Type:    "lvm",
				MinSize: size.Bytes(),
				VGCustomization: int.VGCustomization{
					Name: pl.Name,
				},
			}

			for _, lv := range pl.LogicalVolumes {
				lvSize, err := ParseSize(lv.Minsize)
				if err != nil {
					e.log.Printf("error parsing LVM size %q: %v", lv.Minsize, err)
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
				e.log.Printf("error parsing BTRFS size %q: %v", pb.Minsize, err)
				continue
			}

			part := &int.PartitionCustomization{
				Type:                     "btrfs",
				MinSize:                  size.Bytes(),
				BtrfsVolumeCustomization: int.BtrfsVolumeCustomization{},
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

	if to.Type == "" && to.MinSize == 0 && len(to.Partitions) == 0 {
		return nil
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
			for _, module := range e.from.Installer.Anaconda.EnabledModules {
				to.Modules.Enable = append(to.Modules.Enable, string(module))
			}
		}

		if len(e.from.Installer.Anaconda.DisabledModules) > 0 {
			for _, module := range e.from.Installer.Anaconda.DisabledModules {
				to.Modules.Disable = append(to.Modules.Disable, string(module))
			}
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
	r := e.from.Registration

	var fdo *int.FDOCustomization
	if r.RegistrationFDO != nil {
		fdo = &int.FDOCustomization{}
		fdo.DiMfgStringTypeMacIface = r.RegistrationFDO.DiMfgStringTypeMacIface
		fdo.DiunPubKeyHash = r.RegistrationFDO.DiunPubKeyHash
		fdo.DiunPubKeyInsecure = strconv.FormatBool(r.RegistrationFDO.DiunPubKeyInsecure)
		fdo.DiunPubKeyRootCerts = r.RegistrationFDO.DiunPubKeyRootCerts
		fdo.ManufacturingServerURL = r.RegistrationFDO.ManufacturingServerURL
	}

	var rhsm *int.RHSMCustomization
	if r.RegistrationRedHat != nil && r.RegistrationRedHat.RegistrationRHSM != nil {
		rhsm = &int.RHSMCustomization{
			Config: &int.RHSMConfig{
				SubscriptionManager: &int.SubManConfig{
					RHSMConfig: &int.SubManRHSMConfig{
						ManageRepos:          r.RegistrationRedHat.RegistrationRHSM.RepositoryManagement,
						AutoEnableYumPlugins: r.RegistrationRedHat.RegistrationRHSM.AutoEnable,
					},
					RHSMCertdConfig: &int.SubManRHSMCertdConfig{
						AutoRegistration: r.RegistrationRedHat.RegistrationRHSM.AutoRegistration,
					},
				},
				DNFPlugins: &int.SubManDNFPluginsConfig{
					ProductID: &int.DNFPluginConfig{
						Enabled: r.RegistrationRedHat.RegistrationRHSM.ProductPluginEnabled,
					},
					SubscriptionManager: &int.DNFPluginConfig{
						Enabled: r.RegistrationRedHat.RegistrationRHSM.Enabled,
					},
				},
			},
		}
		e.log.Println("Registration not converted")
	}

	return ptr.EmptyToNil(rhsm), ptr.EmptyToNil(fdo)
}

func (e *InternalExporter) exportOpenSCAP() *int.OpenSCAPCustomization {
	if e.from.OpenSCAP == nil {
		return nil
	}

	to := &int.OpenSCAPCustomization{}
	to.DataStream = e.from.OpenSCAP.Datastream
	to.ProfileID = e.from.OpenSCAP.ProfileID

	if e.from.OpenSCAP.Tailoring != nil {
		tp, tj, err := e.from.OpenSCAP.Tailoring.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse tailoring: %v", err)
		}

		if tj.JSONProfileID != "" || tj.JSONFilePath != "" {
			to.JSONTailoring = &int.OpenSCAPJSONTailoringCustomizations{
				ProfileID: tj.JSONProfileID,
				Filepath:  tj.JSONFilePath,
			}
		} else if len(tp.Selected) > 0 || len(tp.Unselected) > 0 {
			to.Tailoring = &int.OpenSCAPTailoringCustomizations{
				Selected:   tp.Selected,
				Unselected: tp.Unselected,
			}
		} else {
			e.log.Printf("could not parse tailoring: %v", err)
		}
	}

	return ptr.EmptyToNil(to)
}

func (e *InternalExporter) exportIgnition() *int.IgnitionCustomization {
	if e.from.Ignition == nil {
		return nil
	}

	to := &int.IgnitionCustomization{}
	iu, ie, err := e.from.Ignition.SelectUnion()
	if err != nil {
		e.log.Printf("could not parse ignition: %v", err)
	}
	if ie.Text != "" {
		to.Embedded = &int.EmbeddedIgnitionCustomization{
			Config: ie.Text,
		}
	} else if iu.URL != "" {
		to.FirstBoot = &int.FirstBootIgnitionCustomization{
			ProvisioningURL: iu.URL,
		}
	} else {
		e.log.Printf("could not parse ignition: %v", err)
	}

	return ptr.EmptyToNil(to)
}

// parseUGID parses a user/group ID from a string. It returns the
// user/group ID as an int64 if it is a number, or the string itself
// if it is not a number. If the string is empty, it returns nil.
func parseUGID(s string) any {
	if s == "" {
		return nil
	}

	if i, err := strconv.ParseInt(s, 10, 0); err == nil {
		return i
	}

	return s
}

func (e *InternalExporter) exportFSNodes() ([]int.FileCustomization, []int.DirectoryCustomization) {
	if e.from.FSNodes == nil {
		return nil, nil
	}

	var files []int.FileCustomization
	var dirs []int.DirectoryCustomization
	for i, node := range e.from.FSNodes {

		switch node.Type {
		case FSNodeFile, "":
			var contents string
			var err error
			if node.Contents != nil {
				contents, err = node.Contents.String()
				if err != nil {
					e.log.Printf("could not parse contents of node %d: %v", i, err)
				}
			}

			files = append(files, int.FileCustomization{
				Path:  node.Path,
				User:  parseUGID(node.User),
				Group: parseUGID(node.Group),
				Mode:  strconv.FormatInt(int64(node.Mode), 8),
				Data:  contents,
			})
		case FSNodeDir:
			dirs = append(dirs, int.DirectoryCustomization{
				Path:          node.Path,
				User:          parseUGID(node.User),
				Group:         parseUGID(node.Group),
				Mode:          strconv.FormatInt(int64(node.Mode), 8),
				EnsureParents: node.EnsureParents,
			})
		default:
			e.log.Printf("unknown node type %d: %q", i, node.Type)
			continue
		}
	}

	return files, dirs
}

func (e *InternalExporter) exportRepositories() ([]int.RepositoryCustomization, *int.RPMCustomization) {
	if e.from.DNF == nil || e.from.DNF.Repositories == nil {
		return nil, nil
	}

	var repos []int.RepositoryCustomization
	var rpm *int.RPMCustomization
	for _, repo := range e.from.DNF.Repositories {
		burl, bmeta, bmirror, err := repo.Source.SelectUnion()
		if err != nil {
			e.log.Printf("missing source for repository %q: %v", repo.ID, err)
			continue
		}

		if repo.Usage == nil {
			repo.Usage = &DNFRepoUsage{}
		}

		repos = append(repos, int.RepositoryCustomization{
			Id:             repo.ID,
			Name:           repo.Name,
			BaseURLs:       burl.URLs,
			Metalink:       bmeta.Metalink,
			Mirrorlist:     bmirror.Mirrorlist,
			Priority:       ptr.ToNilEmpty(repo.Priority),
			Enabled:        ptr.Or(repo.Usage.Configure, true),
			SSLVerify:      ptr.Or(repo.SSLVerify, true),
			GPGKeys:        repo.GPGKeys,
			GPGCheck:       repo.GPGCheck,
			RepoGPGCheck:   repo.GPGCheckRepo,
			ModuleHotfixes: ptr.ToNilEmpty(repo.ModuleHotfixes),
			Filename:       repo.Filename,
			InstallFrom:    ptr.FromOr(repo.Usage.Install, true),
		})

		if len(repo.GPGKeys) > 0 {
			if rpm == nil {
				rpm = &int.RPMCustomization{
					ImportKeys: &int.RPMImportKeys{},
				}
			}

			rpm.ImportKeys.Files = append(rpm.ImportKeys.Files, repo.GPGKeys...)
		}
	}

	return repos, ptr.EmptyToNil(rpm)
}

func (e *InternalExporter) exportCACerts() *int.CACustomization {
	if e.from.CACerts == nil {
		return nil
	}

	var to *int.CACustomization
	if len(e.from.CACerts) > 0 {
		to = &int.CACustomization{}
		for _, cert := range e.from.CACerts {
			if cert.PEM == "" {
				continue
			}

			to.PEMCerts = append(to.PEMCerts, cert.PEM)
		}
	}

	return to
}
