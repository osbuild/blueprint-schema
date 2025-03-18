package onprem

import (
	"errors"

	int "github.com/osbuild/blueprint-schema/pkg/blueprint"
	ext "github.com/osbuild/blueprint-schema/pkg/conv/onprem/blueprint"
)

func ImportBlueprint(to *int.Blueprint, from *ext.Blueprint) error {
	var errs []error

	return errors.Join(errs...)
}
