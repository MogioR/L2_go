package params

import "flag"

type Params struct {
	SortColumn       int
	SortByNums       bool
	SortReverse      bool
	RemoveDuplicates bool
	// M bool
	// b bool
	// c bool
	// h bool
}

func (p *Params) ParseArguments() {
	k := flag.Int("k", -1, "Указание колонки для сортировки (слова в строке могут выступатьв качестве колонок, по умолчанию разделитель — пробел)")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Сортировать в обратном порядке")

	flag.Parse()

	p.SortColumn = *k
	p.SortByNums = *n
	p.SortReverse = *r
	p.RemoveDuplicates = *u
}
