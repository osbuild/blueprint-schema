package onprem

import (
	"strings"

	int "github.com/osbuild/blueprint-schema"
	"github.com/osbuild/blueprint-schema/conv/notes"
	ext "github.com/osbuild/blueprint-schema/conv/onprem/blueprint"
	ptr "github.com/osbuild/blueprint-schema/conv/ptr"
)

func ExportBlueprint(from *int.Blueprint, nts *notes.ConversionNotes) *ext.Blueprint {
	to := &ext.Blueprint{}
	to.Name = from.Name
	to.Description = from.Description
	nts.Warn("version skipped")
	nts.Warn("distro skipped")

	to.Packages = ExportPackages(from, nts)
	to.Groups = ExportGroups(from, nts)
	to.Modules = ExportModules(from, nts)
	to.Containers = ExportContainers(from, nts)
	to.Customizations = ExportCustomizations(from, nts)

	return to
}

func ExportPackages(from *int.Blueprint, nts *notes.ConversionNotes) []ext.Package {
	var to []ext.Package

	nts.Warn("packages added with version in name")
	for _, pkg := range from.DNF.Packages {
		to = append(to, ext.Package{
			Name: pkg,
		})
	}

	return to
}

func ExportGroups(from *int.Blueprint, nts *notes.ConversionNotes) []ext.Group {
	var to []ext.Group
	nts.Warn("groups added with version in name")

	for _, group := range from.DNF.Groups {
		to = append(to, ext.Group{
			Name: group,
		})
	}

	return to
}

func ExportModules(from *int.Blueprint, nts *notes.ConversionNotes) []ext.Package {
	var to []ext.Package
	nts.Warn("modules added with version in name")

	for _, module := range from.DNF.Modules {
		to = append(to, ext.Package{
			Name: module,
		})
	}

	return to
}

func ExportContainers(from *int.Blueprint, nts *notes.ConversionNotes) []ext.Container {
	var to []ext.Container

	for _, container := range from.Containers {
		to = append(to, ext.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return to
}

//	type Customizations struct {
//		Firewall           *FirewallCustomization         `json:"firewall,omitempty" toml:"firewall,omitempty"`
//		Services           *ServicesCustomization         `json:"services,omitempty" toml:"services,omitempty"`
//		Filesystem         []FilesystemCustomization      `json:"filesystem,omitempty" toml:"filesystem,omitempty"`
//		Disk               *DiskCustomization             `json:"disk,omitempty" toml:"disk,omitempty"`
//		InstallationDevice string                         `json:"installation_device,omitempty" toml:"installation_device,omitempty"`
//		FDO                *FDOCustomization              `json:"fdo,omitempty" toml:"fdo,omitempty"`
//		OpenSCAP           *OpenSCAPCustomization         `json:"openscap,omitempty" toml:"openscap,omitempty"`
//		Ignition           *IgnitionCustomization         `json:"ignition,omitempty" toml:"ignition,omitempty"`
//		Directories        []DirectoryCustomization       `json:"directories,omitempty" toml:"directories,omitempty"`
//		Files              []FileCustomization            `json:"files,omitempty" toml:"files,omitempty"`
//		Repositories       []RepositoryCustomization      `json:"repositories,omitempty" toml:"repositories,omitempty"`
//		FIPS               *bool                          `json:"fips,omitempty" toml:"fips,omitempty"`
//		ContainersStorage  *ContainerStorageCustomization `json:"containers-storage,omitempty" toml:"containers-storage,omitempty"`
//		Installer          *InstallerCustomization        `json:"installer,omitempty" toml:"installer,omitempty"`
//		RPM                *RPMCustomization              `json:"rpm,omitempty" toml:"rpm,omitempty"`
//		RHSM               *RHSMCustomization             `json:"rhsm,omitempty" toml:"rhsm,omitempty"`
//		CACerts            *CACustomization               `json:"cacerts,omitempty" toml:"cacerts,omitempty"`
//	}
func ExportCustomizations(from *int.Blueprint, nts *notes.ConversionNotes) *ext.Customizations {
	if from == nil {
		return nil
	}

	to := &ext.Customizations{}
	to.Hostname = &from.Hostname

	to.Kernel = ExportKernelCustomization(from.Kernel, nts)
	to.User = ExportUserCustomization(from.Accounts.Users, nts)
	to.Group = ExportGroupCustomization(from.Accounts.Groups, nts)
	to.Timezone = ExportTimezoneCustomization(from, nts)
	to.Locale = ExportLocaleCustomization(from.Locale, nts)
	to.Firewall = ExportFirewallCustomization(from.Network.Firewall, nts)

	return to
}

func ExportKernelCustomization(from *int.Kernel, nts *notes.ConversionNotes) *ext.KernelCustomization {
	if from == nil {
		return nil
	}

	to := &ext.KernelCustomization{}
	to.Name = from.Package
	to.Append = strings.Join(from.CmdlineAppend, " ")
	return to
}

func ExportUserCustomization(from []int.UserAccount, nts *notes.ConversionNotes) []ext.UserCustomization {
	if from == nil {
		return nil
	}

	var to []ext.UserCustomization

	nts.Warn("user force password reset ignored")
	for _, fUser := range from {
		toUser := ext.UserCustomization{}
		toUser.Name = fUser.Name
		toUser.Description = &fUser.Description
		toUser.Password = &fUser.Password
		if len(fUser.SshKeys) == 1 {
			toUser.Key = &fUser.SshKeys[0]
		} else if len(fUser.SshKeys) > 1 {
			toUser.Key = &fUser.SshKeys[0]
			nts.Warn("only one ssh key supported for user: %s", fUser.Name)
		}
		toUser.Home = &fUser.Home
		toUser.Shell = &fUser.Shell
		toUser.Groups = fUser.Groups
		if fUser.UID != 0 {
			toUser.UID = &fUser.UID
		}
		if fUser.GID != 0 {
			toUser.GID = &fUser.GID
		}
		if !fUser.Expires.IsZero() {
			toUser.ExpireDate = ptr.To(fUser.Expires.DaysFrom1970())
		}
		to = append(to, toUser)
	}

	return to
}

func ExportGroupCustomization(from []int.GroupAccount, nts *notes.ConversionNotes) []ext.GroupCustomization {
	if from == nil {
		return nil
	}

	var to []ext.GroupCustomization
	for _, fGroup := range from {
		toGroup := ext.GroupCustomization{}
		toGroup.Name = fGroup.Name
		toGroup.GID = &fGroup.GID
		to = append(to, toGroup)
	}

	return to
}

func ExportTimezoneCustomization(from *int.Blueprint, nts *notes.ConversionNotes) *ext.TimezoneCustomization {
	if from.TimeDate == nil {
		return nil
	}

	to := &ext.TimezoneCustomization{}
	to.Timezone = &from.TimeDate.Timezone
	to.NTPServers = from.TimeDate.NTPServers

	return to
}

func ExportLocaleCustomization(from *int.Locale, nts *notes.ConversionNotes) *ext.LocaleCustomization {
	if from == nil {
		return nil
	}

	to := &ext.LocaleCustomization{}
	to.Languages = from.Languages
	if len(from.Keyboards) == 1 {
		to.Keyboard = &from.Keyboards[0]
	} else if len(from.Keyboards) > 1 {
		to.Keyboard = &from.Keyboards[0]
		nts.Warn("only one keyboard layout converted: %s", *to.Keyboard)
	}

	return to
}

func ExportFirewallCustomization(from *int.NetworkFirewall, nts *notes.ConversionNotes) *ext.FirewallCustomization {
	if from == nil {
		return nil
	}

	to := &ext.FirewallCustomization{}
	to.Services = ExportServicesCustomization(from, nts)
	to.Ports = ExportPortsCustomization(from, nts)
	
	return to
}

func ExportServicesCustomization(from *int.NetworkFirewall, nts *notes.ConversionNotes) *ext.FirewallServicesCustomization {
	if from == nil {
		return nil
	}

	enabled, err := from.ServicesAsFirewalld(true)
	if err != nil {
		nts.Fatal("error converting services to firewalld: %s", err)
		return nil
	}

	disabled, err := from.ServicesAsFirewalld(false)
	if err != nil {
		nts.Fatal("error converting services to firewalld: %s", err)
		return nil
	}

	to := &ext.FirewallServicesCustomization{
		Enabled:  make([]string, len(enabled)),
		Disabled: make([]string, len(disabled)),
	}
	copy(to.Enabled, enabled)
	copy(to.Disabled, disabled)
	return to
}

func ExportPortsCustomization(from *int.NetworkFirewall, nts *notes.ConversionNotes) []string {
	if from == nil {
		return nil
	}

	enabled, err := from.PortsAsFirewalld(true)
	if err != nil {
		nts.Fatal("error converting ports to firewalld: %s", err)
		return nil
	}

	disabled, err := from.PortsAsFirewalld(false)
	if err != nil {
		nts.Fatal("error converting ports to firewalld: %s", err)
		return nil
	}
	if len(disabled) > 0 {
		nts.Warn("skipping disabled firewall ports: %q", disabled)
	}

	return enabled
}
