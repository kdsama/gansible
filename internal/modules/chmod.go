package modules

import "fmt"

func NewMode(mode string, file string) string {
	return fmt.Sprintf("chmod %s %s", mode, file)
}
