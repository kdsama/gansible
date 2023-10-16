package main

import (
	"bytes"
	"fmt"
	"sync"

	"golang.org/x/crypto/ssh"
)

// Going to test out Os related information here

func main() {

	// THese are the cli commands that we need to use
	// cat
	port := 22
	hosts := []string{"localhost"}
	username := "kshitij"
	password := "admin@123"
	wg := sync.WaitGroup{}
	wg.Add(1)
	for _, host := range hosts {

		// host := "127.0.0.1"

		// Create an SSH configuration
		go func(host string) {
			defer wg.Done()
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

			cmds := []string{"env", "arch", "cat /proc/cpuinfo", "cat /proc/cmdline", "groups", "whoami", "id -g", "id -u", "df -h / | awk '{print $3, $4}' ", "free -h| awk '{print $3, $4}' "}
			if err != nil {
				panic(err)
			}

			ll := make([]byte, 20)
			mm := make([]byte, 20)
			sshOut := bytes.NewBuffer(ll)
			sshErr := bytes.NewBuffer(mm)
			type cmdOutput struct {
				In  string
				Out string
				Err string
			}
			// cmdOutputs := []cmdOutput{}
			// session.Stdout = sshOut
			var session *ssh.Session
			for _, cmd := range cmds {
				// Create a session
				session, err = client.NewSession()
				defer session.Close()
				if err != nil {

					return
				}
				session.Stdout = sshOut
				session.Stderr = sshErr

				session.Run(cmd)

				co := cmdOutput{
					Out: sshOut.String(),
					Err: sshErr.String(),
				}
				fmt.Println(co.Out)
				// cmdOutputs = append(cmdOutputs, co)

				sshOut.Truncate(0)
				sshErr.Truncate(0)

			}

		}(host)
	}
	wg.Wait()

}
