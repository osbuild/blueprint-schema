package conv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	ubp "github.com/osbuild/blueprint-schema/pkg/ubp"
	bp "github.com/osbuild/blueprint/pkg/blueprint"
)

// InternalImporter is used to convert a blueprint to the internal representation.
type InternalImporter struct {
	from *bp.Blueprint
}

func NewInternalImporter(inputBlueprint *bp.Blueprint) *InternalImporter {
	return &InternalImporter{
		from: inputBlueprint,
	}
}

// Import converts the internal representation to the blueprint.
func (e *InternalImporter) Import() (*ubp.Blueprint, error) {
	to := &ubp.Blueprint{}
	var errs []error
	var err error

	to.Accounts, err = e.importAccounts()
	if err != nil {
		errs = append(errs, err)
	}

	to.Architecture, err = e.importArchitecture()
	if err != nil {
		errs = append(errs, err)
	}

	to.CACerts, err = e.importCACerts()
	if err != nil {
		errs = append(errs, err)
	}

	to.Containers, err = e.importContainers()
	if err != nil {
		errs = append(errs, err)
	}

	to.DNF, err = e.importDNF()
	if err != nil {
		errs = append(errs, err)
	}

	to.Description = e.from.Description
	to.Distribution = e.from.Distro

	to.FIPS, err = e.importFIPS()
	if err != nil {
		errs = append(errs, err)
	}

	to.FSNodes, err = e.importFSNodes()
	if err != nil {
		errs = append(errs, err)
	}

	if e.from.Customizations != nil {
		to.Hostname = ptr.ValueOrEmpty(e.from.Customizations.Hostname)
	}

	to.Ignition, err = e.importIgnition()
	if err != nil {
		errs = append(errs, err)
	}

	to.Installer, err = e.importInstaller()
	if err != nil {
		errs = append(errs, err)
	}

	to.Kernel, err = e.importKernel()
	if err != nil {
		errs = append(errs, err)
	}

	to.Locale, err = e.importLocale()
	if err != nil {
		errs = append(errs, err)
	}

	to.Name = e.from.Name

	to.Network, err = e.importNetwork()
	if err != nil {
		errs = append(errs, err)
	}

	to.OpenSCAP, err = e.importOpenSCAP()
	if err != nil {
		errs = append(errs, err)
	}

	to.Registration, err = e.importRegistration()
	if err != nil {
		errs = append(errs, err)
	}

	to.Storage, err = e.importStorage()
	if err != nil {
		errs = append(errs, err)
	}

	to.Systemd, err = e.importSystemd()
	if err != nil {
		errs = append(errs, err)
	}

	to.Timedate, err = e.importTimedate()
	if err != nil {
		errs = append(errs, err)
	}

	if e.from.Minimal {
		errs = append(errs, fmt.Errorf("minimal flag is not supported"))
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importArchitecture() (ubp.Arch, error) {
	if e.from.Arch == "" {
		return ubp.ArchUnset, nil
	}

	result, err := ubp.ParseArch(e.from.Arch)
	if err != nil {
		return ubp.ArchUnset, fmt.Errorf("error parsing architecture %q: %v", e.from.Arch, err)
	}

	return result, nil
}

func (e *InternalImporter) importDNF() (ubp.DNF, error) {
	to := ubp.DNF{}
	var errs []error
	var err error

	to.Packages, err = e.importPackages()
	if err != nil {
		errs = append(errs, err)
	}

	to.Modules, err = e.importModules()
	if err != nil {
		errs = append(errs, err)
	}

	to.Groups, err = e.importGroups()
	if err != nil {
		errs = append(errs, err)
	}

	to.Repositories, err = e.importRepositories()
	if err != nil {
		errs = append(errs, err)
	}

	if e.from.Customizations != nil && e.from.Customizations.RPM != nil && e.from.Customizations.RPM.ImportKeys != nil {
		for _, keyFile := range e.from.Customizations.RPM.ImportKeys.Files {
			to.ImportKeys = append(to.ImportKeys, strings.TrimPrefix(keyFile, "file://"))
		}
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importPackages() ([]string, error) {
	// Combine packages and modules into a single slice.
	s := make([]string, 0, len(e.from.Packages)+len(e.from.Modules))
	for _, pkg := range e.from.Packages {
		if pkg.Version == "*" {
			pkg.Version = ""
		}
		s = append(s, joinNonEmpty("-", pkg.Name, pkg.Version))
	}
	for _, pkg := range e.from.Modules {
		if pkg.Version == "*" {
			pkg.Version = ""
		}
		s = append(s, joinNonEmpty("-", pkg.Name, pkg.Version))
	}

	if len(s) == 0 {
		return nil, nil // omitzero
	}

	return s, nil
}

func (e *InternalImporter) importModules() ([]string, error) {
	if e.from.EnabledModules == nil {
		return nil, nil
	}

	s := make([]string, len(e.from.EnabledModules))
	for i, pkg := range e.from.EnabledModules {
		if pkg.Stream != "" {
			s[i] = fmt.Sprintf("%s:%s", pkg.Name, pkg.Stream)
		} else {
			s[i] = pkg.Name
		}
	}

	return s, nil
}

func (e *InternalImporter) importGroups() ([]string, error) {
	if e.from.Groups == nil {
		return nil, nil
	}

	s := make([]string, len(e.from.Groups))
	for i, group := range e.from.Groups {
		s[i] = group.Name
	}

	return s, nil
}

func (e *InternalImporter) importRepositories() ([]ubp.DNFRepository, error) {
	if e.from.Customizations == nil || e.from.Customizations.Repositories == nil {
		return nil, nil
	}

	var errs []error
	repos := make([]ubp.DNFRepository, len(e.from.Customizations.Repositories))
	for i, repo := range e.from.Customizations.Repositories {
		repos[i] = ubp.DNFRepository{
			Name:           repo.Name,
			ID:             repo.Id,
			Filename:       repo.Filename,
			GPGCheck:       repo.GPGCheck,
			GPGCheckRepo:   repo.RepoGPGCheck,
			GPGKeys:        repo.GPGKeys,
			ModuleHotfixes: ptr.ValueOr(repo.ModuleHotfixes, false),
			Priority:       repo.Priority,
			TLSVerify:      repo.SSLVerify,
			Usage: ubp.DnfRepositoryUsage{
				Configure: repo.Enabled,
				Install:   &repo.InstallFrom,
			},
		}

		if repo.BaseURLs != nil {
			repos[i].Source = ubp.DNFSourceFromBaseURLs(ubp.DNFSourceBaseURLs{
				URLs: repo.BaseURLs,
			})
		} else if repo.Metalink != "" {
			repos[i].Source = ubp.DNFSourceFromMetalink(ubp.DNFSourceMetalink{
				Metalink: repo.Metalink,
			})
		} else if repo.Mirrorlist != "" {
			repos[i].Source = ubp.DNFSourceFromMirrorlist(ubp.DNFSourceMirrorlist{
				Mirrorlist: repo.Mirrorlist,
			})
		} else {
			errs = append(errs, fmt.Errorf("repository %q has no source defined", repo.Id))
			continue
		}
	}

	return repos, errors.Join(errs...)
}

func (e *InternalImporter) importContainers() ([]ubp.Container, error) {
	if e.from.Containers == nil {
		return nil, nil
	}

	containers := make([]ubp.Container, len(e.from.Containers))
	for i, container := range e.from.Containers {
		containers[i] = ubp.Container{
			Name:         container.Name,
			LocalStorage: container.LocalStorage,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
		}
	}

	return containers, nil
}

func (e *InternalImporter) importKernel() (ubp.Kernel, error) {
	if e.from.Customizations == nil || e.from.Customizations.Kernel == nil {
		return ubp.Kernel{}, nil
	}

	r := ubp.Kernel{
		Package: e.from.Customizations.Kernel.Name,
	}

	if len(e.from.Customizations.Kernel.Append) > 0 {
		r.CmdlineAppend = strings.Split(e.from.Customizations.Kernel.Append, " ")
	}

	return r, nil
}

func (e *InternalImporter) importAccounts() (ubp.Accounts, error) {
	if e.from.Customizations == nil {
		return ubp.Accounts{}, nil
	}

	to := ubp.Accounts{}
	for _, user := range e.from.Customizations.User {
		u := ubp.AccountsUsers{
			Name:                user.Name,
			Description:         ptr.ValueOrEmpty(user.Description),
			Home:                ptr.ValueOrEmpty(user.Home),
			UID:                 ptr.ValueOr(user.UID, 0),
			GID:                 ptr.ValueOr(user.GID, 0),
			Groups:              user.Groups,
			Password:            user.Password,
			Expires:             ubp.NewIntEpochDays(ptr.ValueOrEmpty(user.ExpireDate)),
			ForcePasswordChange: user.ForcePasswordReset,
			Shell:               ptr.ValueOrEmpty(user.Shell),
		}

		if user.Key != nil {
			u.SSHKeys = []string{*user.Key}
		}

		to.Users = append(to.Users, u)
	}

	for _, group := range e.from.Customizations.Group {
		g := ubp.AccountsGroups{
			Name: group.Name,
			GID:  ptr.ValueOr(group.GID, 0),
		}

		to.Groups = append(to.Groups, g)
	}

	return to, nil
}

func (e *InternalImporter) importCACerts() ([]ubp.CACert, error) {
	if e.from.Customizations == nil || e.from.Customizations.CACerts == nil || e.from.Customizations.CACerts.PEMCerts == nil {
		return nil, nil
	}

	caCerts := make([]ubp.CACert, len(e.from.Customizations.CACerts.PEMCerts))
	for i, cert := range e.from.Customizations.CACerts.PEMCerts {
		caCerts[i] = ubp.CACert{
			PEM: cert,
		}
	}

	return caCerts, nil
}

func (e *InternalImporter) importFIPS() (ubp.FIPS, error) {
	if e.from.Customizations == nil || e.from.Customizations.FIPS == nil {
		return ubp.FIPS{}, nil
	}

	fips := ubp.FIPS{
		Enabled: ptr.ValueOr(e.from.Customizations.FIPS, false),
	}

	return fips, nil
}

func (e *InternalImporter) importFSNodes() ([]ubp.FSNode, error) {
	if e.from.Customizations == nil {
		return nil, nil
	}

	var errs []error

	if e.from.Customizations.Filesystem != nil {
		errs = append(errs, fmt.Errorf("filesystem customizations are not converted, ignoring"))
	}

	var res []ubp.FSNode
	for _, file := range e.from.Customizations.Files {
		mode, err := ubp.ParseFSNodeMode(file.Mode)
		if err != nil && file.Mode != "" {
			errs = append(errs, fmt.Errorf("error parsing file mode %q for file %q: %v, using default", file.Mode, file.Path, err))
		}

		if mode == ubp.DefaultFileFSNodeMode {
			mode = ubp.UnsetFSNodeMode
		}

		n := ubp.FSNode{
			Type:  ubp.FSNodeFile,
			Path:  file.Path,
			User:  parseUGIDany(file.User),
			Group: parseUGIDany(file.Group),
			Mode:  mode,
		}

		if file.Data != "" {
			n.Contents = ubp.FSNodeContentsFromText(ubp.FSNodeContentsText{
				Text: file.Data,
			})
		}

		res = append(res, n)
	}

	for _, dir := range e.from.Customizations.Directories {
		mode, err := ubp.ParseFSNodeMode(dir.Mode)
		if err != nil && dir.Mode != "" {
			errs = append(errs, fmt.Errorf("error parsing file mode %q for dir %q: %v, using default", dir.Mode, dir.Path, err))
		}

		if mode == ubp.DefaultDirFSNodeMode {
			mode = ubp.UnsetFSNodeMode
		}

		n := ubp.FSNode{
			Type:          ubp.FSNodeDir,
			Path:          dir.Path,
			User:          parseUGIDany(dir.User),
			Group:         parseUGIDany(dir.Group),
			Mode:          mode,
			EnsureParents: dir.EnsureParents,
		}

		res = append(res, n)
	}

	if len(res) == 0 {
		return nil, errors.Join(errs...) // omitzero
	}

	return res, errors.Join(errs...)
}

func (e *InternalImporter) importIgnition() (ubp.Ignition, error) {
	if e.from.Customizations == nil || e.from.Customizations.Ignition == nil {
		return ubp.Ignition{}, nil
	}

	var res ubp.Ignition
	if e.from.Customizations.Ignition.FirstBoot != nil {
		res = ubp.IgnitionFromURL(ubp.IgnitionURL{
			URL: e.from.Customizations.Ignition.FirstBoot.ProvisioningURL,
		})
	}

	if e.from.Customizations.Ignition.Embedded != nil {
		res = ubp.IgnitionFromText(ubp.IgnitionText{
			Text: e.from.Customizations.Ignition.Embedded.Config,
		})
	}

	return res, nil
}

func (e *InternalImporter) importInstaller() (ubp.Installer, error) {
	if e.from.Customizations == nil {
		return ubp.Installer{}, nil
	}

	to := ubp.Installer{}
	if e.from.Customizations.Installer != nil {
		to.Anaconda = ubp.InstallerAnaconda{
			Unattended:   e.from.Customizations.Installer.Unattended,
			SudoNOPASSWD: e.from.Customizations.Installer.SudoNopasswd,
		}

		if e.from.Customizations.Installer.Kickstart != nil {
			to.Anaconda.Kickstart = e.from.Customizations.Installer.Kickstart.Contents
		}

		if e.from.Customizations.Installer.Modules != nil {
			for _, m := range e.from.Customizations.Installer.Modules.Enable {
				if pm := ubp.ParseAnacondaModule(m); pm != "" {
					to.Anaconda.EnabledModules = append(to.Anaconda.EnabledModules, pm)
				}
			}

			for _, m := range e.from.Customizations.Installer.Modules.Disable {
				if pm := ubp.ParseAnacondaModule(m); pm != "" {
					to.Anaconda.DisabledModules = append(to.Anaconda.DisabledModules, pm)
				}
			}
		}
	}

	if e.from.Customizations.InstallationDevice != "" {
		to.CoreOS = ubp.InstallerCoreOS{
			InstallationDevice: e.from.Customizations.InstallationDevice,
		}
	}

	return to, nil
}

func (e *InternalImporter) importLocale() (ubp.Locale, error) {
	if e.from.Customizations == nil || e.from.Customizations.Locale == nil {
		return ubp.Locale{}, nil
	}

	to := ubp.Locale{}
	to.Languages = append(to.Languages, e.from.Customizations.Locale.Languages...)
	if e.from.Customizations.Locale.Keyboard != nil {
		to.Keyboards = []string{*e.from.Customizations.Locale.Keyboard}
	}

	return to, nil
}

func (e *InternalImporter) importNetwork() (ubp.Network, error) {
	to := ubp.Network{
		Firewall: ubp.NetworkFirewall{},
	}
	var errs []error

	if e.from.Customizations == nil || e.from.Customizations.Firewall == nil {
		return to, nil
	}

	if e.from.Customizations.Firewall.Services != nil {
		for _, srv := range e.from.Customizations.Firewall.Services.Enabled {
			ns := ubp.FirewallService{
				Service: srv,
			}
			if service := ubp.NetworkServiceFromService(ns); service != nil {
				to.Firewall.Services = append(to.Firewall.Services, *service)
			}
		}

		for _, srv := range e.from.Customizations.Firewall.Services.Disabled {
			ns := ubp.FirewallService{
				Service: srv,
				Enabled: ptr.To(ubp.FirewallEnabled(false)),
			}
			if service := ubp.NetworkServiceFromService(ns); service != nil {
				to.Firewall.Services = append(to.Firewall.Services, *service)
			}
		}
	}

	for _, port := range e.from.Customizations.Firewall.Ports {
		if strings.Contains(port, "-") {
			fromTo, err := ubp.ParseFirewalldFromTo(port)
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing firewall port range %q: %v, ignoring", port, err))
				continue
			}

			ns := ubp.NetworkServiceFromFromTo(fromTo)
			to.Firewall.Services = append(to.Firewall.Services, *ns)
			continue
		} else {
			firewallPort, err := ubp.ParseFirewalldPort(port)
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing firewall port %q: %v, ignoring", port, err))
				continue
			}

			ns := ubp.NetworkServiceFromPort(firewallPort)
			to.Firewall.Services = append(to.Firewall.Services, *ns)
		}
	}

	if e.from.Customizations.Firewall.Zones != nil {
		errs = append(errs, fmt.Errorf("firewall zones are not supported, ignoring"))
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importOpenSCAP() (ubp.OpenSCAP, error) {
	if e.from.Customizations == nil || e.from.Customizations.OpenSCAP == nil {
		return ubp.OpenSCAP{}, nil
	}

	var errs []error
	to := ubp.OpenSCAP{
		ProfileID:  e.from.Customizations.OpenSCAP.ProfileID,
		Datastream: e.from.Customizations.OpenSCAP.DataStream,
	}

	if e.from.Customizations.OpenSCAP.PolicyID != "" {
		// https://github.com/osbuild/blueprint-schema/issues/29
		errs = append(errs, fmt.Errorf("policy ID %q is not supported, ignoring", e.from.Customizations.OpenSCAP.PolicyID))
	}

	if e.from.Customizations.OpenSCAP.JSONTailoring != nil {
		to.Tailoring = ubp.OpenSCAPTailoringFromJSON(ubp.TailoringJSON{
			JSONProfileID: e.from.Customizations.OpenSCAP.JSONTailoring.ProfileID,
			JSONFilePath:  e.from.Customizations.OpenSCAP.JSONTailoring.Filepath,
		})
	}

	if e.from.Customizations.OpenSCAP.Tailoring != nil {
		to.Tailoring = ubp.OpenSCAPTailoringFromProfiles(ubp.TailoringProfiles{
			Selected:   e.from.Customizations.OpenSCAP.Tailoring.Selected,
			Unselected: e.from.Customizations.OpenSCAP.Tailoring.Unselected,
		})
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importRegistration() (ubp.Registration, error) {
	if e.from.Customizations == nil {
		return ubp.Registration{}, nil
	}

	var errs []error
	to := ubp.Registration{}
	if e.from.Customizations.RHSM != nil && e.from.Customizations.RHSM.Config != nil {
		to.RegistrationRedHat = ubp.RegistrationRedHat{
			RegistrationRHSM: ubp.RegistrationRHSM{},
		}

		if e.from.Customizations.RHSM.Config.SubscriptionManager != nil {
			if e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMConfig != nil {
				to.RegistrationRedHat.RegistrationRHSM.AutoEnable = e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMConfig.AutoEnableYumPlugins
				to.RegistrationRedHat.RegistrationRHSM.RepositoryManagement = e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMConfig.ManageRepos
			}

			if e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMCertdConfig != nil {
				to.RegistrationRedHat.RegistrationRHSM.AutoRegistration = e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMCertdConfig.AutoRegistration
			}
		}

		if e.from.Customizations.RHSM.Config.DNFPlugins != nil {
			to.RegistrationRedHat.RegistrationRHSM.Enabled = e.from.Customizations.RHSM.Config.DNFPlugins.SubscriptionManager.Enabled
			to.RegistrationRedHat.RegistrationRHSM.ProductPluginEnabled = e.from.Customizations.RHSM.Config.DNFPlugins.ProductID.Enabled
		}
	}

	if e.from.Customizations.FDO != nil {
		var insecure bool
		_, err := fmt.Sscanf(e.from.Customizations.FDO.DiunPubKeyInsecure, "%t", &insecure)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot parse DiunPubKeyInsecure %q: %v, using default false", e.from.Customizations.FDO.DiunPubKeyInsecure, err))
		}

		to.RegistrationFDO = ubp.RegistrationFDO{
			ManufacturingServerURL:  e.from.Customizations.FDO.ManufacturingServerURL,
			DiMfgStringTypeMacIface: e.from.Customizations.FDO.DiMfgStringTypeMacIface,
			DiunPubKeyHash:          e.from.Customizations.FDO.DiunPubKeyHash,
			DiunPubKeyInsecure:      insecure,
			DiunPubKeyRootCerts:     e.from.Customizations.FDO.DiunPubKeyRootCerts,
		}
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importStorage() (ubp.Storage, error) {
	if e.from.Customizations == nil || e.from.Customizations.Disk == nil {
		return ubp.Storage{}, nil
	}

	var errs []error
	to := ubp.Storage{
		Type: ubp.StorageType(e.from.Customizations.Disk.Type),
	}

	if e.from.Customizations.Disk.MinSize > 0 {
		to.Minsize = ubp.ToByteSize(e.from.Customizations.Disk.MinSize)
	}

	for _, part := range e.from.Customizations.Disk.Partitions {
		switch strings.ToLower(part.Type) {
		case "plain":
			fst, err := ubp.ParseFSType(part.FSType)
			if err != nil {
				errs = append(errs, fmt.Errorf("error parsing filesystem type %q for partition %q: %v, using default", part.FSType, part.Name, err))
			}
			np := ubp.PartitionPlain{
				Type:       ubp.PartTypePlain,
				FSType:     fst,
				Label:      part.Label,
				Minsize:    ubp.ToByteSize(part.MinSize),
				Mountpoint: part.Mountpoint,
			}
			to.Partitions = append(to.Partitions, ubp.StoragePartitionFromPlain(np))
		case "btrfs":
			np := ubp.PartitionBTRFS{
				Type:    ubp.PartTypeBTRFS,
				Minsize: ubp.ToByteSize(part.MinSize),
			}
			for _, sv := range part.Subvolumes {
				nsv := ubp.PartitionSubvolumes{
					Name:       sv.Name,
					Mountpoint: sv.Mountpoint,
				}
				np.Subvolumes = append(np.Subvolumes, nsv)
			}
			to.Partitions = append(to.Partitions, ubp.StoragePartitionFromBTRFS(np))
		case "lvm":
			np := ubp.PartitionLVM{
				Type:    ubp.PartTypeLVM,
				Name:    part.Name,
				Minsize: ubp.ToByteSize(part.MinSize),
			}
			for _, lv := range part.LogicalVolumes {
				fst, err := ubp.ParseFSType(lv.FSType)
				if err != nil {
					errs = append(errs, fmt.Errorf("error parsing filesystem type %q for lv %q: %v, using default", lv.FSType, lv.Name, err))
				}
				nlv := ubp.PartitionLV{
					Name:       lv.Name,
					Label:      lv.Label,
					FSType:     fst,
					Minsize:    ubp.ToByteSize(lv.MinSize),
					Mountpoint: lv.Mountpoint,
				}
				np.LogicalVolumes = append(np.LogicalVolumes, nlv)
			}
			to.Partitions = append(to.Partitions, ubp.StoragePartitionFromLVM(np))
		}
	}

	return to, errors.Join(errs...)
}

func (e *InternalImporter) importSystemd() (ubp.Systemd, error) {
	if e.from.Customizations == nil || e.from.Customizations.Services == nil {
		return ubp.Systemd{}, nil
	}

	to := ubp.Systemd{
		Enabled:  e.from.Customizations.Services.Enabled,
		Disabled: e.from.Customizations.Services.Disabled,
		Masked:   e.from.Customizations.Services.Masked,
	}

	return to, nil
}

func (e *InternalImporter) importTimedate() (ubp.TimeDate, error) {
	if e.from.Customizations == nil || e.from.Customizations.Timezone == nil {
		return ubp.TimeDate{}, nil
	}

	to := ubp.TimeDate{
		Timezone:   ptr.ValueOrEmpty(e.from.Customizations.Timezone.Timezone),
		NTPServers: e.from.Customizations.Timezone.NTPServers,
	}

	return to, nil
}
