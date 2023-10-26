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

type Document struct {
	hosts []string
	tasks []*Task
}

func (pb *PlayBook) Generate(index int) Document {

	// Here there is another mistake
	// The task may or may not be appended according to module being used.
	// YOu forgot about it didnt you.
	// Think what you need to do , and go drink water
	// so I have an idea
	// We are going to gather facts at the start of each play anyway
	// So there is no trouble tbf
	// The thing is
	// We already have a command
	// How are we going to skip an execution
	// Now this becomes our problem
	// Maybe what we can do is provide playbook with facts first
	// Thats not right
	// Or we should return OS information alongside tasks
	// There can be an enumeration for that
	// If OS matches, only then we should process it, or we should skip
	// But what if we want to actually skip in case of where clause
	// We can add another field for that
	// But that field is not exclusive to just server skipping, it can be anycondition
	// Should I not consider the where clause ?
	// Lets do that , Lets skip the where clause
	// Now there is another concern
	// We would also have to return the hosts we must run this on
	// Lets first gather everything
	hosts := strings.Split(pb.Plays[index].Hosts, ",")
	fn := []*Task{}
	for _, task := range pb.Plays[index].Tasks {
		t, err := parseTask(task)
		if err != nil || t == nil {
			continue
		}
		fn = append(fn, t)
	}
	return Document{
		hosts: hosts,
		tasks: fn,
	}
}
