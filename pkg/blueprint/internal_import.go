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

	to.Name = e.from.Name
	to.Description = e.from.Description
	to.DNF = e.importDNF()
	to.Containers = e.importContainers()
	to.Kernel = e.exportKernel()
	to.Distribution = e.from.Distro
	to.Architecture, err = ParseArch(e.from.Arch)
	if err != nil {
		e.log.Printf("error parsing architecture %q: %v", e.from.Arch, err)
	}

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

func (e *InternalImporter) exportKernel() *Kernel {
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
