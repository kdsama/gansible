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

func (ms *mockSshService) execute(name string, cmd string) ExecOutput {
	return ExecOutput{
		Cmd: fmt.Sprintf("cmd %s", name),
		Out: name,
		Err: "",
	}
}

func TestFreeStrategy(t *testing.T) {

}

func TestLinearStrategy(t *testing.T) {

}
