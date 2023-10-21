package modules

import (
	"fmt"
	"log"
	"strings"
)

func NewCopy(task map[string]interface{}) string {
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
