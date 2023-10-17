package internal

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

type Login struct {
	host   string
	user   string
	pw     string
	pkey   string
	client *ssh.Client
	// protocol Protocol
}

func NewLogin(host, user, pw, pkey string) *ssh.Client {
	lg := &Login{
		host: host,
		user: user,
		pw:   pw,
		pkey: pkey,
		// protocol: pr,
	}
	client := lg.Ping()
	lg.client = client
	return client
}

func (lg *Login) Ping() *ssh.Client {
	config := &ssh.ClientConfig{
		User: lg.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(lg.pw),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", lg.host, 22), config)
	if err != nil {

		d := color.New(color.FgRed)
		d.Println("Cannot Connect", err)
		os.Exit(1)
	}
	d := color.New(color.FgCyan, color.Bold)
	d.Println("Successful Connection")
	return client
}
