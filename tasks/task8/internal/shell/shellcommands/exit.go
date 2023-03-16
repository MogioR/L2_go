package shellcommands

import "os"

type ExitCommand struct{}

func (pwd ExitCommand) Exec(args []string) {
	os.Exit(0)
}
