package internal

import (
	"fmt"
	"os"

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

type Inventory struct {
	All struct {
		Hosts    map[string]Host `yaml:"hosts"`
		Children map[string]struct {
			Hosts map[string]struct{} `yaml:"hosts"`
		} `yaml:"children"`
	} `yaml:"all"`
}

func New() {
	yamlData, _ := os.ReadFile("./ff.txt")
	// Unmarshal the YAML into the Go struct
	var inventory Inventory
	err := yaml.Unmarshal(yamlData, &inventory)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can now work with the 'inventory' struct as needed
	fmt.Printf("%+v\n", inventory.All.Children)
}
