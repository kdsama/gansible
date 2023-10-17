package internal

import (
	"golang.org/x/crypto/ssh"
)

var factList = []string{"env", "arch", "cat /proc/cpuinfo", "cat /proc/cmdline", "groups", "whoami",
	"id -g", "id -u", "df -h / | awk '{print $3, $4}' ", "free -h| awk '{print $3, $4}' "}

func GetFacts(client *ssh.Client) *[]ExecOutput {

	cmdOutputs := []ExecOutput{}

	for _, cmd := range factList {
		// Create a session
		co := execute(client, cmd)
		cmdOutputs = append(cmdOutputs, co)

	}
	return &cmdOutputs
}
