package commands

import (
	"kube-ci/providers"
	"os"
)

type Auth struct {
}

func (l *Auth) Run() {

	cloudProvider := os.Getenv("K8S_CLOUD_PROVIDER")
	var p providers.CloudCommand
	switch cloudProvider {
	case "AWS":
		p = providers.Factory(providers.AWSType)
	default:
		p = providers.Factory(providers.GCPType)
	}

	p.Authenticate()
}

func newAuthCommand() Command {
	return &Auth{}
}
