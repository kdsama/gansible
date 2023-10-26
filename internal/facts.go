package internal

var factList = []string{"/bin/bash -ic '/usr/bin/env'"}

func GetFacts(client *sshConn) *[]ExecOutput {

	cmdOutputs := []ExecOutput{}

	for _, cmd := range factList {
		// Create a session
		co := client.execute(cmd)
		cmdOutputs = append(cmdOutputs, co)

	}
	return &cmdOutputs
}
