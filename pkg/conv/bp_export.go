package conv

import (
	"strconv"
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	ubp "github.com/osbuild/blueprint-schema/pkg/ubp"
	bp "github.com/osbuild/blueprint/pkg/blueprint"
)

// InternalExporter is used to convert a blueprint to the internal representation.
type InternalExporter struct {
	from *ubp.Blueprint
	log  *errs
}

func NewInternalExporter(inputBlueprint *ubp.Blueprint) *InternalExporter {
	return &InternalExporter{
		from: inputBlueprint,
		log:  newErrorCollector(),
	}
}

// ExportInternal converts the blueprint to the internal representation.
func (e *InternalExporter) Export() (*bp.Blueprint, error) {
	to := &bp.Blueprint{}

	if e.from == nil {
		return nil, nil
	}

	// Create monotonic incremental version number based on miliseconds
	to.Version = int64ToVersion(uint64(time.Now().UTC().UnixMilli()))

	to.Name = e.from.Name
	to.Description = e.from.Description
	to.Packages = e.exportPackages()
	to.EnabledModules = e.exportModules()
	to.Groups = e.exportGroups()
	to.Containers = e.exportContainers()
	to.Customizations = e.exportCustomizations()
	to.Distro = e.from.Distribution
	to.Arch = e.from.Architecture.String()

	return to, e.log.Errors()
}

func (e *InternalExporter) exportPackages() []bp.Package {
	if e.from.DNF == nil || e.from.DNF.Packages == nil {
		return nil
	}

	var s []bp.Package
	for _, pkg := range e.from.DNF.Packages {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, bp.Package{
			Name:    p[0],
			Version: p[1],
		})
	}

	return s
}

func (e *InternalExporter) exportGroups() []bp.Group {
	if e.from.DNF == nil || e.from.DNF.Groups == nil {
		return nil
	}

	var s []bp.Group
	for _, pkg := range e.from.DNF.Groups {
		s = append(s, bp.Group{
			Name: pkg,
		})
	}

	return s
}

func (e *InternalExporter) exportModules() []bp.EnabledModule {
	if e.from.DNF == nil || e.from.DNF.Modules == nil {
		return nil
	}

	var s []bp.EnabledModule
	for _, pkg := range e.from.DNF.Modules {
		p := splitStringEmptyN(pkg, ":", 2)

		s = append(s, bp.EnabledModule{
			Name:   p[0],
			Stream: p[1],
		})
	}

	return s
}

func (e *InternalExporter) exportContainers() []bp.Container {
	var s []bp.Container

	for _, container := range e.from.Containers {
		s = append(s, bp.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return s
}

func (e *InternalExporter) exportCustomizations() *bp.Customizations {
	to := &bp.Customizations{}

	to.Hostname = ptr.ToNilIfEmpty(e.from.Hostname)
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
		to.FIPS = ptr.ToNilIfEmpty(e.from.FIPS.Enabled)
	}
	to.CACerts = e.exportCACerts()

	return to
}

func (e *InternalExporter) exportKernel() *bp.KernelCustomization {
	if e.from.Kernel == nil {
		return nil
	}

	to := &bp.KernelCustomization{}
	to.Name = e.from.Kernel.Package
	if len(e.from.Kernel.CmdlineAppend) > 0 {
		to.Append = strings.Join(e.from.Kernel.CmdlineAppend, " ")
	}

	return ptr.EmptyToNil(to)
}

func (e *InternalExporter) exportUserCustomization() []bp.UserCustomization {
	if e.from.Accounts == nil || e.from.Accounts.Users == nil {
		return nil
	}

	var s []bp.UserCustomization
	for _, u := range e.from.Accounts.Users {
		uc := bp.UserCustomization{}
		uc.Name = u.Name
		uc.Description = ptr.ToNilIfEmpty(u.Description)
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = ptr.ToNilIfEmpty(u.SSHKeys[0])
		} else if len(u.SSHKeys) > 1 {
			uc.Key = ptr.ToNilIfEmpty(u.SSHKeys[0])
			e.log.Printf("only one ssh key supported for user: %s", u.Name)
		}
		uc.Home = ptr.ToNilIfEmpty(u.Home)
		uc.Shell = ptr.ToNilIfEmpty(u.Shell)
		uc.Groups = u.Groups
		if u.UID != 0 {
			uc.UID = ptr.ToNilIfEmpty(u.UID)
		}
		if u.GID != 0 {
			uc.GID = ptr.ToNilIfEmpty(u.GID)
		}
		if u.Expires != nil {
			uc.ExpireDate = ptr.ToNilIfEmpty(u.Expires.Days())
		}
		if u.ForcePasswordChange != nil {
			uc.ForcePasswordReset = u.ForcePasswordChange
		}

		s = append(s, uc)
	}

	return s
}

func (e *InternalExporter) exportGroupCustomization() []bp.GroupCustomization {
	if e.from.Accounts == nil || e.from.Accounts.Groups == nil {
		return nil
	}

	var s []bp.GroupCustomization
	for _, g := range e.from.Accounts.Groups {
		gc := bp.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = ptr.ToNilIfEmpty(g.GID)
		}
		s = append(s, gc)
	}

	return s
}

func (e *InternalExporter) exportTimezoneCustomization() *bp.TimezoneCustomization {
	if e.from.Timedate == nil {
		return nil
	}

	to := &bp.TimezoneCustomization{}
	to.Timezone = ptr.ToNilIfEmpty(e.from.Timedate.Timezone)
	to.NTPServers = e.from.Timedate.NTPServers

	return to
}

func (e *InternalExporter) exportLocaleCustomization() *bp.LocaleCustomization {
	if e.from.Locale == nil {
		return nil
	}

	to := &bp.LocaleCustomization{}
	if len(e.from.Locale.Keyboards) > 0 {
		to.Keyboard = ptr.ToNilIfEmpty(e.from.Locale.Keyboards[0])
		if len(e.from.Locale.Keyboards) > 1 {
			e.log.Println("only one keyboard layout supported, selecting first one")
		}
	}
	to.Languages = append(to.Languages, e.from.Locale.Languages...)

	return to
}

func (e *InternalExporter) exportFirewallCustomization() *bp.FirewallCustomization {
	if e.from.Network == nil || e.from.Network.Firewall == nil || len(e.from.Network.Firewall.Services) == 0 {
		return nil
	}

	to := &bp.FirewallCustomization{
		Services: &bp.FirewallServicesCustomization{},
	}
	for i, s := range e.from.Network.Firewall.Services {
		fs, fp, fft, err := s.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse network service %d: %v", i, err)
			continue
		}

		if fs.Service != "" {
			if fs.Enabled == nil || *fs.Enabled {
				to.Services.Enabled = append(to.Services.Enabled, fs.Service)
			} else {
				to.Services.Disabled = append(to.Services.Disabled, fs.Service)
			}
		} else if fp.Port != 0 {
			srv := ubp.PortProtoToFirewalld(fp.Port, fp.Protocol)

			if fp.Enabled == nil || *fp.Enabled {
				to.Ports = append(to.Ports, srv)
			} else {
				e.log.Printf("network service %d error: port number %d cannot be disabled", i, fp.Port)
				continue
			}
		} else if fft.From != 0 && fft.To != 0 {
			srv := ubp.PortsProtoToFirewalld(fft.From, fft.To, fp.Protocol)

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

func (e *InternalExporter) exportSystemdCustomization() *bp.ServicesCustomization {
	if e.from.Systemd == nil {
		return nil
	}

	to := &bp.ServicesCustomization{}
	to.Enabled = e.from.Systemd.Enabled
	to.Disabled = e.from.Systemd.Disabled
	to.Masked = e.from.Systemd.Masked

	return to
}

func (e *InternalExporter) exportStorage() *bp.DiskCustomization {
	if e.from.Storage == nil {
		return nil
	}

	to := &bp.DiskCustomization{}
	to.Type = e.from.Storage.Type.String()
	to.MinSize = e.from.Storage.Minsize.Bytes()

	for i, p := range e.from.Storage.Partitions {
		pp, pl, pb, err := p.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse partition %d: %v", i, err)
			continue
		}

		if pp.Type == ubp.PartTypePlain {
			part := &bp.PartitionCustomization{
				Type:    "plain",
				MinSize: pp.Minsize.Bytes(),
				FilesystemTypedCustomization: bp.FilesystemTypedCustomization{
					Label:      pp.Label,
					Mountpoint: pp.Mountpoint,
					FSType:     pp.FSType.String(),
				},
			}

			to.Partitions = append(to.Partitions, *part)
		} else if pl.Type == ubp.PartTypeLVM {
			part := &bp.PartitionCustomization{
				Type:    "lvm",
				MinSize: pl.Minsize.Bytes(),
				VGCustomization: bp.VGCustomization{
					Name: pl.Name,
				},
			}

			for _, lv := range pl.LogicalVolumes {
				lvc := bp.LVCustomization{
					Name:    lv.Name,
					MinSize: lv.Minsize.Bytes(),
					FilesystemTypedCustomization: bp.FilesystemTypedCustomization{
						Label:      lv.Label,
						Mountpoint: lv.Mountpoint,
						FSType:     lv.FSType.String(),
					},
				}
				part.LogicalVolumes = append(part.LogicalVolumes, lvc)
			}

			to.Partitions = append(to.Partitions, *part)
		} else if pb.Type == ubp.PartTypeBTRFS {
			part := &bp.PartitionCustomization{
				Type:                     "btrfs",
				MinSize:                  pb.Minsize.Bytes(),
				BtrfsVolumeCustomization: bp.BtrfsVolumeCustomization{},
			}

			for _, sv := range pb.Subvolumes {
				svc := bp.BtrfsSubvolumeCustomization{
					Name:       sv.Name,
					Mountpoint: sv.Mountpoint,
				}
				part.Subvolumes = append(part.Subvolumes, svc)
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

func (e *InternalExporter) exportInstaller() (string, *bp.InstallerCustomization) {
	var installationDevice string
	if e.from.Installer == nil {
		return installationDevice, nil
	}

	var to *bp.InstallerCustomization
	if e.from.Installer.Anaconda != nil {
		to = &bp.InstallerCustomization{
			Modules: &bp.AnacondaModules{},
		}
		to.Unattended = e.from.Installer.Anaconda.Unattended
		to.SudoNopasswd = e.from.Installer.Anaconda.SudoNOPASSWD
		if e.from.Installer.Anaconda.Kickstart != "" {
			to.Kickstart = &bp.Kickstart{Contents: e.from.Installer.Anaconda.Kickstart}
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

func (e *InternalExporter) exportRegistration() (*bp.RHSMCustomization, *bp.FDOCustomization) {
	if e.from.Registration == nil {
		return nil, nil
	}
	r := e.from.Registration

	var fdo *bp.FDOCustomization
	if r.RegistrationFDO != nil {
		fdo = &bp.FDOCustomization{}
		fdo.DiMfgStringTypeMacIface = r.RegistrationFDO.DiMfgStringTypeMacIface
		fdo.DiunPubKeyHash = r.RegistrationFDO.DiunPubKeyHash
		fdo.DiunPubKeyInsecure = strconv.FormatBool(r.RegistrationFDO.DiunPubKeyInsecure)
		fdo.DiunPubKeyRootCerts = r.RegistrationFDO.DiunPubKeyRootCerts
		fdo.ManufacturingServerURL = r.RegistrationFDO.ManufacturingServerURL
	}

	var rhsm *bp.RHSMCustomization
	if r.RegistrationRedHat != nil && r.RegistrationRedHat.RegistrationRHSM != nil {
		rhsm = &bp.RHSMCustomization{
			Config: &bp.RHSMConfig{
				SubscriptionManager: &bp.SubManConfig{
					RHSMConfig: &bp.SubManRHSMConfig{
						ManageRepos:          r.RegistrationRedHat.RegistrationRHSM.RepositoryManagement,
						AutoEnableYumPlugins: r.RegistrationRedHat.RegistrationRHSM.AutoEnable,
					},
					RHSMCertdConfig: &bp.SubManRHSMCertdConfig{
						AutoRegistration: r.RegistrationRedHat.RegistrationRHSM.AutoRegistration,
					},
				},
				DNFPlugins: &bp.SubManDNFPluginsConfig{
					ProductID: &bp.DNFPluginConfig{
						Enabled: r.RegistrationRedHat.RegistrationRHSM.ProductPluginEnabled,
					},
					SubscriptionManager: &bp.DNFPluginConfig{
						Enabled: r.RegistrationRedHat.RegistrationRHSM.Enabled,
					},
				},
			},
		}
		e.log.Println("registration not converted")
	}

	return ptr.EmptyToNil(rhsm), ptr.EmptyToNil(fdo)
}

func (e *InternalExporter) exportOpenSCAP() *bp.OpenSCAPCustomization {
	if e.from.OpenSCAP == nil {
		return nil
	}

	to := &bp.OpenSCAPCustomization{}
	to.DataStream = e.from.OpenSCAP.Datastream
	to.ProfileID = e.from.OpenSCAP.ProfileID

	if e.from.OpenSCAP.Tailoring != nil {
		tp, tj, err := e.from.OpenSCAP.Tailoring.SelectUnion()
		if err != nil {
			e.log.Printf("could not parse tailoring: %v", err)
		}

		if tj.JSONProfileID != "" || tj.JSONFilePath != "" {
			to.JSONTailoring = &bp.OpenSCAPJSONTailoringCustomizations{
				ProfileID: tj.JSONProfileID,
				Filepath:  tj.JSONFilePath,
			}
		} else if len(tp.Selected) > 0 || len(tp.Unselected) > 0 {
			to.Tailoring = &bp.OpenSCAPTailoringCustomizations{
				Selected:   tp.Selected,
				Unselected: tp.Unselected,
			}
		} else {
			e.log.Printf("could not parse tailoring: %v", err)
		}
	}

	return ptr.EmptyToNil(to)
}

func (e *InternalExporter) exportIgnition() *bp.IgnitionCustomization {
	if e.from.Ignition == nil {
		return nil
	}

	to := &bp.IgnitionCustomization{}
	iu, ie, err := e.from.Ignition.SelectUnion()
	if err != nil {
		e.log.Printf("could not parse ignition: %v", err)
	}
	if ie.Text != "" {
		to.Embedded = &bp.EmbeddedIgnitionCustomization{
			Config: ie.Text,
		}
	} else if iu.URL != "" {
		to.FirstBoot = &bp.FirstBootIgnitionCustomization{
			ProvisioningURL: iu.URL,
		}
	} else {
		e.log.Printf("could not parse ignition: %v", err)
	}

	return ptr.EmptyToNil(to)
}

func (e *InternalExporter) exportFSNodes() ([]bp.FileCustomization, []bp.DirectoryCustomization) {
	if e.from.FSNodes == nil {
		return nil, nil
	}

	var files []bp.FileCustomization
	var dirs []bp.DirectoryCustomization
	for i, node := range e.from.FSNodes {

		switch node.Type {
		case ubp.FSNodeFile, "":
			var contents string
			var err error
			if node.Contents != nil {
				contents, err = node.Contents.String()
				if err != nil {
					e.log.Printf("could not parse contents of node %d: %v, contents skipped", i, err)
				}
			}

			if node.State != nil && node.State.IsAbsent() {
				e.log.Printf("fs node %d is marked as absent, unsupported", i)
				continue
			}

			fc := bp.FileCustomization{
				Path:  node.Path,
				User:  parseUGIDstr(node.User),
				Group: parseUGIDstr(node.Group),
				Data:  contents,
			}
			mode := strconv.FormatInt(int64(node.Mode), 8)
			if mode != "0" {
				fc.Mode = mode
			}

			files = append(files, fc)
		case ubp.FSNodeDir:
			fc := bp.DirectoryCustomization{
				Path:          node.Path,
				User:          parseUGIDstr(node.User),
				Group:         parseUGIDstr(node.Group),
				Mode:          strconv.FormatInt(int64(node.Mode), 8),
				EnsureParents: node.EnsureParents,
			}
			mode := strconv.FormatInt(int64(node.Mode), 8)
			if mode != "0" {
				fc.Mode = mode
			}

			dirs = append(dirs, fc)
		default:
			e.log.Printf("unknown node type %d: %q", i, node.Type)
			continue
		}
	}

	return files, dirs
}

func (e *InternalExporter) exportRepositories() ([]bp.RepositoryCustomization, *bp.RPMCustomization) {
	if e.from.DNF == nil || e.from.DNF.Repositories == nil {
		return nil, nil
	}

	var repos []bp.RepositoryCustomization
	var rpm *bp.RPMCustomization
	for _, repo := range e.from.DNF.Repositories {
		burl, bmeta, bmirror, err := repo.Source.SelectUnion()
		if err != nil {
			e.log.Printf("missing source for repository %q: %v", repo.ID, err)
			continue
		}

		var configure, install *bool
		if repo.Usage != nil {
			configure = repo.Usage.Configure
			install = repo.Usage.Install
		}

		repos = append(repos, bp.RepositoryCustomization{
			Id:             repo.ID,
			Name:           repo.Name,
			BaseURLs:       burl.URLs,
			Metalink:       bmeta.Metalink,
			Mirrorlist:     bmirror.Mirrorlist,
			Priority:       ptr.ToNilIfEmpty(repo.Priority),
			Enabled:        configure,
			SSLVerify:      repo.SSLVerify,
			GPGKeys:        repo.GPGKeys,
			GPGCheck:       repo.GPGCheck,
			RepoGPGCheck:   repo.GPGCheckRepo,
			ModuleHotfixes: ptr.ToNilIfEmpty(repo.ModuleHotfixes),
			Filename:       repo.Filename,
			InstallFrom:    ptr.ValueOr(install, true),
		})

		if len(e.from.DNF.ImportKeys) > 0 {
			if rpm == nil {
				rpm = &bp.RPMCustomization{
					ImportKeys: &bp.RPMImportKeys{
						Files: e.from.DNF.ImportKeys,
					},
				}
			}
		}
	}

	return repos, ptr.EmptyToNil(rpm)
}

func (e *InternalExporter) exportCACerts() *bp.CACustomization {
	if e.from.CACerts == nil {
		return nil
	}

	var to *bp.CACustomization
	if len(e.from.CACerts) > 0 {
		to = &bp.CACustomization{}
		for _, cert := range e.from.CACerts {
			if cert.PEM == "" {
				continue
			}

			to.PEMCerts = append(to.PEMCerts, cert.PEM)
		}
	}

	return to
}
