package commands

import (
	"log"
	"os"
)

type Command interface {
	Run()
}

type commandType string

const (
	DeployType commandType = "deploy"
	AuthType               = "auth"
	HelpType               = "help"
)

func init() {
	if _, isProvided := os.LookupEnv("K8S_CLOUD_PROVIDER"); !isProvided {
		log.Fatalln("Env Var K8S_CLOUD_PROVIDER not provided")
	}
}

func Factory(executor commandType) Command {
	log.Printf("Factory from package command, Executor: %s\n", executor)
	switch executor {
	case DeployType:
		return newDeployCommand()
	case AuthType:
		return newAuthCommand()
	case HelpType:
		return newHelpCommand()
	default:
		return newDeployCommand()
	}
}
