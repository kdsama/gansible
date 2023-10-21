package modules

import "fmt"

func NewOwner(owner, group string, file string) string {
	var str string

	if owner != "" {
		str += owner
	}
	if group != "" {
		str += fmt.Sprintf(":%s", group)
	}
	return fmt.Sprintf("chown %s %s", str, file)
}
