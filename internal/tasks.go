package internal

import (
	"fmt"

	"github.com/kdsama/gansible/internal/modules"
)

type Task struct {
	cmds []string
}

func parseTask(task map[string]interface{}) ([]*Task, error) {

	var result []*Task
	for key, _ := range task {
		switch key {
		// case "copy":
		// 	res = modules.NewCopy(task[key].(map[string]interface{}))
		case "lineinfile":
			cmds, err := modules.NewLineInFile(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = append(result, &Task{cmds: cmds})
		case "file":
			cmds, err := modules.NewFilePermissions(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = append(result, &Task{cmds: cmds})
		case "user":
			cmds, err := modules.NewUser(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = append(result, &Task{cmds: cmds})
		case "shell":
			cmds, err := modules.NewShell(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = append(result, &Task{cmds: cmds})
		default:
			fmt.Println("Yeah its fine for now")
		}
	}
	return result, nil
}
