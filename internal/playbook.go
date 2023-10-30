package internal

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type PlayBook struct {
	Plays []Play
}

type Play struct {
	Name     string                   `yaml:"name"`
	Hosts    string                   `yaml:"hosts"`
	Tasks    []map[string]interface{} `yaml:"tasks"`
	Strategy string                   `yaml:"strategy"`
	Serial   int                      `yaml:"serial"`
}

func NewPlaybook(playbookPath string) *PlayBook {
	data, err := os.ReadFile(playbookPath)
	if err != nil {
		log.Fatal(err)
	}
	var Plays []Play
	err = yaml.Unmarshal(data, &Plays)
	if err != nil {
		log.Fatal(err)
	}
	// There is no such thing as playbooks in parallel
	// Main thing this layer would need to do in the future is the load all the tasks and variables and all the other things according to Options given
	// Or we can have separate files for that as well
	// But what I am talking about is include and roles and all that stuff
	// The execution strategy can be given to the engine for now
	// So the engine will do it on all hosts the way it wants to
	// So we are returning several plays
	// These plays are supposed to run sequentially, Remember that.
	//

	return &PlayBook{
		Plays: Plays,
	}
}

type PlayDoc struct {
	hosts []string
	tasks []*Task
}

func (pb *PlayBook) Generate(index int) PlayDoc {
	hosts := strings.Split(pb.Plays[index].Hosts, ",")
	fn := []*Task{}
	for _, task := range pb.Plays[index].Tasks {
		t, err := parseTask(task)
		if err != nil || t == nil {
			continue
		}
		fn = append(fn, t)
	}
	return PlayDoc{
		hosts: hosts,
		tasks: fn,
	}
}
