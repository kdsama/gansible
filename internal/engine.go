package internal

import (
	"fmt"
	"sync"
)

type Engine struct {
	inventory     *MainInventory
	playbook      *PlayBook
	maxConcurrent int
	wg            *sync.WaitGroup
	sshService    ssher
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
		wg:            &sync.WaitGroup{},
		sshService:    NewSSHService(),
	}
}

func (e *Engine) Run() {
	// os := "ubuntu"
	fmt.Println(e.playbook.Plays)

	for i := range e.playbook.Plays {
		// fmt.Println("O ERROR", o.Err, strings.Trim(o.Err, " "))
		switch e.playbook.Plays[i].Strategy {
		case "free":
			e.FreeStrategy(i)
		default:
			e.LinearStrategy(i)
		}

	}
}

func (e *Engine) LinearStrategy(i int) {
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)
	if e.playbook.Plays[i].Serial > 0 {
		e.maxConcurrent = e.playbook.Plays[i].Serial
	}
	for _, h := range respObj.hosts {
		if _, ok := e.sshService.get(h); ok {
			obj := e.inventory.inv.All.Hosts[h]
			e.sshService.add(h, obj.SshHost, obj.SshUser, obj.SshPass, "", obj.SshPort)
		}
	}
	opts := []ExecOutput{}
	for k := 0; k < len(respObj.hosts)/e.maxConcurrent; k += e.maxConcurrent {
		start, end := k*e.maxConcurrent, ((k + 1) * e.maxConcurrent)
		if end > len(respObj.hosts) {
			end = len(respObj.hosts)
		}
		for _, t := range respObj.tasks {
			wg.Add(len(respObj.hosts))

			for _, h := range respObj.hosts[start:end] {
				h := h
				t := t
				go func() {
					defer wg.Done()
					for _, c := range t.cmds {
						opts = append(opts, e.sshService.execute(h, c))
					}

				}()
			}
		}
	}

	wg.Wait()

}

func (e *Engine) FreeStrategy(i int) {
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)
	if e.playbook.Plays[i].Serial > 0 {
		e.maxConcurrent = e.playbook.Plays[i].Serial
	}
	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		e.sshService.add(h, obj.SshHost, obj.SshUser, obj.SshPass, "", obj.SshPort)
	}
	wg.Add(len(respObj.hosts))
	opts := []ExecOutput{}
	for _, h := range respObj.hosts {
		h := h
		go func() {
			defer wg.Done()
			for _, t := range respObj.tasks {
				h := h

				for _, c := range t.cmds {

					fmt.Println(c)
					opts = append(opts, e.sshService.execute(h, c))
				}

			}
		}()
	}

	wg.Wait()

}
