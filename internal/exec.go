package internal

import (
	"bytes"
	"log"

	"golang.org/x/crypto/ssh"
)

type ExecOutput struct {
	Cmd string
	Out string
	Err string
}

func execute(client *ssh.Client, cmd string) ExecOutput {
	ll := make([]byte, 0)
	mm := make([]byte, 0)
	sshOut := bytes.NewBuffer(ll)
	sshErr := bytes.NewBuffer(mm)

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	session.Stdout = sshOut
	session.Stderr = sshErr

	session.Run(cmd)
	co := ExecOutput{
		Out: sshOut.String(),
		Err: sshErr.String(),
		Cmd: cmd,
	}
	return co
}
