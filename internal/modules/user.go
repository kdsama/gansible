package modules

import (
	"fmt"
	"strings"
)

func NewUser(userMap map[string]interface{}) ([]string, error) {
	var (
		result []string
		name   string
		state  bool
	)
	if _, ok := userMap["name"]; !ok {
		return result, fmt.Errorf("name %w", ErrNotFound)
	}
	name = userMap["name"].(string)
	if _, ok := userMap["remove"]; ok {
		if userMap["remove"].(bool) {
			return []string{fmt.Sprintf("userdel -r %s", name)}, nil
		}
	}

	if _, ok := userMap["state"]; !ok {
		state = true
	} else {
		if userMap["state"].(string) == "absent" {
			state = false
		} else {
			state = true
		}
	}
	if state {
		result = append(result, fmt.Sprintf("useradd %s", name), fmt.Sprintf("passwd -u %s", name))

	} else {
		result = append(result, fmt.Sprintf("passwd -l %s", name))
	}

	if _, ok := userMap["groups"]; ok {
		groups := strings.Split(userMap["groups"].(string), ",")

		for _, g := range groups {
			result = append(result, fmt.Sprintf("groupadd %s", g))
		}
		result = append(result, fmt.Sprintf("usermod -aG %s %s", userMap["groups"].(string), name))
	}

	if val, ok := userMap["create_home"]; ok {
		if val.(bool) {
			result = append(result, fmt.Sprintf("mkdir /home/%s", name), fmt.Sprintf("chown %s:%s /home/%s", name, name, name))
		}
	}
	return result, nil
}
