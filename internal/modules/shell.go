package modules

import "fmt"

func NewShell(shellMap map[string]interface{}) ([]string, error) {

	var cmds []string
	if _, ok := shellMap[chdirParam]; ok {
		cmds = append(cmds, fmt.Sprintf("cd %s", shellMap[chdirParam].(string)))
	}
	if _, ok := shellMap[cmdParam]; !ok {
		return cmds, fmt.Errorf("command %w", ErrNotFound)
	}

	if _, ok := shellMap[executableParam]; ok {
		cmds = append(cmds, fmt.Sprintf("chsh -s %s", shellMap[executableParam].(string)))
	}
	cmds = append(cmds, shellMap[cmdParam].(string))

	return cmds, nil
}
