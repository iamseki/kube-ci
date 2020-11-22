package commands

import (
	"log"
)

// Command Interface is implemented by commands
type Command interface {
	Run()
}

type commandType string

//Commands value used to instantiate Factory
const (
	DeployType commandType = "deploy"
	AuthType               = "auth"
	HelpType               = "help"
	ConfigType             = "config"
)

func init() {
	/*if _, isProvided := os.LookupEnv("K8S_CLOUD_PROVIDER"); !isProvided {
		log.Fatalln("Env Var K8S_CLOUD_PROVIDER not provided")
	}*/
}

// Factory returns a command struct that run a specific command
func Factory(executor commandType) Command {
	log.Printf("Factory from kube-ci, Yielding command structure: %s\n", executor)
	switch executor {
	case DeployType:
		return newDeployCommand()
	case AuthType:
		return newAuthCommand()
	case ConfigType:
		return newConfigCommand()
	default:
		return newDeployCommand()
	}
}
