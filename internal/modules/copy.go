package modules

import (
	"fmt"
	"log"
	"strings"
)

func NewCopy(task map[string]interface{}) string {
	// We need to do a remote copy
	// Without breaking the barrier
	// Can it just be done using commandline ?

	return "Hi"
}

// This is for copy from one place to another within remote. Not required
func newCopy(task map[string]interface{}) string {
	constraints := []string{"src", "dest"}
	for _, c := range constraints {
		if _, ok := task[c]; !ok {
			log.Fatalf("Error, %s not present", c)
		}
	}
	var cmd strings.Builder
	if _, ok := task["chdir"]; ok {

		cmd.WriteString(fmt.Sprintf("cd %s && ", task["chdir"].(string)))
	}
	cmd.WriteString(fmt.Sprintf("cp %s %s", task["src"].(string), task["dest"].(string)))
	return cmd.String()
}
