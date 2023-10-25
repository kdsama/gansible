package internal

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Host struct {
	AnsibleHost    string `yaml:"ansible_host"`
	AnsibleUser    string `yaml:"ansible_user"`
	AnsibleSshPass string `yaml:"ansible_ssh_pass"`
	AnsiblePort    int    `yaml:"ansible_ssh_port"`
}
type MainInventory struct {
	inv   *Inventory
	paths []string
}

type InventoryOpts func(inv *MainInventory)

type Inventory struct {
	All struct {
		Hosts    map[string]Host `yaml:"hosts"`
		Children map[string]struct {
			Hosts map[string]struct{} `yaml:"hosts"`
		} `yaml:"children"`
	} `yaml:"all"`
}

func InventoryPathOptions(path string) InventoryOpts {
	return func(inv *MainInventory) {
		inv.paths = append(inv.paths, path)
	}
}

func NewInventory(invOpts ...InventoryOpts) *MainInventory {
	// We need to check in Command line, Current Directory and and default location
	// and merge all of them.
	mainInv := MainInventory{}
	for _, opts := range invOpts {
		opts(&mainInv)
	}
	// append the default path as well
	mainInv.paths = append(mainInv.paths, "ff.yml")
	// Argument, Current Folder and default

	for _, inv := range mainInv.paths {
		if strings.Trim(inv, " ") == "" {
			continue
		}
		yamlData, _ := os.ReadFile(inv)

		// if err != nil {
		// 	log.Println(err)
		// }
		var inventory Inventory
		err := yaml.Unmarshal(yamlData, &inventory)
		if err != nil {
			log.Fatal(err)
		}
		if mainInv.inv == nil {
			mainInv.inv = &inventory
		} else {
			// Add keys that are not present to mainInv from Hosts
			for k, v := range inventory.All.Hosts {
				if _, ok := mainInv.inv.All.Hosts[k]; !ok {
					mainInv.inv.All.Hosts[k] = v

				}
			}
			// Add Children keys as well
			for k, v := range inventory.All.Children {
				if _, ok := mainInv.inv.All.Children[k]; !ok {
					mainInv.inv.All.Children[k] = v
				}
			}
		}

	}
	mainInv.Validate()
	return &mainInv
	// Unmarshal the YAML into the Go struct
}

func (mv *MainInventory) Validate() {
	// We need to check if there is any cyclic dependency in children
	// There cant be one in hosts

	visited := map[string]struct{}{}

	var dfs func(node string) bool
	dfs = func(node string) bool {
		if _, ok := visited[node]; ok {
			return false
		}
		visited[node] = struct{}{}
		for k, _ := range mv.inv.All.Children[node].Hosts {
			if !dfs(k) {
				log.Fatal("cyclic dependency found at child: ", k)
			}
		}
		return true
	}
	for k, _ := range mv.inv.All.Children {
		visited = map[string]struct{}{}
		dfs(k)
	}
	// Have to visualise better for the cyclic dependency on the children here
}
