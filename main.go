package main

import (
	kube "kube-ci/commands"
)

func main() {
	auth := kube.Factory(kube.AuthType)
	deploy := kube.Factory(kube.DeployType)

	auth.Run()
	deploy.Run()
}
