package shellcommands

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (pwd EchoCommand) Exec(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: empty arguments!")
		return
	}

	fmt.Println(strings.Join(args, " "))
}
