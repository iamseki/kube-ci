package commands

import (
	"log"
	"os"
	"os/exec"
)

// Config struct represents a configuration object
type Config struct {
}

// Run the config logic
func (c *Config) Run() {
	configFileURL := os.Getenv("CONFIG_FILE_URL")
	configProvider := os.Getenv("CONFIG_PROVIDER")

	if len(configFileURL) != 0 {
		log.Println(configFileURL)
		log.Println(configProvider)
		// ensuring that is not kubeconfig file already in root folder
		cmd := exec.Command("rm", "kubeconfig")
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		// then download it from some cloud provided storage
		// *** implementation here *** //
	} else {
		log.Println("Getting config file from /app/kubeconfig ..")
	}
}

func newConfigCommand() Command {
	return &Config{}
}
