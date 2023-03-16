package shell

import (
	"fmt"
)

type CommandI interface {
	Exec(args []string)
}

type Shell struct {
	comands map[string]CommandI
}

func NewShell() *Shell {
	return &Shell{
		comands: make(map[string]CommandI),
	}
}

func (s *Shell) RegisterComand(name string, comand CommandI) {
	s.comands[name] = comand
}

func (s *Shell) ExecComand(name string, args []string) {
	if command, ok := s.comands[name]; ok {
		command.Exec(args)
	} else {
		fmt.Println("unknown comand")
	}
}

// func (s *Shell) ExecComands(names []string, args [][]string) {
// 	for _, name := range names {
// 		s.comands[name].Exec(args)
// 	}
// }
