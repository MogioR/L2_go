package stringsorter

import (
	"sort"
	"strconv"
	"strings"
)

type StringSorter struct {
	strings    []string
	sortColumn int
	sortByNums bool
}

func NewStringSorter(strings []string, sortColumn int, sortByNums bool) *StringSorter {
	return &StringSorter{
		strings:    strings,
		sortColumn: sortColumn,
		sortByNums: sortByNums,
	}
}
func (ss StringSorter) Len() int      { return len(ss.strings) }
func (ss StringSorter) Swap(i, j int) { ss.strings[i], ss.strings[j] = ss.strings[j], ss.strings[i] }
func (ss StringSorter) stringsLess(a, b string) bool {
	if !ss.sortByNums {
		return a < b // Сравниваем обычные строки
	} else { // Сравниваем строки как числа
		a, err := strconv.Atoi(a)
		if err != nil {
			return true
		}
		b, err := strconv.Atoi(b)
		if err != nil {
			return false
		}
		return a < b
	}
}
func (ss StringSorter) Less(i, j int) bool {
	// Выбираем что сравнивать
	var a, b string
	if ss.sortColumn == -1 {
		a = ss.strings[i]
		b = ss.strings[j]
	} else {
		aColumns := strings.Split(ss.strings[i], " ")
		bColumns := strings.Split(ss.strings[j], " ")

		if len(a) <= ss.sortColumn {
			return true
		} else if len(b) <= ss.sortColumn {
			return false
		}

		a = aColumns[ss.sortColumn]
		b = bColumns[ss.sortColumn]
	}
	// Сравниваем
	return ss.stringsLess(a, b)
}
func (ss StringSorter) Sort() {
	sort.Sort(ss)
}

func (ss StringSorter) Get(reversed bool, removeDuplicates bool) []string {
	result := make([]string, 0)

	// Определяем сторону обхода
	begin := 0
	end := 0
	step := 0
	if !reversed {
		begin = 0
		end = len(ss.strings)
		step = 1
	} else {
		begin = len(ss.strings) - 1
		end = -1
		step = -1
	}

	// Проверяем нужны ли дубликаты
	if !removeDuplicates {
		for ; begin != end; begin += step {
			result = append(result, ss.strings[begin])
		}
	} else {
		result = append(result, ss.strings[begin])
		begin += step

		for ; begin != end; begin += step {
			if ss.strings[begin-step] != ss.strings[begin] {
				result = append(result, ss.strings[begin])
			}
		}
	}

	return result
}
