package shellcommands

import (
	"fmt"
	"os"
)

type CDCommand struct{}

func (pwd CDCommand) Exec(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: don't have root!")
		return
	}
	if len(args) > 1 {
		fmt.Println("Error: to many arguments!")
		return
	}

	var root string
	if args[0] == ".." {
		root = "../"
	} else {
		root = args[0]
	}

	err := os.Chdir(root)
	if err != nil {
		fmt.Println(err)
		return
	}
}
