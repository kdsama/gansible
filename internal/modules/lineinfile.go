package modules

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
)

func NewLineInFile(task map[string]interface{}) ([]string, error) {

	var (
		constraints = []string{"path", "line"}
		path        string
		result      []string
	)

	for _, k := range constraints {
		if _, ok := task[k]; !ok {
			return result, fmt.Errorf("%s %w", k, ErrNotFound)
		}
	}
	path = task["path"].(string)
	if _, ok := task["state"]; ok {
		if task["state"] != "present" && task["state"] != "absent" {
			return result, fmt.Errorf("%w for %s", ErrInvalidInput, "state")
		}

		if task["state"] == "absent" {
			result = append(result, fmt.Sprintf("if grep -q %s %s; then echo 'false'; else echo 'true'; fi", task["line"].(string), path))
		} else if task["state"] == "present" {
			result = append(result, fmt.Sprintf("if grep -q %s %s; then echo 'true'; else echo 'false'; fi", task["line"].(string), path))
		} else {
			fmt.Println("Invalid Input --> skipping")
		}

		return result, nil
	}

	// TODO: move this to file.go as lineinfile does not change ownershup of the file
	result = append(result, fmt.Sprintf("echo \"%s\" >> %s", task["line"].(string), path))

	return result, nil

}
