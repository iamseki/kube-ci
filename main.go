package main

import (
	kube "kube-ci/commands"
	"log"
)

func main() {
	log.Println("Starting Project")
	kube.Factory()
}
