package shellcommands

import (
	"fmt"
	"os"
	"syscall"
)

type ExecCommand struct{}

func (pwd ExecCommand) Exec(args []string) {
	if ok := syscall.Exec(args[0], args[0:], os.Environ()); ok != nil {
		fmt.Println(ok.Error())
	}
}
