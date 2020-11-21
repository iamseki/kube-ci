package providers

import "log"

type AWS struct {
}

func (a *AWS) Authenticate() {
	log.Println("AWS Provider Authentication")
}

func newAWSProvider() CloudCommand {
	return &AWS{}
}
