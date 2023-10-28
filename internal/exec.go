package internal

import (
	"bytes"
	"log"
	"math/rand"
	"time"
)

type ExecOutput struct {
	Cmd string
	Out string
	Err string
}

func (sc *sshConn) execute(cmd string) ExecOutput {
	ll := make([]byte, 0)
	mm := make([]byte, 0)
	time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
	sshOut := bytes.NewBuffer(ll)
	sshErr := bytes.NewBuffer(mm)

	session, err := sc.client.NewSession()
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
