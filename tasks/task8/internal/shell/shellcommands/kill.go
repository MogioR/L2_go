package shellcommands

import (
	"fmt"
	"os"
	"strconv"
)

type KillCommand struct{}

func (pwd KillCommand) Exec(args []string) {

	if len(args) == 0 {
		fmt.Println("Error: hasn't targets!")
		return
	}

	for _, v := range args[1:] {
		if pid, ok := strconv.Atoi(v); ok != nil {
			fmt.Println("kill: pid:", v, "is not valid")
		} else {
			if proc, ok := os.FindProcess(pid); ok != nil {
				fmt.Println(pid, "not found")
			} else if ok = proc.Kill(); ok != nil {
				fmt.Println(ok.Error())
			} else {
				fmt.Println(pid, "has killed")
			}
		}
	}
}
