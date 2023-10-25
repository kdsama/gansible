package internal

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type PlayBook struct {
	Plays []Play
}

type Play struct {
	Name  string                   `yaml:"name"`
	Hosts string                   `yaml:"hosts"`
	Tasks []map[string]interface{} `yaml:"tasks"`
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

func (pb *PlayBook) Run(index int) {
	currentPlay := pb.Plays[index]
	// Here there is another mistake 
	// The task may or may not be appended according to module being used. 
	// YOu forgot about it didnt you. 
	// Think what you need to do , and go drink water
	for currentPlay.Tasks
}
