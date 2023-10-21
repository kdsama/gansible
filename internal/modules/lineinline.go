package modules

import (
	"fmt"
	"log"
)

func NewLineInFile(task map[string]interface{}) string {

	constraints := []string{"path", "line"}
	for _, k := range constraints {
		if _, ok := task[k]; !ok {
			log.Fatal("Damn it where is ", k)
		}
	}
	if _, ok := task["state"]; ok {
		var str string
		if task["state"] != "present" && task["state"] != "absent" {
			log.Fatal("Invalid input : ", task["state"])
		}

		str = fmt.Sprintf("if grep -q %s %s; then echo 'true'; else echo 'false'; fi", task["line"].(string), task["path"].(string))
		if task["state"] == "absent" {
			str = fmt.Sprintf("if grep -q %s %s; then echo 'false'; else echo 'true'; fi", task["line"].(string), task["path"].(string))
		}

		return str
	}
	if _, ok := task["append"]; !ok {
		log.Fatal("Expected append function")
	}
	return fmt.Sprintf("%s >> %s", task["line"].(string), task["path"].(string))

}
