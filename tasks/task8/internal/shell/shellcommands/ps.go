package shellcommands

import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
)

type PSCommand struct{}

func (pwd PSCommand) Exec(args []string) {
	var proceses []ps.Process

	if len(args) > 0 {
		fmt.Println("Error: to many arguments!")
		return
	}

	proceses, ok := ps.Processes()
	if ok != nil {
		fmt.Println(ok.Error())
		return
	}

	fmt.Println("pid\tproc")
	for _, proc := range proceses {
		fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
	}
}
