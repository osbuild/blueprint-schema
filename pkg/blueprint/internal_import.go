package blueprint

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	int "github.com/osbuild/blueprint/pkg/blueprint"
)

// InternalImporter is used to convert a blueprint to the internal representation.
type InternalImporter struct {
	from *int.Blueprint
	to   *Blueprint
	log  *logs
}

func NewInternalImporter(inputBlueprint *int.Blueprint) *InternalImporter {
	return &InternalImporter{
		from: inputBlueprint,
		log:  newCollector(),
	}
}

// Import converts the internal representation to the blueprint.
func (e *InternalImporter) Import() error {
	to := &Blueprint{}

	to.Accounts = e.importAccounts()
	to.Architecture = e.importArchitecture()
	to.CACerts = e.importCACerts()
	to.Containers = e.importContainers()
	to.DNF = e.importDNF()
	to.Description = e.from.Description
	to.Distribution = e.from.Distro
	to.FIPS = e.importFIPS()
	to.FSNodes = e.importFSNodes()
	if e.from.Customizations != nil {
		to.Hostname = ptr.ValueOrEmpty(e.from.Customizations.Hostname)
	}
	to.Ignition = e.importIgnition()
	to.Installer = e.importInstaller()
	to.Kernel = e.importKernel()
	to.Locale = e.importLocale()
	to.Name = e.from.Name
	to.Network = e.importNetwork()
	to.OpenSCAP = e.importOpenSCAP()
	to.Registration = e.importRegistration()
	to.Storage = e.importStorage()
	to.Systemd = e.importSystemd()
	to.Timedate = e.importTimedate()

	e.to = to
	return e.log.Errors()
}

func (e *InternalImporter) Result() *Blueprint {
	return e.to
}

func (e *InternalImporter) importArchitecture() Arch {
	if e.from.Arch == "" {
		return ArchUnset
	}

	result, err := ParseArch(e.from.Arch)
	if err != nil {
		e.log.Printf("error parsing architecture %q: %v", e.from.Arch, err)
	}

	return result
}

func (e *InternalImporter) importDNF() *DNF {
	if e.from.Customizations == nil || e.from.Customizations.RPM == nil {
		return nil
	}

	to := DNF{}
	to.Packages = e.importPackages()
	to.Modules = e.importModules()
	to.Groups = e.importGroups()
	to.Repositories = e.importRepositories()

	if e.from.Customizations.RPM.ImportKeys != nil {
		for _, keyFile := range e.from.Customizations.RPM.ImportKeys.Files {
			to.ImportKeys = append(to.ImportKeys, strings.TrimPrefix(keyFile, "file://"))
		}
	}

	if reflect.DeepEqual(to, DNF{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importPackages() []string {
	if e.from.Packages == nil {
		return nil
	}

	// Combine packages and modules into a single slice.
	s := make([]string, len(e.from.Packages)+len(e.from.Modules))
	for i, pkg := range e.from.Packages {
		s[i] = joinNonEmpty("-", pkg.Name, pkg.Version)
	}
	for i, pkg := range e.from.Modules {
		s[i] = joinNonEmpty("-", pkg.Name, pkg.Version)
	}

	return s
}

func (e *InternalImporter) importModules() []string {
	if e.from.EnabledModules == nil {
		return nil
	}

	s := make([]string, len(e.from.EnabledModules))
	for i, pkg := range e.from.EnabledModules {
		if pkg.Stream != "" {
			s[i] = fmt.Sprintf("%s:%s", pkg.Name, pkg.Stream)
		} else {
			s[i] = pkg.Name
		}
	}

	return s
}

func (e *InternalImporter) importGroups() []string {
	if e.from.Groups == nil {
		return nil
	}

	s := make([]string, len(e.from.Groups))
	for i, group := range e.from.Groups {
		s[i] = group.Name
	}

	return s
}

func (e *InternalImporter) importRepositories() []DNFRepository {
	if e.from.Customizations == nil || e.from.Customizations.Repositories == nil {
		return nil
	}

	repos := make([]DNFRepository, len(e.from.Customizations.Repositories))
	for i, repo := range e.from.Customizations.Repositories {
		repos[i] = DNFRepository{
			Name:           repo.Name,
			ID:             repo.Id,
			Filename:       repo.Filename,
			GPGCheck:       repo.GPGCheck,
			GPGCheckRepo:   repo.RepoGPGCheck,
			GPGKeys:        repo.GPGKeys,
			ModuleHotfixes: ptr.ValueOr(repo.ModuleHotfixes, false),
			Priority:       ptr.ValueOr(repo.Priority, 0),
			SSLVerify:      repo.SSLVerify,
			Usage: &DnfRepositoryUsage{
				Configure: repo.Enabled,
				Install:   &repo.InstallFrom,
			},
		}

		if repo.BaseURLs != nil {
			repos[i].Source = DNFSourceFromBaseURLs(DNFSourceBaseURLs{
				URLs: repo.BaseURLs,
			})
		} else if repo.Metalink != "" {
			repos[i].Source = DNFSourceFromMetalink(DNFSourceMetalink{
				Metalink: repo.Metalink,
			})
		} else if repo.Mirrorlist != "" {
			repos[i].Source = DNFSourceFromMirrorlist(DNFSourceMirrorlist{
				Mirrorlist: repo.Mirrorlist,
			})
		} else {
			e.log.Printf("repository %q has no source defined", repo.Id)
			continue
		}
	}

	return repos
}

func (e *InternalImporter) importContainers() []Container {
	if e.from.Containers == nil {
		return nil
	}

	containers := make([]Container, len(e.from.Containers))
	for i, container := range e.from.Containers {
		containers[i] = Container{
			Name:         container.Name,
			LocalStorage: container.LocalStorage,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
		}
	}

	return containers
}

func (e *InternalImporter) importKernel() *Kernel {
	if e.from.Customizations == nil || e.from.Customizations.Kernel == nil {
		return nil
	}

	r := &Kernel{
		Package: e.from.Customizations.Kernel.Name,
	}

	if len(e.from.Customizations.Kernel.Append) > 0 {
		r.CmdlineAppend = strings.Split(e.from.Customizations.Kernel.Append, " ")
	}

	return r
}

func (e *InternalImporter) importAccounts() *Accounts {
	if e.from.Customizations == nil {
		return nil
	}

	to := Accounts{}
	for _, user := range e.from.Customizations.User {
		u := AccountsUsers{
			Name:                user.Name,
			Description:         ptr.ValueOrEmpty(user.Description),
			Home:                ptr.ValueOrEmpty(user.Home),
			UID:                 ptr.ValueOr(user.UID, 0),
			GID:                 ptr.ValueOr(user.GID, 0),
			Groups:              user.Groups,
			Password:            user.Password,
			Expires:             NewIntEpochDays(ptr.ValueOrEmpty(user.ExpireDate)),
			ForcePasswordChange: user.ForcePasswordReset,
			Shell:               ptr.ValueOrEmpty(user.Shell),
		}

		if user.Key != nil {
			u.SSHKeys = []string{*user.Key}
		}

		to.Users = append(to.Users, u)
	}

	for _, group := range e.from.Customizations.Group {
		g := AccountsGroups{
			Name: group.Name,
			GID:  ptr.ValueOr(group.GID, 0),
		}

		to.Groups = append(to.Groups, g)
	}

	if reflect.DeepEqual(to, Accounts{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importCACerts() []CACert {
	if e.from.Customizations == nil || e.from.Customizations.CACerts == nil || e.from.Customizations.CACerts.PEMCerts == nil {
		return nil
	}

	caCerts := make([]CACert, len(e.from.Customizations.CACerts.PEMCerts))
	for i, cert := range e.from.Customizations.CACerts.PEMCerts {
		caCerts[i] = CACert{
			PEM: cert,
		}
	}

	return caCerts
}

func (e *InternalImporter) importFIPS() *FIPS {
	if e.from.Customizations == nil || e.from.Customizations.FIPS == nil {
		return nil
	}

	fips := FIPS{
		Enabled: ptr.ValueOr(e.from.Customizations.FIPS, false),
	}

	if reflect.DeepEqual(fips, FIPS{}) {
		return nil // omitzero
	}

	return &fips
}

func (e *InternalImporter) importFSNodes() []FSNode {
	if e.from.Customizations == nil {
		return nil
	}

	var res []FSNode
	for _, file := range e.from.Customizations.Files {
		mode, err := ParseFSNodeMode(file.Mode)
		if err != nil {
			e.log.Printf("error parsing file mode %q for file %q: %v, using default", file.Mode, file.Path, err)
		}

		n := FSNode{
			Type:  FSNodeFile,
			Path:  file.Path,
			User:  parseUGIDany(file.User),
			Group: parseUGIDany(file.Group),
			Mode:  mode,
		}

		if file.Data != "" {
			n.Contents = FSNodeContentsFromText(FSNodeContentsText{
				Text: file.Data,
			})
		}

		res = append(res, n)
	}

	for _, dir := range e.from.Customizations.Directories {
		mode, err := ParseFSNodeMode(dir.Mode)
		if err != nil {
			e.log.Printf("error parsing file mode %q for dir %q: %v, using default", dir.Mode, dir.Path, err)
		}

		n := FSNode{
			Type:          FSNodeDir,
			Path:          dir.Path,
			User:          parseUGIDany(dir.User),
			Group:         parseUGIDany(dir.Group),
			Mode:          mode,
			EnsureParents: dir.EnsureParents,
		}

		res = append(res, n)
	}
	return res
}

func (e *InternalImporter) importIgnition() *Ignition {
	if e.from.Customizations == nil || e.from.Customizations.Ignition == nil {
		return nil
	}

	var res *Ignition
	if e.from.Customizations.Ignition.FirstBoot != nil {
		res = IgnitionFromURL(IgnitionURL{
			URL: e.from.Customizations.Ignition.FirstBoot.ProvisioningURL,
		})
	}

	if e.from.Customizations.Ignition.Embedded != nil {
		res = IgnitionFromText(IgnitionText{
			Text: e.from.Customizations.Ignition.Embedded.Config,
		})
	}

	return res
}

func (e *InternalImporter) importInstaller() *Installer {
	if e.from.Customizations == nil || e.from.Customizations.Installer == nil {
		return nil
	}

	to := Installer{
		Anaconda: &InstallerAnaconda{
			Unattended:   e.from.Customizations.Installer.Unattended,
			SudoNOPASSWD: e.from.Customizations.Installer.SudoNopasswd,
		},
	}

	if e.from.Customizations.Installer.Kickstart != nil {
		to.Anaconda.Kickstart = e.from.Customizations.Installer.Kickstart.Contents
	}

	if e.from.Customizations.Installer.Modules != nil {
		for _, m := range e.from.Customizations.Installer.Modules.Enable {
			if pm := ParseAnacondaModule(m); pm != "" {
				to.Anaconda.EnabledModules = append(to.Anaconda.EnabledModules, pm)
			}
		}

		for _, m := range e.from.Customizations.Installer.Modules.Disable {
			if pm := ParseAnacondaModule(m); pm != "" {
				to.Anaconda.DisabledModules = append(to.Anaconda.DisabledModules, pm)
			}
		}
	}

	if e.from.Customizations.InstallationDevice != "" {
		to.CoreOS = &InstallerCoreOS{
			InstallationDevice: e.from.Customizations.InstallationDevice,
		}
	}

	if reflect.DeepEqual(to, Installer{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importLocale() *Locale {
	if e.from.Customizations == nil || e.from.Customizations.Locale == nil {
		return nil
	}

	to := Locale{}
	to.Languages = append(to.Languages, e.from.Customizations.Locale.Languages...)
	if e.from.Customizations.Locale.Keyboard != nil {
		to.Keyboards = []string{*e.from.Customizations.Locale.Keyboard}
	}

	if reflect.DeepEqual(to, Locale{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importNetwork() *Network {
	if e.from.Customizations == nil || e.from.Customizations.Firewall == nil {
		return nil
	}

	to := Network{
		Firewall: &NetworkFirewall{},
	}

	if e.from.Customizations.Firewall.Services != nil {
		for _, srv := range e.from.Customizations.Firewall.Services.Enabled {
			ns := FirewallService{
				Service: srv,
			}
			if service := NetworkServiceFromService(ns); service != nil {
				to.Firewall.Services = append(to.Firewall.Services, *service)
			}
		}
	}

	for _, port := range e.from.Customizations.Firewall.Ports {
		if strings.Contains(port, "-") {
			fromTo, err := ParseFirewalldFromTo(port)
			if err != nil {
				e.log.Printf("error parsing firewall port range %q: %v, ignoring", port, err)
				continue
			}

			ns := NetworkServiceFromFromTo(fromTo)
			to.Firewall.Services = append(to.Firewall.Services, *ns)
			continue
		} else {
			firewallPort, err := ParseFirewalldPort(port)
			if err != nil {
				e.log.Printf("error parsing firewall port %q: %v, ignoring", port, err)
				continue
			}

			ns := NetworkServiceFromPort(firewallPort)
			to.Firewall.Services = append(to.Firewall.Services, *ns)
		}
	}

	if e.from.Customizations.Firewall.Zones != nil {
		e.log.Printf("firewall zones are not supported, ignoring")
	}

	if reflect.DeepEqual(to.Firewall, NetworkFirewall{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importOpenSCAP() *OpenSCAP {
	if e.from.Customizations == nil || e.from.Customizations.OpenSCAP == nil {
		return nil
	}

	to := OpenSCAP{
		ProfileID:  e.from.Customizations.OpenSCAP.ProfileID,
		Datastream: e.from.Customizations.OpenSCAP.DataStream,
	}

	if e.from.Customizations.OpenSCAP.PolicyID != "" {
		// https://github.com/osbuild/blueprint-schema/issues/29
		e.log.Printf("policy ID %q is not supported, ignoring", e.from.Customizations.OpenSCAP.PolicyID)
	}

	if e.from.Customizations.OpenSCAP.JSONTailoring != nil {
		to.Tailoring = OpenSCAPTailoringFromJSON(TailoringJSON{
			JSONProfileID: e.from.Customizations.OpenSCAP.JSONTailoring.ProfileID,
			JSONFilePath:  e.from.Customizations.OpenSCAP.JSONTailoring.Filepath,
		})
	}

	if e.from.Customizations.OpenSCAP.Tailoring != nil {
		to.Tailoring = OpenSCAPTailoringFromProfiles(TailoringProfiles{
			Selected:   e.from.Customizations.OpenSCAP.Tailoring.Selected,
			Unselected: e.from.Customizations.OpenSCAP.Tailoring.Unselected,
		})
	}

	if reflect.DeepEqual(to, OpenSCAP{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importRegistration() *Registration {
	if e.from.Customizations == nil || e.from.Customizations.RHSM == nil || e.from.Customizations.RHSM.Config == nil {
		return nil
	}

	to := Registration{
		RegistrationRedHat: &RegistrationRedHat{
			RegistrationRHSM: &RegistrationRHSM{
				AutoEnable:           e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMConfig.AutoEnableYumPlugins,
				RepositoryManagement: e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMConfig.ManageRepos,
				AutoRegistration:     e.from.Customizations.RHSM.Config.SubscriptionManager.RHSMCertdConfig.AutoRegistration,
				Enabled:              e.from.Customizations.RHSM.Config.DNFPlugins.SubscriptionManager.Enabled,
				ProductPluginEnabled: e.from.Customizations.RHSM.Config.DNFPlugins.ProductID.Enabled,
			},
		},
	}

	if e.from.Customizations.FDO != nil {
		var insecure bool
		_, err := fmt.Sscanf(e.from.Customizations.FDO.DiunPubKeyInsecure, "%t", &insecure)
		if err != nil {
			e.log.Printf("cannot parse DiunPubKeyInsecure %q: %v, using default false", e.from.Customizations.FDO.DiunPubKeyInsecure, err)
		}

		to.RegistrationFDO = &RegistrationFDO{
			ManufacturingServerURL:  e.from.Customizations.FDO.ManufacturingServerURL,
			DiMfgStringTypeMacIface: e.from.Customizations.FDO.DiMfgStringTypeMacIface,
			DiunPubKeyHash:          e.from.Customizations.FDO.DiunPubKeyHash,
			DiunPubKeyInsecure:      insecure,
			DiunPubKeyRootCerts:     e.from.Customizations.FDO.DiunPubKeyRootCerts,
		}
	}

	if reflect.DeepEqual(to, Registration{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importStorage() *Storage {
	if e.from.Customizations == nil || e.from.Customizations.Disk == nil {
		return nil
	}

	to := Storage{
		Type: StorageType(e.from.Customizations.Disk.Type),
	}

	if e.from.Customizations.Disk.MinSize > 0 {
		to.Minsize = ToByteSize(e.from.Customizations.Disk.MinSize)
	}

	for _, part := range e.from.Customizations.Disk.Partitions {
		switch strings.ToLower(part.Type) {
		case "plain":
			fst, err := ParseFSType(part.FSType)
			if err != nil {
				e.log.Printf("error parsing filesystem type %q for partition %q: %v, using default", part.FSType, part.Name, err)
			}
			np := PartitionPlain{
				Type:       PartTypePlain,
				FSType:     fst,
				Label:      part.Label,
				Minsize:    ToByteSize(part.MinSize),
				Mountpoint: part.Mountpoint,
			}
			to.Partitions = append(to.Partitions, StoragePartitionFromPlain(np))
		case "btrfs":
			np := PartitionBTRFS{
				Type:    PartTypeBTRFS,
				Minsize: ToByteSize(part.MinSize),
			}
			for _, sv := range part.Subvolumes {
				nsv := PartitionSubvolumes{
					Name:       sv.Name,
					Mountpoint: sv.Mountpoint,
				}
				np.Subvolumes = append(np.Subvolumes, nsv)
			}
			to.Partitions = append(to.Partitions, StoragePartitionFromBTRFS(np))
		case "lvm":
			np := PartitionLVM{
				Type:    PartTypeLVM,
				Name:    part.Name,
				Minsize: ToByteSize(part.MinSize),
			}
			for _, lv := range part.LogicalVolumes {
				fst, err := ParseFSType(part.FSType)
				if err != nil {
					e.log.Printf("error parsing filesystem type %q for lv %q: %v, using default", part.FSType, lv.Name, err)
				}
				nlv := PartitionLV{
					Name:       lv.Name,
					Label:      lv.Label,
					FSType:     fst,
					Minsize:    ToByteSize(lv.MinSize),
					Mountpoint: lv.Mountpoint,
				}
				np.LogicalVolumes = append(np.LogicalVolumes, nlv)
			}
			to.Partitions = append(to.Partitions, StoragePartitionFromLVM(np))
		}
	}
	if reflect.DeepEqual(to, Storage{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importSystemd() *Systemd {
	if e.from.Customizations == nil || e.from.Customizations.Services == nil {
		return nil
	}

	to := Systemd{
		Enabled:  e.from.Customizations.Services.Enabled,
		Disabled: e.from.Customizations.Services.Disabled,
		Masked:   e.from.Customizations.Services.Masked,
	}

	if reflect.DeepEqual(to, Systemd{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importTimedate() *TimeDate {
	if e.from.Customizations == nil || e.from.Customizations.Timezone == nil {
		return nil
	}

	to := TimeDate{
		Timezone:   ptr.ValueOrEmpty(e.from.Customizations.Timezone.Timezone),
		NTPServers: e.from.Customizations.Timezone.NTPServers,
	}

	if reflect.DeepEqual(to, TimeDate{}) {
		return nil // omitzero
	}

	return &to
}
