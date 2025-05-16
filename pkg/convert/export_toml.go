package convert

import (
	"strings"

	from "github.com/osbuild/blueprint-schema/pkg/blueprint"
	"github.com/osbuild/blueprint-schema/pkg/ptr"
	to "github.com/osbuild/blueprint/pkg/blueprint"
)

func ExportBlueprint(b *from.Blueprint) *to.Blueprint {
	to := &to.Blueprint{}
	to.Name = b.Name
	to.Description = b.Description
	log.Println("TODO: skip the version or create a time-based one")

	to.Packages = exportPackages(b)
	to.EnabledModules = exportModules(b)
	to.Groups = exportGroups(b)
	to.Containers = exportContainers(b)
	to.Customizations = exportCustomizations(b)

	return to
}

func exportPackages(b *from.Blueprint) []to.Package {
	var s []to.Package

	for _, pkg := range b.DNF.Packages {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, to.Package{
			Name:    p[0],
			Version: p[1],
		})
	}

	return s
}

func exportGroups(b *from.Blueprint) []to.Group {
	var s []to.Group

	for _, pkg := range b.DNF.Groups {
		s = append(s, to.Group{
			Name: pkg,
		})
	}

	return s
}

func exportModules(b *from.Blueprint) []to.EnabledModule {
	var s []to.EnabledModule

	for _, pkg := range b.DNF.Modules {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, to.EnabledModule{
			Name:   p[0],
			Stream: p[1],
		})
	}

	return s
}

func exportContainers(b *from.Blueprint) []to.Container {
	var s []to.Container

	for _, container := range b.Containers {
		s = append(s, to.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    &container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return s
}

func exportCustomizations(from *from.Blueprint) *to.Customizations {
	if from == nil {
		return nil
	}

	to := &to.Customizations{}
	to.Hostname = &from.Hostname

	to.Kernel = ExportKernelCustomization(from.Kernel)
	to.User = ExportUserCustomization(from.Accounts.Users)

	return to
}

func ExportKernelCustomization(from *from.Kernel) *to.KernelCustomization {
	if from == nil {
		return nil
	}

	to := &to.KernelCustomization{}
	to.Name = from.Package
	to.Append = strings.Join(from.CmdlineAppend, " ")
	return to
}

func ExportUserCustomization(in []from.AccountsUsers) []to.UserCustomization {
	if in == nil {
		return nil
	}

	var s []to.UserCustomization

	log.Println("user force password reset ignored")
	for _, u := range in {
		uc := to.UserCustomization{}
		uc.Name = u.Name
		uc.Description = &u.Description
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = &u.SSHKeys[0]
		} else if len(u.SSHKeys) > 1 {
			uc.Key = &u.SSHKeys[0]
			log.Println("only one ssh key supported for user: %s", u.Name)
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
			uc.ExpireDate, err = ptr.ToErr(from.ExpireDateToEpochDays(*u.Expires))
			if err != nil {
				log.Printf("error converting expire date for user %s: %v", u.Name, err)
			}
		}
		if u.ForcePasswordChange != nil {
			uc.ForcePasswordReset = u.ForcePasswordChange
		}

		s = append(s, uc)
	}

	return s
}
