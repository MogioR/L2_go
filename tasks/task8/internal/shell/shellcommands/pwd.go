package shellcommands

import (
	"fmt"
	"os"
)

type PWDCommand struct{}

func (pwd PWDCommand) Exec(args []string) {
	if len(args) > 0 {
		fmt.Println("Error: to many arguments!")
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
}
