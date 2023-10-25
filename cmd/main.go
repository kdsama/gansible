package main

import "github.com/kdsama/gansible/internal"

// "golang.org/x/crypto/ssh"

func main() {

	engine := internal.NewEngine("./pb.yml", "")
	engine.Run()
}
