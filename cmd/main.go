package main

import (
	"flag"

	"github.com/kdsama/gansible/internal"
)

// "golang.org/x/crypto/ssh"

func main() {

	var (
		inventory = flag.String("i", "", "inventory file for executing the playbook ")
		playbook  = flag.String("p", "", "Playbook to run ")
	)
	flag.Parse()

	engine := internal.NewEngine(*playbook, *inventory)
	engine.Run()
}
