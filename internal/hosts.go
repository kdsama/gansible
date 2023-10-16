package internal

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// all:  # This is the top-level group, which contains all hosts
//   hosts:
//     host1.example.com:  # Individual host
//       ansible_host: 192.168.1.1  # IP address or hostname
//       ansible_user: user1  # SSH username
//       ansible_ssh_pass: password1  # SSH password (not recommended, use SSH keys)
//     host2.example.com:
//       ansible_host: 192.168.1.2
//       ansible_user: user2
//       ansible_ssh_pass: password2

//   children:  # You can define groups of hosts
//     web_servers:
//       hosts:
//         web1.example.com:
//         web2.example.com:
//     db_servers:
//       hosts:
//         db1.example.com:
//         db2.example.com:

type Host struct {
	AnsibleHost    string `yaml:"ansible_host"`
	AnsibleUser    string `yaml:"ansible_user"`
	AnsibleSshPass string `yaml:"ansible_ssh_pass"`
}
type MainInventory struct {
	inv *Inventory
}

type Inventory struct {
	All struct {
		Hosts    map[string]Host `yaml:"hosts"`
		Children map[string]struct {
			Hosts map[string]struct{} `yaml:"hosts"`
		} `yaml:"children"`
	} `yaml:"all"`
}

func New() {
	// We need to check in Command line, Current Directory and and default location
	// and merge all of them.
	mainInv := MainInventory{}

	// Argument, Current Folder and default
	inventoryArr := []string{"", "./ff.txt", "./ff1.txt"}

	for _, inv := range inventoryArr {
		if strings.Trim(inv, " ") == "" {
			continue
		}
		yamlData, err := os.ReadFile(inv)

		if err != nil {
			log.Println(err)
		}
		var inventory Inventory
		err = yaml.Unmarshal(yamlData, &inventory)
		if err != nil {
			return
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
	fmt.Println(mainInv.inv.All.Hosts)
	// Unmarshal the YAML into the Go struct

	// You can now work with the 'inventory' struct as needed
	// validate for cyclic dependency

}
