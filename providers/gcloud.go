package providers

import "log"

type GCP struct {
}

func (a *GCP) Authenticate() {
	log.Println("GCP Provider Authentication")
}

func newGCPProvider() CloudCommand {
	return &GCP{}
}
