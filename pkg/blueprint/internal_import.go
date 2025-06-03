package blueprint

import (
	"fmt"

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

	to.Name = e.from.Name
	to.Description = e.from.Description
	to.DNF.Packages = e.importPackages()

	e.to = to
	return e.log.Errors()
}

func (e *InternalImporter) Result() *Blueprint {
	return e.to
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
