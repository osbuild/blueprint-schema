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
	var err error

	to.Accounts = e.importAccounts()
	to.Architecture, err = ParseArch(e.from.Arch)
	if err != nil {
		e.log.Printf("error parsing architecture %q: %v", e.from.Arch, err)
	}
	to.CACerts = e.importCACerts()
	to.Containers = e.importContainers()
	to.DNF = e.importDNF()
	to.Description = e.from.Description
	to.Distribution = e.from.Distro
	to.FIPS = e.importFIPS()
	to.FSNodes = e.importFSNodes()

	to.Name = e.from.Name
	to.Kernel = e.importKernel()

	e.to = to
	return e.log.Errors()
}

func (e *InternalImporter) Result() *Blueprint {
	return e.to
}

func (e *InternalImporter) importDNF() *DNF {
	to := DNF{}
	to.Packages = e.importPackages()
	to.Modules = e.importModules()
	to.Groups = e.importGroups()
	to.ImportKeys = e.from.Customizations.RPM.ImportKeys.Files
	to.Repositories = e.importRepositories()

	if reflect.DeepEqual(to, DNF{}) {
		return nil // omitzero
	}

	return &to
}

func (e *InternalImporter) importPackages() []string {
	if e.from.Packages == nil {
		return nil
	}

	s := make([]string, len(e.from.Packages))
	for i, pkg := range e.from.Packages {
		s[i] = fmt.Sprintf("%s-%s", pkg.Name, pkg.Version)
	}

	return s
}

func (e *InternalImporter) importModules() []string {
	if e.from.Modules == nil {
		return nil
	}

	s := make([]string, len(e.from.EnabledModules))
	for i, pkg := range e.from.Modules {
		s[i] = fmt.Sprintf("%s-%s", pkg.Name, pkg.Version)
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
			ModuleHotfixes: ptr.FromOr(repo.ModuleHotfixes, false),
			Priority:       ptr.FromOr(repo.Priority, 0),
			SSLVerify:      repo.SSLVerify,
			Usage: &DnfRepositoryUsage{
				Configure: ptr.To(true),
			},
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
			Description:         ptr.From(user.Description),
			Home:                ptr.From(user.Home),
			UID:                 ptr.FromOr(user.UID, 0),
			GID:                 ptr.FromOr(user.GID, 0),
			Groups:              user.Groups,
			Password:            user.Password,
			Expires:             ParseExpireDate(user.ExpireDate),
			ForcePasswordChange: user.ForcePasswordReset,
			Shell:               ptr.From(user.Shell),
		}

		if user.Key != nil {
			u.SSHKeys = []string{*user.Key}
		}

		to.Users = append(to.Users, u)
	}

	for _, group := range e.from.Customizations.Group {
		g := AccountsGroups{
			Name: group.Name,
			GID:  ptr.FromOr(group.GID, 0),
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
		Enabled: ptr.FromOr(e.from.Customizations.FIPS, false),
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
		mode, err := parseOctalString(file.Mode)
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
		mode, err := parseOctalString(dir.Mode)
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
