package commands

import (
	"log"
	"os"
	"os/exec"
)

type Deploy struct {
	manifestsDir string
}

func (d *Deploy) Run() {

	// GET DEPLOY NAME
	// SHOW HISTORY WITH KUBECTL ROLLOUT HISTORY DEPLOY/NAME
	// KUBECTL APPLY -F KUBERNETES/BRANCH
	// SHOW STATUS KUBECTL ROLLOUT STATUS DEPLOY/NAME
	// DONE
	log.Println("Deploy type runner")

	cmd := exec.Command("kubectl", "rollout", "history")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	log.Println(string(out))

	/*	s1 := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s1)

		log.Println(r.Intn(61))
		log.Println(r.Intn(61))
		log.Println(r.Intn(61))
		log.Println(r.Intn(61))
		log.Println(r.Intn(61))
		log.Println(r.Intn(61)) */
}

func newDeployCommand() Command {
	branch := os.Getenv("COMMIT_BRANCH")

	return &Deploy{
		manifestsDir: "kubernetes/" + branch,
	}
}
