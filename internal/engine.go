package internal

import (
	"fmt"
	"strings"
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

func (e *Engine) LinearStrategy(i int, cache map[string]*ssh.Client) {
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)

	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		cache[h] = NewLogin(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
	}
	for _, t := range respObj.tasks {
		wg.Add(len(respObj.hosts))

		for _, h := range respObj.hosts {
			h := h
			go func() {
				defer wg.Done()
				for _, c := range t.cmds {
					o := execute(cache[h], c)

					if strings.Trim(o.Err, " ") != "" {

					}

				}

			}()
		}
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		wg.Wait()

	}
}

func (e *Engine) FreeStrategy(i int, cache map[string]*ssh.Client) {
	fmt.Println("Free Strategy Start")
	wg := sync.WaitGroup{}
	respObj := e.playbook.Generate(i)

	for _, h := range respObj.hosts {
		obj := e.inventory.inv.All.Hosts[h]
		cache[h] = NewLogin(obj.AnsibleHost, obj.AnsibleUser, obj.AnsibleSshPass, "", obj.AnsiblePort)
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
					execute(cache[h], c)
					// if strings.Trim(o.Err, " ") != "" {
					// }

				}

			}
		}()
	}
	wg.Wait()
	fmt.Println("Free Strategy ENd ---------------------------------------------------------")
}
