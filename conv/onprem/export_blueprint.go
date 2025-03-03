package onprem

import (
	"strings"

	int "github.com/osbuild/blueprint-schema"
	ext "github.com/osbuild/blueprint-schema/conv/onprem/blueprint"
	ptr "github.com/osbuild/blueprint-schema/conv/ptr"
)

func ExportBlueprint(to *ext.Blueprint, from *int.Blueprint, errs *Errors) {
	to.Name = from.Name
	to.Description = from.Description
	errs.Warn("version skipped")
	errs.Warn("distro skipped")

	ExportPackages(to.Packages, from, errs)
	ExportGroups(to.Groups, from, errs)
	ExportModules(to.Modules, from, errs)
	ExportContainers(to.Containers, from, errs)
	to.Customizations = &ext.Customizations{}
	ExportCustomizations(to.Customizations, from, errs)
}

func ExportPackages(to []ext.Package, from *int.Blueprint, errs *Errors) {
	errs.Warn("packages added with version in name")
	for _, pkg := range from.DNF.Packages {
		to = append(to, ext.Package{
			Name: pkg,
		})
	}
}

func ExportGroups(to []ext.Group, from *int.Blueprint, errs *Errors) {
	errs.Warn("groups added with version in name")
	for _, group := range from.DNF.Groups {
		to = append(to, ext.Group{
			Name: group,
		})
	}
}

func ExportModules(to []ext.Package, from *int.Blueprint, errs *Errors) {
	errs.Warn("modules added with version in name")
	for _, module := range from.DNF.Modules {
		to = append(to, ext.Package{
			Name: module,
		})
	}
}

func ExportContainers(to []ext.Container, from *int.Blueprint, errs *Errors) {
	for _, container := range from.Containers {
		to = append(to, ext.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}
}

//	type Customizations struct {
//		User               []UserCustomization            `json:"user,omitempty" toml:"user,omitempty"`
//		Group              []GroupCustomization           `json:"group,omitempty" toml:"group,omitempty"`
//		Timezone           *TimezoneCustomization         `json:"timezone,omitempty" toml:"timezone,omitempty"`
//		Locale             *LocaleCustomization           `json:"locale,omitempty" toml:"locale,omitempty"`
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
func ExportCustomizations(to *ext.Customizations, from *int.Blueprint, errs *Errors) {
	to.Hostname = &from.Hostname

	to.Kernel = &ext.KernelCustomization{}
	ExportKernelCustomization(to.Kernel, from.Kernel, errs)

	to.User = []ext.UserCustomization{}
	ExportUserCustomization(to.User, from.Accounts.Users, errs)
}

func ExportKernelCustomization(to *ext.KernelCustomization, from *int.Kernel, errs *Errors) {
	to.Name = from.Package
	to.Append = strings.Join(from.CmdlineAppend, " ")
}

func ExportUserCustomization(to []ext.UserCustomization, from []int.UserAccount, errs *Errors) {
	if from == nil {
		return
	}

	errs.Warn("user force password reset ignored")
	for _, fUser := range from {
		toUser := ext.UserCustomization{}
		toUser.Name = fUser.Name
		toUser.Description = &fUser.Description
		toUser.Password = &fUser.Password
		if len(fUser.SshKeys) == 1 {
			toUser.Key = &fUser.SshKeys[0]
		} else if len(fUser.SshKeys) > 1 {
			toUser.Key = &fUser.SshKeys[0]
			errs.Fatal("only one ssh key supported for user", fUser.Name)
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
}
