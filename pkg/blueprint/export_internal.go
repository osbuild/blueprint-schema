package blueprint

import (
	"fmt"
	"strings"
	"time"

	"github.com/osbuild/blueprint-schema/pkg/ptr"
	int "github.com/osbuild/blueprint/pkg/blueprint"
)

type InternalExporter struct {
	b   *Blueprint
	to  *int.Blueprint
	log *logs
}

func NewInternalExporter(inputBlueprint *Blueprint) *InternalExporter {
	return &InternalExporter{
		b:   inputBlueprint,
		log: newCollector(),
	}
}

// ExportInternal converts the blueprint to the internal representation.
func (e *InternalExporter) Export(ed ComposeRequest) error {
	to := &int.Blueprint{}

	to.Name = e.b.Name
	to.Description = e.b.Description
	if ed.Version == "" {
		// Create monotonic incremental version number based on miliseconds
		to.Version = int64ToVersion(uint64(time.Now().UTC().UnixMilli()))
	} else {
		to.Version = ed.Version
	}

	to.Packages = e.exportPackages()
	to.EnabledModules = e.exportModules()
	to.Groups = e.exportGroups()
	to.Containers = e.exportContainers()
	to.Customizations = e.exportCustomizations()
	to.Distro = ed.Distro
	to.Arch = ed.Arch

	e.to = to
	return e.log.Errors()
}

func (e *InternalExporter) Result() *int.Blueprint {
	return e.to
}

func (e *InternalExporter) exportPackages() []int.Package {
	var s []int.Package

	for _, pkg := range e.b.DNF.Packages {
		p := splitStringEmptyN(pkg, "-", 2)

		s = append(s, int.Package{
			Name:    p[0],
			Version: p[1],
		})
	}

	return s
}

func (e *InternalExporter) exportGroups() []int.Group {
	var s []int.Group

	for _, pkg := range e.b.DNF.Groups {
		s = append(s, int.Group{
			Name: pkg,
		})
	}

	return s
}

func (e *InternalExporter) exportModules() []int.EnabledModule {
	var s []int.EnabledModule

	for _, pkg := range e.b.DNF.Modules {
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

	for _, container := range e.b.Containers {
		s = append(s, int.Container{
			Name:         container.Name,
			Source:       container.Source,
			TLSVerify:    &container.TLSVerify,
			LocalStorage: container.LocalStorage,
		})
	}

	return s
}

func (e *InternalExporter) exportCustomizations() *int.Customizations {
	to := &int.Customizations{}

	to.Hostname = &e.b.Hostname
	to.Kernel = e.exportKernel()
	to.User = e.exportUserCustomization()
	to.Group = e.exportGroupCustomization()
	to.Timezone = e.exportTimezoneCustomization()
	to.Locale = e.exportLocaleCustomization()
	to.Firewall = e.exportFirewallCustomization()
	to.Services = e.exportSystemdCustomization()
	to.Disk = e.exportStorage()

	return to
}

func (e *InternalExporter) exportKernel() *int.KernelCustomization {
	if e.b.Kernel == nil {
		return nil
	}

	to := &int.KernelCustomization{}
	to.Name = e.b.Kernel.Package
	to.Append = strings.Join(e.b.Kernel.CmdlineAppend, " ")

	return to
}

func (e *InternalExporter) exportUserCustomization() []int.UserCustomization {
	if e.b.Accounts.Users == nil {
		return nil
	}

	var s []int.UserCustomization
	for _, u := range e.b.Accounts.Users {
		uc := int.UserCustomization{}
		uc.Name = u.Name
		uc.Description = &u.Description
		uc.Password = u.Password
		if len(u.SSHKeys) == 1 {
			uc.Key = &u.SSHKeys[0]
		} else if len(u.SSHKeys) > 1 {
			uc.Key = &u.SSHKeys[0]
			e.log.Printf("only one ssh key supported for user: %s", u.Name)
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
	if e.b.Accounts.Groups == nil {
		return nil
	}

	var s []int.GroupCustomization
	for _, g := range e.b.Accounts.Groups {
		gc := int.GroupCustomization{}
		gc.Name = g.Name
		if g.GID != 0 {
			gc.GID = &g.GID
		}
		s = append(s, gc)
	}

	return s
}

func (e *InternalExporter) exportTimezoneCustomization() *int.TimezoneCustomization {
	if e.b.Timedate == nil {
		return nil
	}

	to := &int.TimezoneCustomization{}
	to.Timezone = &e.b.Timedate.Timezone
	to.NTPServers = e.b.Timedate.NTPServers

	return to
}

func (e *InternalExporter) exportLocaleCustomization() *int.LocaleCustomization {
	if e.b.Locale == nil {
		return nil
	}

	to := &int.LocaleCustomization{}
	if len(e.b.Locale.Keyboards) > 0 {
		to.Keyboard = &e.b.Locale.Keyboards[0]
		if len(e.b.Locale.Keyboards) > 1 {
			e.log.Println("only one keyboard layout supported, selecting first one")
		}
	}
	to.Languages = e.b.Locale.Languages

	return to
}

func (e *InternalExporter) exportFirewallCustomization() *int.FirewallCustomization {
	if e.b.Network.Firewall == nil || len(e.b.Network.Firewall.Services) == 0 {
		return nil
	}

	to := &int.FirewallCustomization{
		Ports: make([]string, 0),
		Services: &int.FirewallServicesCustomization{
			Enabled:  make([]string, 0),
			Disabled: make([]string, 0),
		},
	}
	for i, s := range e.b.Network.Firewall.Services {
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

	if len(to.Ports) == 0 {
		to.Ports = nil
	}
	if len(to.Services.Enabled) == 0 {
		to.Services.Enabled = nil
	}
	if len(to.Services.Disabled) == 0 {
		to.Services.Disabled = nil
	}
	if to.Services.Enabled == nil && to.Services.Disabled == nil {
		to.Services = nil
	}

	return to
}

func (e *InternalExporter) exportSystemdCustomization() *int.ServicesCustomization {
	if e.b.Systemd == nil {
		return nil
	}

	to := &int.ServicesCustomization{}

	if len(e.b.Systemd.Enabled) > 0 {
		to.Enabled = make([]string, len(e.b.Systemd.Enabled))
		copy(to.Enabled, e.b.Systemd.Enabled)
	}
	if len(e.b.Systemd.Disabled) > 0 {
		to.Disabled = make([]string, len(e.b.Systemd.Disabled))
		copy(to.Disabled, e.b.Systemd.Disabled)
	}
	if len(e.b.Systemd.Masked) > 0 {
		to.Masked = make([]string, len(e.b.Systemd.Masked))
		copy(to.Masked, e.b.Systemd.Masked)
	}

	return to
}

func (e *InternalExporter) exportStorage() *int.DiskCustomization {
	if e.b.Storage == nil {
		return nil
	}

	to := &int.DiskCustomization{}
	to.Type = e.b.Storage.Type.String()
	size, err := ParseSize(e.b.Storage.Minsize)
	if err != nil {
		e.log.Printf("error parsing size %s: %v", e.b.Storage.Minsize, err)
	}
	to.MinSize = size.Bytes()

	if len(e.b.Storage.Partitions) == 0 {
		return to
	}

	// TODO: the rest

	return to
}
