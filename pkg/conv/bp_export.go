package conv

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	ubp "github.com/osbuild/blueprint-schema/pkg/ubp"
	bp "github.com/osbuild/blueprint/pkg/blueprint"
)

// InternalExporter is used to convert a blueprint to the internal representation.
type InternalExporter struct {
	from *ubp.Blueprint
}

func NewInternalExporter(inputBlueprint *ubp.Blueprint) *InternalExporter {
	return &InternalExporter{
		from: inputBlueprint,
	}
}

// ExportInternal converts the blueprint to the internal representation.
func (e *InternalExporter) Export() (*bp.Blueprint, error) {
	to := &bp.Blueprint{}
	var errs []error
	var err error

	if e.from == nil {
		return nil, nil
	}

	to.Name = e.from.Name
	to.Description = e.from.Description

	to.Packages, err = e.exportPackages()
	if err != nil {
		errs = append(errs, err)
	}

	to.EnabledModules, err = e.exportModules()
	if err != nil {
		errs = append(errs, err)
	}

	to.Groups, err = e.exportGroups()
	if err != nil {
		errs = append(errs, err)
	}

	to.Containers, err = e.exportContainers()
	if err != nil {
		errs = append(errs, err)
	}

	to.Customizations, err = e.exportCustomizations()
	if err != nil {
		errs = append(errs, err)
	}
	to.Distro = e.from.Distribution
	to.Arch = e.from.Architecture.String()

	return to, errors.Join(errs...)
}

func (e *InternalExporter) exportPackages() ([]bp.Package, error) {
	var s []bp.Package
	for _, pkg := range e.from.DNF.Packages {
		// It is not possible to reliably detect version of a package with a dash in its name,
		// let's do best effort and split it at the second last dash and issue a warning.
		en, vr := splitEnVr(pkg)

		s = append(s, bp.Package{
			Name:    en, // Epoch + Name
			Version: vr, // Version + Release
		})
	}

	return s, nil
}

func (e *InternalExporter) exportGroups() ([]bp.Group, error) {
	var s []bp.Group
	for _, pkg := range e.from.DNF.Groups {
		s = append(s, bp.Group{
			Name: pkg,
		})
	}
	return s, nil
}

func (e *InternalExporter) exportModules() ([]bp.EnabledModule, error) {
	var s []bp.EnabledModule
	for _, pkg := range e.from.DNF.Modules {
		p := splitStringEmptyN(pkg, ":", 2)

		s = append(s, bp.EnabledModule{
			Name:   p[0],
			Stream: p[1],
		})
	}

	return s, nil
}

func (e *InternalExporter) exportContainers() ([]bp.Container, error) {
	var s []bp.Container
	for _, container := range e.from.Containers {
		s = append(s, bp.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return s, nil
}

func (e *InternalExporter) exportCustomizations() (*bp.Customizations, error) {
	to := &bp.Customizations{}
	var errs []error
	var err error

	to.Hostname = ptr.ToNilIfEmpty(e.from.Hostname)
	to.FIPS = ptr.ToNilIfEmpty(e.from.FIPS.Enabled)

	to.Kernel, err = e.exportKernel()
	if err != nil {
		errs = append(errs, err)
	}

	to.User, err = e.exportUserCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Group, err = e.exportGroupCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Timezone, err = e.exportTimezoneCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Locale, err = e.exportLocaleCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Firewall, err = e.exportFirewallCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Services, err = e.exportSystemdCustomization()
	if err != nil {
		errs = append(errs, err)
	}

	to.Disk, err = e.exportStorage()
	if err != nil {
		errs = append(errs, err)
	}

	to.InstallationDevice, to.Installer, err = e.exportInstaller()
	if err != nil {
		errs = append(errs, err)
	}

	to.RHSM, to.FDO, err = e.exportRegistration()
	if err != nil {
		errs = append(errs, err)
	}

	to.OpenSCAP, err = e.exportOpenSCAP()
	if err != nil {
		errs = append(errs, err)
	}

	to.Ignition, err = e.exportIgnition()
	if err != nil {
		errs = append(errs, err)
	}

	to.Files, to.Directories, err = e.exportFSNodes()
	if err != nil {
		errs = append(errs, err)
	}

	to.Repositories, to.RPM, err = e.exportRepositories()
	if err != nil {
		errs = append(errs, err)
	}

	to.CACerts, err = e.exportCACerts()
	if err != nil {
		errs = append(errs, err)
	}

	return to, errors.Join(errs...)
}

func (e *InternalExporter) exportKernel() (*bp.KernelCustomization, error) {
	to := &bp.KernelCustomization{}

	to.Name = e.from.Kernel.Package
	if len(e.from.Kernel.CmdlineAppend) > 0 {
		to.Append = strings.Join(e.from.Kernel.CmdlineAppend, " ")
	}

	return ptr.EmptyToNil(to), nil
}

func (e *InternalExporter) exportUserCustomization() ([]bp.UserCustomization, error) {
	var s []bp.UserCustomization
	var err error

	for _, u := range e.from.Accounts.Users {
		uc := bp.UserCustomization{}
		uc.Name = u.Name
		uc.Description = ptr.ToNilIfEmpty(u.Description)
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = ptr.ToNilIfEmpty(u.SSHKeys[0])
		} else if len(u.SSHKeys) > 1 {
			uc.Key = ptr.ToNilIfEmpty(u.SSHKeys[0])
			err = fmt.Errorf("only one SSH key supported, selecting the first one for user %q", u.Name)
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

	return s, err
}

func (e *InternalExporter) exportGroupCustomization() ([]bp.GroupCustomization, error) {
	var s []bp.GroupCustomization

	for _, g := range e.from.Accounts.Groups {
		gc := bp.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = ptr.ToNilIfEmpty(g.GID)
		}
		s = append(s, gc)
	}

	return s, nil
}

func (e *InternalExporter) exportTimezoneCustomization() (*bp.TimezoneCustomization, error) {
	to := &bp.TimezoneCustomization{}

	to.Timezone = ptr.ToNilIfEmpty(e.from.Timedate.Timezone)
	to.NTPServers = e.from.Timedate.NTPServers

	return to, nil
}

func (e *InternalExporter) exportLocaleCustomization() (*bp.LocaleCustomization, error) {
	to := &bp.LocaleCustomization{}
	var err error

	if len(e.from.Locale.Keyboards) > 0 {
		to.Keyboard = ptr.ToNilIfEmpty(e.from.Locale.Keyboards[0])
		if len(e.from.Locale.Keyboards) > 1 {
			err = fmt.Errorf("only one keyboard supported, selecting the first one: %q", e.from.Locale.Keyboards[0])
		}
	}
	to.Languages = append(to.Languages, e.from.Locale.Languages...)

	return to, err
}

func (e *InternalExporter) exportFirewallCustomization() (*bp.FirewallCustomization, error) {
	to := &bp.FirewallCustomization{
		Services: &bp.FirewallServicesCustomization{},
	}

	for i, s := range e.from.Network.Firewall.Services {
		fs, fp, fft, err := s.SelectUnion()
		if err != nil {
			return nil, fmt.Errorf("network service %d parsing error: %v", i, err)
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
				return nil, fmt.Errorf("network service %d error: port number %d cannot be disabled", i, fp.Port)
			}
		} else if fft.From != 0 && fft.To != 0 {
			srv := ubp.PortsProtoToFirewalld(fft.From, fft.To, fp.Protocol)

			if fft.Enabled == nil || *fft.Enabled {
				to.Ports = append(to.Ports, srv)
			} else {
				return nil, fmt.Errorf("network service %d error: port number %d cannot be disabled", i, fp.Port)
			}
		} else {
			return nil, fmt.Errorf("network service %d error: one of service, port or from and to present", i)
		}
	}

	return to, nil
}

func (e *InternalExporter) exportSystemdCustomization() (*bp.ServicesCustomization, error) {
	to := &bp.ServicesCustomization{}

	to.Enabled = e.from.Systemd.Enabled
	to.Disabled = e.from.Systemd.Disabled
	to.Masked = e.from.Systemd.Masked

	return to, nil
}

func (e *InternalExporter) exportStorage() (*bp.DiskCustomization, error) {
	to := &bp.DiskCustomization{}

	to.Type = e.from.Storage.Type.String()
	to.MinSize = e.from.Storage.Minsize.Bytes()

	for i, p := range e.from.Storage.Partitions {
		pp, pl, pb, err := p.SelectUnion()
		if err != nil {
			return nil, fmt.Errorf("could not parse partition %d: %v", i, err)
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
			return nil, fmt.Errorf("unknown partition type %q", pl.Type)
		}
	}

	return to, nil
}

func (e *InternalExporter) exportInstaller() (string, *bp.InstallerCustomization, error) {
	var installationDevice string
	to := &bp.InstallerCustomization{
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

	installationDevice = e.from.Installer.CoreOS.InstallationDevice

	return installationDevice, to, nil
}

func (e *InternalExporter) exportRegistration() (*bp.RHSMCustomization, *bp.FDOCustomization, error) {
	r := e.from.Registration

	fdo := &bp.FDOCustomization{}
	fdo.DiMfgStringTypeMacIface = r.RegistrationFDO.DiMfgStringTypeMacIface
	fdo.DiunPubKeyHash = r.RegistrationFDO.DiunPubKeyHash
	fdo.DiunPubKeyInsecure = strconv.FormatBool(r.RegistrationFDO.DiunPubKeyInsecure)
	fdo.DiunPubKeyRootCerts = r.RegistrationFDO.DiunPubKeyRootCerts
	fdo.ManufacturingServerURL = r.RegistrationFDO.ManufacturingServerURL

	rhsm := &bp.RHSMCustomization{
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

	emptyRHSM := &bp.RHSMCustomization{
		Config: &bp.RHSMConfig{
			DNFPlugins: &bp.SubManDNFPluginsConfig{
				ProductID:           &bp.DNFPluginConfig{},
				SubscriptionManager: &bp.DNFPluginConfig{},
			},
			SubscriptionManager: &bp.SubManConfig{
				RHSMConfig:      &bp.SubManRHSMConfig{},
				RHSMCertdConfig: &bp.SubManRHSMCertdConfig{},
			},
		},
	}

	if reflect.DeepEqual(rhsm.Config.DNFPlugins, emptyRHSM.Config.DNFPlugins) {
		rhsm.Config.DNFPlugins = nil // omitzero
	}
	if reflect.DeepEqual(rhsm.Config.SubscriptionManager, emptyRHSM.Config.SubscriptionManager) {
		rhsm.Config.SubscriptionManager = nil // omitzero
	}

	return ptr.EmptyToNil(rhsm), ptr.EmptyToNil(fdo), nil
}

func (e *InternalExporter) exportOpenSCAP() (*bp.OpenSCAPCustomization, error) {
	to := &bp.OpenSCAPCustomization{}

	to.DataStream = e.from.OpenSCAP.Datastream
	to.ProfileID = e.from.OpenSCAP.ProfileID

	if e.from.OpenSCAP.Tailoring != nil {
		tp, tj, err := e.from.OpenSCAP.Tailoring.SelectUnion()
		if err != nil {
			return nil, fmt.Errorf("could not parse tailoring: %v", err)
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
			return nil, fmt.Errorf("could not parse tailoring: %v", err)
		}
	}

	return ptr.EmptyToNil(to), nil
}

func (e *InternalExporter) exportIgnition() (*bp.IgnitionCustomization, error) {
	to := &bp.IgnitionCustomization{}

	iu, ie, err := e.from.Ignition.SelectUnion()
	if err != nil {
		return nil, fmt.Errorf("could not parse ignition: %v", err)
	}

	if ie.Text != "" {
		to.Embedded = &bp.EmbeddedIgnitionCustomization{
			Config: ie.Text,
		}
	} else if iu.URL != "" {
		to.FirstBoot = &bp.FirstBootIgnitionCustomization{
			ProvisioningURL: iu.URL,
		}
	}

	// XXX: use ptr.EmptyToNil to omit empty struct
	return to, nil
}

func (e *InternalExporter) exportFSNodes() ([]bp.FileCustomization, []bp.DirectoryCustomization, error) {
	if e.from.FSNodes == nil {
		return nil, nil, nil
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
					return nil, nil, fmt.Errorf("could not parse contents of node %d: %v, contents skipped", i, err)
				}
			}

			if node.State.IsAbsent() {
				return nil, nil, fmt.Errorf("fs node %d is marked as absent, unsupported", i)
			}

			fc := bp.FileCustomization{
				Path:  node.Path,
				User:  parseUGIDstr(node.User),
				Group: parseUGIDstr(node.Group),
				Data:  contents,
			}

			if node.Mode != ubp.UnsetFSNodeMode {
				fc.Mode = strconv.FormatInt(int64(node.Mode), 8)
			}

			files = append(files, fc)
		case ubp.FSNodeDir:
			fc := bp.DirectoryCustomization{
				Path:          node.Path,
				User:          parseUGIDstr(node.User),
				Group:         parseUGIDstr(node.Group),
				EnsureParents: node.EnsureParents,
			}

			if node.Mode != ubp.UnsetFSNodeMode {
				fc.Mode = strconv.FormatInt(int64(node.Mode), 8)
			}

			dirs = append(dirs, fc)
		default:
			return nil, nil, fmt.Errorf("unknown node type %d: %q", i, node.Type)
		}
	}

	return files, dirs, nil
}

func (e *InternalExporter) exportRepositories() ([]bp.RepositoryCustomization, *bp.RPMCustomization, error) {
	var repos []bp.RepositoryCustomization

	for _, repo := range e.from.DNF.Repositories {
		burl, bmeta, bmirror, err := repo.Source.SelectUnion()
		if err != nil {
			return nil, nil, fmt.Errorf("missing source for repository %q: %v", repo.ID, err)
		}

		var configure, install *bool
		configure = repo.Usage.Configure
		install = repo.Usage.Install

		repos = append(repos, bp.RepositoryCustomization{
			Id:             repo.ID,
			Name:           repo.Name,
			BaseURLs:       burl.URLs,
			Metalink:       bmeta.Metalink,
			Mirrorlist:     bmirror.Mirrorlist,
			Priority:       repo.Priority,
			Enabled:        configure,
			SSLVerify:      repo.TLSVerify,
			GPGKeys:        repo.GPGKeys,
			GPGCheck:       repo.GPGCheck,
			RepoGPGCheck:   repo.GPGCheckRepo,
			ModuleHotfixes: ptr.ToNilIfEmpty(repo.ModuleHotfixes),
			Filename:       repo.Filename,
			InstallFrom:    ptr.ValueOr(install, true),
		})
	}

	var rpm *bp.RPMCustomization

	if len(e.from.DNF.ImportKeys) > 0 {
		rpm = &bp.RPMCustomization{
			ImportKeys: &bp.RPMImportKeys{
				Files: e.from.DNF.ImportKeys,
			},
		}
	}

	return repos, rpm, nil
}

func (e *InternalExporter) exportCACerts() (*bp.CACustomization, error) {
	if e.from.CACerts == nil {
		return nil, nil
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

	return to, nil
}
