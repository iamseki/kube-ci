package providers

type CloudCommand interface {
	Authenticate()
}

type CloudProvider string

const (
	AWSType CloudProvider = "aws"
	GCPType               = "gcp"
	DGOType               = "dgo"
	HERType               = "her"
)

func Factory(provider CloudProvider) CloudCommand {
	switch provider {
	case AWSType:
		return newAWSProvider()
	case GCPType:
		return newGCPProvider()
	default:
		return newAWSProvider()
	}
}
