package modules

import "fmt"

func NewShell(shellMap map[string]interface{}) ([]string, error) {

	var cmds []string
	if _, ok := shellMap["chdir"]; ok {
		cmds = append(cmds, fmt.Sprintf("cd %s", shellMap["chdir"].(string)))
	}
	if _, ok := shellMap["cmd"]; !ok {
		return cmds, fmt.Errorf("command %w", ErrNotFound)
	}

	if _, ok := shellMap["executable"]; ok {
		cmds = append(cmds, fmt.Sprintf("chsh -s %s", shellMap["executable"].(string)))
	}
	cmds = append(cmds, shellMap["cmd"].(string))

	return cmds, nil
}
