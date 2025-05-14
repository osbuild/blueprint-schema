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
	log.Println("version skipped")
	log.Println("distro skipped")

	to.Packages = ExportPackages(b)
	to.Groups = ExportGroups(b)
	to.Modules = ExportModules(b)
	to.Containers = ExportContainers(b)
	to.Customizations = ExportCustomizations(b)

	return to
}

func ExportPackages(b *from.Blueprint) []to.Package {
	var s []to.Package

	for _, pkg := range b.DNF.Packages {
		p := strings.SplitN(pkg, "-", 2)
		name := p[0]
		version := ""
		if len(p) > 1 {
			version = p[1]
		}

		s = append(s, to.Package{
			Name:    name,
			Version: version,
		})
	}

	return s
}

func ExportGroups(b *from.Blueprint) []to.Group {
	var s []to.Group

	for _, pkg := range b.DNF.Packages {
		s = append(s, to.Group{
			Name: pkg,
		})
	}

	return s
}

func ExportModules(b *from.Blueprint) []to.Package {
	var s []to.Package

	for _, pkg := range b.DNF.Packages {
		s = append(s, to.Package{
			Name: pkg,
		})
	}

	return s
}

func ExportContainers(b *from.Blueprint) []to.Container {
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

func ExportCustomizations(from *from.Blueprint) *to.Customizations {
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
		s = append(s, uc)
	}

	return s
}
