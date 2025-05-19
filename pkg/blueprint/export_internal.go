package blueprint

import (
	"fmt"
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	int "github.com/osbuild/blueprint/pkg/blueprint"
)

// ExportData is used for feeding the export function with the
// information needed to export the blueprint.
type ExportData struct {
	Version string
	Distro  string
	Arch    string
}

// ExportInternal converts the blueprint to the internal representation.
func (b *Blueprint) ExportInternal(ed ExportData) *int.Blueprint {
	to := &int.Blueprint{}
	to.Name = b.Name
	to.Description = b.Description
	if ed.Version == "" {
		// Create monotonic incremental version number based on miliseconds
		to.Version = int64ToVersion(uint64(time.Now().UTC().UnixMilli()))
	} else {
		to.Version = ed.Version
	}

	to.Packages = exportPackages(b)
	to.EnabledModules = exportModules(b)
	to.Groups = exportGroups(b)
	to.Containers = exportContainers(b)
	to.Customizations = exportCustomizations(b)
	to.Distro = ed.Distro
	to.Arch = ed.Arch
	return to
}

func exportPackages(b *Blueprint) []int.Package {
	var s []int.Package

	for _, pkg := range b.DNF.Packages {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, int.Package{
			Name:    p[0],
			Version: p[1],
		})
	}
	return s
}

func exportGroups(b *Blueprint) []int.Group {
	var s []int.Group

	for _, pkg := range b.DNF.Groups {
		s = append(s, int.Group{
			Name: pkg,
		})
	}
	return s
}

func exportModules(b *Blueprint) []int.EnabledModule {
	var s []int.EnabledModule

	for _, pkg := range b.DNF.Modules {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, int.EnabledModule{
			Name:   p[0],
			Stream: p[1],
		})
	}
	return s
}

func exportContainers(b *Blueprint) []int.Container {
	var s []int.Container

	for _, container := range b.Containers {
		s = append(s, int.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    &container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}
	return s
}

func exportCustomizations(from *Blueprint) *int.Customizations {
	if from == nil {
		return nil
	}

	to := &int.Customizations{}
	to.Hostname = &from.Hostname
	// SSHKey is not supported, use user customization instead
	to.Kernel = ExportKernelCustomization(from.Kernel)
	if from.Accounts != nil {
		to.User = ExportUserCustomization(from.Accounts.Users)
		to.Group = ExportGroupCustomization(from.Accounts.Groups)
	}
	to.Timezone = ExportTimezoneCustomization(from.Timedate)
	to.Locale = ExportLocaleCustomization(from.Locale)
	if from.Network != nil {
		to.Firewall = ExportFirewallCustomization(from.Network.Firewall)
	}
	return to
}

func ExportKernelCustomization(from *Kernel) *int.KernelCustomization {
	if from == nil {
		return nil
	}

	to := &int.KernelCustomization{}
	to.Name = from.Package
	to.Append = strings.Join(from.CmdlineAppend, " ")
	return to
}

func ExportUserCustomization(in []AccountsUsers) []int.UserCustomization {
	if in == nil {
		return nil
	}

	var s []int.UserCustomization

	for _, u := range in {
		uc := int.UserCustomization{}
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
			uc.ExpireDate, err = ptr.ToErr(ExpireDateToEpochDays(*u.Expires))
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

func ExportGroupCustomization(in []AccountsGroups) []int.GroupCustomization {
	if in == nil {
		return nil
	}

	var s []int.GroupCustomization

	for _, g := range in {
		gc := int.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = &g.GID
		}
		s = append(s, gc)
	}
	return s
}

func ExportTimezoneCustomization(from *TimeDate) *int.TimezoneCustomization {
	if from == nil {
		return nil
	}

	to := &int.TimezoneCustomization{}
	to.Timezone = &from.Timezone
	to.NTPServers = from.NTPServers
	return to
}

func ExportLocaleCustomization(from *Locale) *int.LocaleCustomization {
	if from == nil {
		return nil
	}

	to := &int.LocaleCustomization{}
	if len(from.Keyboards) > 0 {
		to.Keyboard = &from.Keyboards[0]
		if len(from.Keyboards) > 1 {
			log.Println("only one keyboard layout supported, selecting first one")
		}
	}
	to.Languages = from.Languages
	return to
}

func ExportFirewallCustomization(from *NetworkFirewall) *int.FirewallCustomization {
	if from == nil {
		return nil
	}

	to := &int.FirewallCustomization{
		Ports: make([]string, 0),
		Services: &int.FirewallServicesCustomization{
			Enabled:  make([]string, 0),
			Disabled: make([]string, 0),
		},
	}
	for i, s := range from.Services {
		fs, fp, fft, err := s.SelectUnion()
		if err != nil {
			log.Printf("could not parse network service %i: %v", i, err)
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
				log.Printf("network service %i error: port number %d cannot be disabled", i, fp.Port)
			}
		} else if fft.From != 0 && fft.To != 0 {
			if fft.Protocol != "" {
				proto = fft.Protocol.String()
			}
			srv := fmt.Sprintf("%d-%d/%s", fft.From, fft.To, proto)

			if fft.Enabled == nil || *fft.Enabled {
				to.Ports = append(to.Ports, srv)
			} else {
				log.Printf("network service %i error: port number %d cannot be disabled", i, fp.Port)
			}
		} else {
			log.Printf("network service %i error: one of service, port or from and to present", i)
			continue
		}
	}
	return to
}
