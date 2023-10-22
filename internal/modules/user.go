package modules

import (
	"fmt"
	"strings"
)

func NewUser(userMap map[string]interface{}) (string, error) {
	var (
		str   string
		name  string
		state bool
	)
	if _, ok := userMap["name"]; !ok {
		return "", fmt.Errorf("name %w", ErrNotFound)
	}
	name = userMap["name"].(string)
	if _, ok := userMap["remove"]; ok {
		if userMap["remove"].(bool) {
			return fmt.Sprintf("userdel -r %s", name), nil
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
		str = fmt.Sprintf(" useradd  %s ||  passwd -u %s", name, name)
	} else {
		str = fmt.Sprintf("passwd -l %s", name)
	}

	if _, ok := userMap["groups"]; ok {
		groups := strings.Split(userMap["groups"].(string), ",")
		ns := fmt.Sprint(str)
		for _, g := range groups {
			ns = fmt.Sprintf("%s || groupadd %s", ns, g)
		}
		str = fmt.Sprintf("%s && usermod -aG %s %s", ns, userMap["groups"].(string), name)
	}

	if val, ok := userMap["create_home"]; ok {
		if val.(bool) {
			str = fmt.Sprintf("%s && mkdir /home/%s && chown %s:%s /home/%s", str, name, name, name, name)
		}
	}
	return str, nil
}
