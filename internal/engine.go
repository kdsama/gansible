package internal

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/ssh"
)

type Engine struct {
	inventory *MainInventory
	playbook  *PlayBook
}

func NewEngine(playbookPath string, hostPath string) *Engine {

	var (
		inventory *MainInventory
		pb        *PlayBook
	)
	if hostPath != "" {
		opts := InventoryPathOptions(hostPath)
		inventory = NewInventory(opts)
	} else {
		inventory = NewInventory()
	}

	pb = NewPlaybook(playbookPath)

	return &Engine{
		inventory: inventory,
		playbook:  pb,
	}
}

func (e *Engine) Run() {
	// os := "ubuntu"
	cache := map[string]*ssh.Client{}

	wg := sync.WaitGroup{}
	for i := range e.playbook.Plays {
		respObj := e.playbook.Generate(i)
		fmt.Println(respObj.hosts)
		for _, h := range respObj.hosts {
			obj := e.inventory.inv.All.Hosts[h]
			fmt.Println("Obj is ", obj)
			cache[h] = NewLogin(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
		}
		for _, t := range respObj.tasks {
			wg.Add(len(respObj.hosts))

			for _, h := range respObj.hosts {
				h := h
				go func() {
					defer wg.Done()
					for _, c := range t.cmds {
						execute(cache[h], c)
					}

				}()
			}
			wg.Wait()
			fmt.Println("Finished")
		}
	}
}
