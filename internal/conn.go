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
	port   int
	// protocol Protocol
}

func NewLogin(host, user, pw, pkey string, port int) *ssh.Client {
	if port == 0 {
		port = 22
	}
	fmt.Println("Port is", port)
	lg := &Login{
		host: host,
		user: user,
		pw:   pw,
		pkey: pkey,
		port: port,
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
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", lg.host, lg.port), config)
	if err != nil {

		d := color.New(color.FgRed)
		d.Println("Cannot Connect", err)
		os.Exit(1)
	}
	d := color.New(color.FgCyan, color.Bold)
	d.Println("Successful Connection")
	return client
}
