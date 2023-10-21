package internal

import (
	"golang.org/x/crypto/ssh"
)

var factList = []string{"/bin/bash -ic '/usr/bin/env'"}

func GetFacts(client *ssh.Client) *[]ExecOutput {

	cmdOutputs := []ExecOutput{}

	for _, cmd := range factList {
		// Create a session
		co := execute(client, cmd)
		cmdOutputs = append(cmdOutputs, co)

	}
	return &cmdOutputs
}
