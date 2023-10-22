package modules

import "fmt"

func NewFileOwner(owner, group string, file string) string {
	var str string

	if owner != "" {
		str += owner
	}
	if group != "" {
		str += fmt.Sprintf(":%s", group)
	}

	return fmt.Sprintf("chown %s %s", str, file)
}

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
	if _, ok := userMap["name"]; !ok {
		if _, ok := userMap["remove"]; ok {
			if userMap["remove"].(bool) {
				return fmt.Sprintf("userdel -r %s", name), nil
			}
		}
	}

	if _, ok := userMap["state"]; !ok {
		state = true
	} else {
		state = userMap["state"].(bool)
	}
	if state {
		str = fmt.Sprintf("sudo useradd  %s || sudo passwd -u %s", name, name)
	} else {
		str = fmt.Sprintf("passwd -l %s", name)
	}
	if val, ok := userMap["create_home"]; ok {
		if val.(bool) {
			str = fmt.Sprintf("%s && mkdir /home/%s && chown %s:%s /home/%s", str, name, name, name, name)
		}
	}

	return str, nil
}
