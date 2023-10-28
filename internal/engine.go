package internal

import (
	"fmt"
	"sync"
)

type Engine struct {
	inventory     *MainInventory
	playbook      *PlayBook
	maxConcurrent int
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
		inventory:     inventory,
		playbook:      pb,
		maxConcurrent: 10,
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
	// So  I need to run this with x concurrency
	// If its length is less than x, we run as is
	// Else,
	// lets say we have 22
	// 0 to 10 , 11 to 20, 21 to 22
	// 0:11
	//11:21
	//21:22

	for _, t := range respObj.tasks {
		wg.Add(len(respObj.hosts))

		for _, h := range respObj.hosts {
			h := h
			t := t
			go func() {
				defer wg.Done()
				for _, c := range t.cmds {
					cache[h].execute(c)
				}

			}()
		}
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	wg.Wait()

}

func (e *Engine) FreeStrategy(i int, cache map[string]*sshConn) {
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)
	if e.playbook.Plays[i].Serial > 0 {
		e.maxConcurrent = e.playbook.Plays[i].Serial
	}
	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		cache[h] = NewSshConn(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
	}
	wg.Add(len(respObj.hosts))
	for k := 0; k < len(respObj.hosts)/e.maxConcurrent; k += e.maxConcurrent {
		start, end := k*e.maxConcurrent, ((k + 1) * e.maxConcurrent)
		if end > len(respObj.hosts) {
			end = len(respObj.hosts)
		}
		for _, h := range respObj.hosts[start:end] {
			h := h
			go func() {
				defer wg.Done()
				for _, t := range respObj.tasks {
					h := h
					for _, c := range t.cmds {
						cache[h].execute(c)
					}

				}
			}()
		}
	}
	wg.Wait()
}
