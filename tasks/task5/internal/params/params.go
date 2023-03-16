package params

import (
	"errors"
	"flag"
)

type Params struct {
	After  int
	Before int
	// Context    int
	Counting   bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func (p *Params) ParseArguments() error {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать +-N строк вокруг совпадения")
	count := flag.Bool("c", false, "подсчитать количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	if *context != 0 {
		if *after != 0 || *before != 0 {
			return errors.New("-A and -B mutually exclusive with -C")
		} else {
			*after = *context
			*before = *context
		}
	}

	p.After = *after
	p.Before = *before
	// p.Context = *context
	p.Counting = *count
	p.IgnoreCase = *ignoreCase
	p.Invert = *invert
	p.Fixed = *fixed
	p.LineNum = *lineNum

	return nil
}
