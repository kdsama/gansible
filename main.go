package main

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
	// "golang.org/x/crypto/ssh"
)

func main() {
	// Replace these with your own SSH credentials
	username := "kshitij"
	password := "admin@123"
	host := "localhost"
	port := 22

	// Create an SSH configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer client.Close()

	cmds := []string{"ls -l", "cat lam.txt"}
	if err != nil {
		panic(err)
	}

	ll := make([]byte, 20)
	sshOut := bytes.NewBuffer(ll)
	// session.Stdout = sshOut

	for _, cmd := range cmds {
		// Create a session
		session, err := client.NewSession()
		defer session.Close()
		if err != nil {
			fmt.Println("Failed to create session:", err)
			return
		}
		session.Stdout = sshOut
		session.Run(cmd)
		// to := sshOut.Len()
		fmt.Println(sshOut.String())
		fmt.Println("????????????????????????????????????????????")
		sshOut.Truncate(0)
	}

	// fmt.Println("???????")

}
