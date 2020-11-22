package main

import (
	"flag"
	kube "kube-ci/commands"
)

func main() {
	var config bool

	flag.BoolVar(&config, "config", false, "Config flag must be provided if wanted to run config command")
	flag.BoolVar(&config, "c", false, "Config flag must be provided if wanted to run config command")

	flag.Parse()

	if config {
		config := kube.Factory(kube.ConfigType)
		config.Run()
	} else {
		auth := kube.Factory(kube.AuthType)
		deploy := kube.Factory(kube.DeployType)

		auth.Run()
		deploy.Run()
	}
}
