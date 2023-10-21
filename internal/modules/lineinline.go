package modules

import (
	"fmt"
	"log"
)

func NewLineInFile(task map[string]interface{}) string {

	var (
		constraints = []string{"path", "line"}
		owner       string
		group       string
		query       string
		path        string
	)

	for _, k := range constraints {
		if _, ok := task[k]; !ok {
			log.Fatal("Damn it where is ", k)
		}
	}
	path = task["path"].(string)
	if _, ok := task["state"]; ok {
		var str string
		if task["state"] != "present" && task["state"] != "absent" {
			log.Fatal("Invalid input : ", task["state"])
		}

		str = fmt.Sprintf("if grep -q %s %s; then echo 'true'; else echo 'false'; fi", task["line"].(string), path)
		if task["state"] == "absent" {
			str = fmt.Sprintf("if grep -q %s %s; then echo 'false'; else echo 'true'; fi", task["line"].(string), path)
		}

		return str
	}
	if _, ok := task["append"]; !ok {
		log.Fatal("Expected append function")
	}
	query = fmt.Sprintf("%s >> %s", task["line"].(string), path)
	if _, ok := task["group"]; ok {
		group = task["group"].(string)
	}
	if _, ok := task["owner"]; ok {
		owner = task["owner"].(string)
	}
	if owner != "" || group != "" {
		ns := NewOwner(owner, group, path)
		query = fmt.Sprintf("%s && %s", query, ns)
	}

	if _, ok := task["mod"]; ok {
		ns := NewMode(task["mod"].(string), path)
		query = fmt.Sprintf("%s && %s", query, ns)
	}
	return query

}
