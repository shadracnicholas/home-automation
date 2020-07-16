package deployer

import (
	"fmt"

	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/config"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/deployer/systemd"
)

// Deployer deploys services
type Deployer interface {
	Revision() (string, error)
	Deploy(revision string) error
}

// Choose returns an appropriate deployer for the service and target
func Choose(service *config.Service, target *config.Target) (Deployer, error) {
	switch target.System {
	case config.SysSystemd:
		return &systemd.Systemd{
			Service: service,
			Target:  target,
		}, nil
	}

	return nil, fmt.Errorf("unsupported system %q", target.System)
}
