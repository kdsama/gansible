package internal

import (
	"fmt"
	"testing"
)

type mockSshService struct {
	hosts map[string]bool
}

func (ms *mockSshService) add(name, host, user, pw, pkey string, port int) {
	ms.hosts[name] = true
}

func (ms *mockSshService) get(name string) (*sshConn, bool) {
	return nil, true
}

func (ms *mockSshService) getOS(name string) string {
	return "any"
}

func (ms *mockSshService) execute(name string, cmd string) (ExecOutput, error) {
	return ExecOutput{
		Cmd: fmt.Sprintf("cmd %s", name),
		Out: name,
		Err: "",
	}, nil
}
func testSetupEngine() *Engine {
	eg := Engine{}
	eg.inventory = &MainInventory{
		inv: &Inventory{
			All: struct {
				Hosts    map[string]Host "yaml:\"hosts\""
				Children map[string]struct {
					Hosts map[string]struct{} "yaml:\"hosts\""
				} "yaml:\"children\""
			}{
				Hosts: map[string]Host{
					"server": {
						SshHost: "localhost",
						SshUser: "someuser",
						SshPass: "SomePass",
						SshPort: 22,
					},
					"server1": {
						SshHost: "localhost1",
						SshUser: "someuser1",
						SshPass: "SomePass1",
						SshPort: 22,
					},
				},
			},
		},
	}
	eg.playbook = &PlayBook{
		Plays: []Play{
			{
				Name:  "playname",
				Hosts: "localhost",
				Tasks: []map[string]interface{}{
					{
						"lineinfile": map[string]interface{}{
							"path": "dummyPath",
							"line": "dummyLine",
						},
					},
					{
						"user": map[string]interface{}{
							"name":        "kd",
							"state":       "present",
							"create_home": true,
						},
					},
				},
			},
		},
	}
	eg.sshService = &mockSshService{}
	return &eg
}

func TestFreeStrategy(t *testing.T) {

}

func TestLinearStrategy(t *testing.T) {

}
