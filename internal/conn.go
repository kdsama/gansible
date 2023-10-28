package internal

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

type ssher interface {
	add(name, host, user, pw, pkey string, port int)
	get(name string) (*sshConn, bool)
	execute(name string, cmd string) ExecOutput
}

type sshService map[string]*sshConn

func NewSSHService() *sshService {

	return new(sshService)
}

func (ss *sshService) add(name, host, user, pw, pkey string, port int) {
	conn := NewSshConn(host, user, pw, pkey, port)
	(*ss)[name] = conn
}
func (ss *sshService) get(name string) (*sshConn, bool) {
	val, ok := (*ss)[name]
	return val, ok
}
func (ss *sshService) execute(name string, cmd string) ExecOutput {
	return (*ss)[name].execute(cmd)
}

type sshConn struct {
	host   string
	user   string
	pw     string
	pkey   string
	client *ssh.Client
	port   int
	// protocol Protocol
}

func NewSshConn(host, user, pw, pkey string, port int) *sshConn {
	if port == 0 {
		port = 22
	}
	fmt.Println("Port is", port)
	lg := &sshConn{
		host: host,
		user: user,
		pw:   pw,
		pkey: pkey,
		port: port,
		// protocol: pr,
	}
	client := lg.Ping()
	lg.client = client
	return lg
}

func (lg *sshConn) Ping() *ssh.Client {
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
