package internal

import (
	"fmt"
	"sync"
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
	cache := map[string]*sshConn{}

	for i := range e.playbook.Plays {
		// fmt.Println("O ERROR", o.Err, strings.Trim(o.Err, " "))
		switch e.playbook.Plays[i].Strategy {
		case "free":
			e.FreeStrategy(i, cache)
		default:
			e.LinearStrategy(i, cache)
		}

	}
}

func (e *Engine) LinearStrategy(i int, cache map[string]*sshConn) {
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)

	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		cache[h] = NewSshConn(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
	}
	for _, t := range respObj.tasks {
		wg.Add(len(respObj.hosts))

		for _, h := range respObj.hosts {
			h := h
			go func() {
				defer wg.Done()
				for _, c := range t.cmds {
					cache[h].execute(c)
				}

			}()
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		wg.Wait()

	}
}

func (e *Engine) FreeStrategy(i int, cache map[string]*sshConn) {
	fmt.Println("Free Strategy Start")
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)

	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		cache[h] = NewSshConn(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
	}
	wg.Add(len(respObj.hosts))
	for _, h := range respObj.hosts {
		h := h
		go func() {
			defer func() {
				fmt.Println("Host", h, "Finished")
				wg.Done()
			}()
			for _, t := range respObj.tasks {
				h := h
				for _, c := range t.cmds {
					cache[h].execute(c)
				}

			}
		}()
	}
	wg.Wait()
	fmt.Println("Free Strategy ENd ---------------------------------------------------------")
}
