package modules

import "fmt"

func NewMode(mode string, file string) string {
	fmt.Println("Are we here ")
	return fmt.Sprintf("chown %s %s", mode, file)
}
