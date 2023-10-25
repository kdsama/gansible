package internal

import (
	"fmt"

	"github.com/kdsama/gansible/internal/modules"
)

type Task struct {
	cmds []string
	os   string
}

func parseTask(task map[string]interface{}) (*Task, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// TODO: Work the other task level variables that may be present
	var result *Task
	for key, _ := range task {
		switch key {
		// case "copy":
		// 	res = modules.NewCopy(task[key].(map[string]interface{}))
		case "lineinfile":
			cmds, err := modules.NewLineInFile(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = &Task{
				cmds: cmds,
				os:   "any",
			}
		case "file":
			cmds, err := modules.NewFilePermissions(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = &Task{
				cmds: cmds,
				os:   "any",
			}
		case "user":
			cmds, err := modules.NewUser(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = &Task{
				cmds: cmds,
				os:   "any",
			}
		case "shell":
			cmds, err := modules.NewShell(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result = &Task{
				cmds: cmds,
				os:   "any",
			}

		case "default":
			fmt.Println(key)
		}
	}
	return result, nil
}
