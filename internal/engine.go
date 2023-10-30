package internal

import (
	"fmt"
	"strings"
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

	for i := range e.playbook.Plays {

		respObj := e.prepareTasks(i)

		switch e.playbook.Plays[i].Strategy {
		case "free":
			e.FreeStrategy(respObj)
		default:
			e.LinearStrategy(respObj)
		}

	}
}

func (e *Engine) LinearStrategy(respObj PlayDoc) {

	opts := []ExecOutput{}
	for k := 0; k < len(respObj.hosts)/e.maxConcurrent; k += e.maxConcurrent {
		start, end := k*e.maxConcurrent, ((k + 1) * e.maxConcurrent)
		if end > len(respObj.hosts) {
			end = len(respObj.hosts)
		}
		for _, t := range respObj.tasks {
			e.wg.Add(len(respObj.hosts))

			for _, h := range respObj.hosts[start:end] {
				h := h
				t := t
				if !e.sameOS(t, h) {
					continue
				}
				go func() {
					defer e.wg.Done()
					for _, c := range t.cmds {
						res, err := e.sshService.execute(h, c)
						if err != nil {
							fmt.Println("Needs to be skipped")
							continue
						}
						if strings.Trim(res.Err, " ") != "" && !t.skip_errors {

							break
						}
						opts = append(opts, res)
					}

				}()
			}
		}
	}

	e.wg.Wait()

}

func (e *Engine) FreeStrategy(respObj PlayDoc) {

	e.wg.Add(len(respObj.hosts))
	opts := []ExecOutput{}
	for _, h := range respObj.hosts {
		h := h
		go func() {
			defer e.wg.Done()
			for _, t := range respObj.tasks {
				h := h
				if !e.sameOS(t, h) {
					continue
				}

				for _, c := range t.cmds {
					res, err := e.sshService.execute(h, c)
					if err != nil {
						fmt.Println("Needs to be skipped")
						continue
					}
					if strings.Trim(res.Err, " ") != "" && !t.skip_errors {
						break
					}
					opts = append(opts, res)
				}

			}
		}()
	}

	e.wg.Wait()

}

func (e *Engine) sameOS(t *Task, h string) bool {
	if t.os != "any" && t.os != strings.ToLower(e.sshService.getOS(h)) {
		return false
	}
	return true
}

func (e *Engine) prepareTasks(i int) PlayDoc {
	respObj := e.playbook.Generate(i)
	if e.playbook.Plays[i].Serial > 0 {
		e.maxConcurrent = e.playbook.Plays[i].Serial
	}
	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		e.sshService.add(h, obj.SshHost, obj.SshUser, obj.SshPass, "", obj.SshPort)
	}
	return respObj
}
