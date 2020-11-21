package commands

import (
	"log"
	"math/rand"
	"time"
)

type Deploy struct {
}

func (d *Deploy) Run() {
	log.Println("Deploy type runner")

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	log.Println(r.Intn(61))
	log.Println(r.Intn(61))
	log.Println(r.Intn(61))
	log.Println(r.Intn(61))
	log.Println(r.Intn(61))
	log.Println(r.Intn(61))
}

func newDeployCommand() Command {
	return &Deploy{}
}
