package internal

import (
	"fmt"

	"github.com/kdsama/gansible/internal/modules"
)

type Task struct {
	name        string
	cmds        []string
	skip_errors bool
	os          string
}

func parseTask(task map[string]interface{}) (*Task, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// TODO: Work the other task level variables that may be present
	var result = &Task{}
	result.os = "any"
	for key, _ := range task {
		switch key {
		// case "copy":
		// 	res = modules.NewCopy(task[key].(map[string]interface{}))
		case LineinfileMod:
			cmds, err := modules.NewLineInFile(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result.cmds = cmds

		case fileMod:
			cmds, err := modules.NewFilePermissions(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result.cmds = cmds

		case userMod:
			cmds, err := modules.NewUser(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result.cmds = cmds
		case shellMod:
			cmds, err := modules.NewShell(task[key].(map[string]interface{}))
			if err != nil {
				return result, err
			}
			result.cmds = cmds
		case "skip_errors":
			result.skip_errors = true
		case "name":
			result.name = task[key].(string)
		case "default":
			fmt.Println(key)
		}
	}
	return result, nil
}
