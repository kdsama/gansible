package internal

import (
	"fmt"

	"github.com/kdsama/gansible/internal/modules"
)

func parseTask(task map[string]interface{}) string {

	var res string
	for key, _ := range task {
		switch key {
		case "copy":
			res = modules.NewCopy(task[key].(map[string]interface{}))
		case "lineinfile":
			res, _ = modules.NewLineInFile(task[key].(map[string]interface{}))
		default:
			fmt.Println("Yeah its fine for now")
		}
	}
	return res
}
