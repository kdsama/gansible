package internal

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Playbook struct {
	Name  string                   `yaml:"name"`
	Hosts string                   `yaml:"hosts"`
	Tasks []map[string]interface{} `yaml:"tasks"`
}

func NewPlaybook(playfile string) {
	data, e := os.ReadFile(playfile)
	// fmt.Println(string(data))
	if e != nil {
		log.Fatal(e)
	}
	var pb []Playbook
	e = yaml.Unmarshal(data, &pb)
	if e != nil {
		log.Fatal(e)
	}
	taskArray := [][]string{}
	for _, play := range pb {
		for _, task := range play.Tasks {
			result := parseTask(task)
			taskArray = append(taskArray, result)
		}
	}

	fmt.Println(taskArray)
}
