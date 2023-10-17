package main

import "github.com/kdsama/gansible/internal"

// Going to test out Os related information here

func main() {

	hosts := []string{"localhost"}
	username := "kshitij"
	password := "admin@123"
	client := internal.NewLogin(hosts[0], username, password, "")
	defer client.Close()
	internal.GetFacts(client)

}
