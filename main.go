package main

import "github.com/kdsama/gansible/internal"

// "golang.org/x/crypto/ssh"

func main() {
	internal.New()
	internal.NewPlaybook("./pb.yml")

	// internal.New()
	// return
	// Replace these with your own SSH credentials

	// port := 22
	// hosts := []string{"localhost", "127.0.0.1", "0.0.0.0"}
	// username := "kshitij"
	// password := "admin@123"
	// cl := internal.NewLogin(hosts[0], username, password, "")
	// internal.GetFacts(cl)
	// wg := sync.WaitGroup{}
	// wg.Add(3)
	// for _, host := range hosts {

	// 	// host := "127.0.0.1"

	// 	// Create an SSH configuration
	// 	go func(host string) {
	// 		defer wg.Done()
	// 		config := &ssh.ClientConfig{
	// 			User: username,
	// 			Auth: []ssh.AuthMethod{
	// 				ssh.Password(password),
	// 			},
	// 			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// 		}

	// 		// Connect to the SSH server
	// 		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	// 		if err != nil {
	// 			fmt.Println("Failed to dial:", err)
	// 			return
	// 		}
	// 		defer client.Close()

	// 		cmds := []string{"ls -l", "cat lam.txt", "lasd", "pwd"}
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		ll := make([]byte, 20)
	// 		mm := make([]byte, 20)
	// 		sshOut := bytes.NewBuffer(ll)
	// 		sshErr := bytes.NewBuffer(mm)
	// 		type cmdOutput struct {
	// 			In  string
	// 			Out string
	// 			Err string
	// 		}
	// 		cmdOutputs := []cmdOutput{}
	// 		// session.Stdout = sshOut
	// 		var session *ssh.Session
	// 		for _, cmd := range cmds {
	// 			// Create a session
	// 			session, err = client.NewSession()
	// 			defer session.Close()
	// 			if err != nil {

	// 				return
	// 			}
	// 			session.Stdout = sshOut
	// 			session.Stderr = sshErr

	// 			session.Run(cmd)

	// 			co := cmdOutput{
	// 				Out: sshOut.String(),
	// 				Err: sshErr.String(),
	// 			}
	// 			cmdOutputs = append(cmdOutputs, co)

	// 			sshOut.Truncate(0)
	// 			sshErr.Truncate(0)

	// 		}
	// 		fmt.Printf("Host %s::::: %v", host, cmdOutputs)
	// 	}(host)
	// }
	// wg.Wait()
	// fmt.Println("Done, now Exiting ")
}
