package automator

import (
	"errors"
	"strings"

	"github.com/weaveworks/flux/instance"
	"github.com/weaveworks/flux/jobs"
)

// Config collects the parameters to the automator. All fields are mandatory.
type Config struct {
	Releaser   jobs.JobReadPusher
	InstanceDB instance.DB
}

// Validate returns an error if the config is underspecified.
func (cfg Config) Validate() error {
	var errs []string
	if cfg.Releaser == nil {
		errs = append(errs, "releaser not supplied")
	}
	if cfg.InstanceDB == nil {
		errs = append(errs, "instance configuration DB not supplied")
	}
	if len(errs) > 0 {
		return errors.New("invalid: " + strings.Join(errs, "; "))
	}
	return nil
}
