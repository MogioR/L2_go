package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			[]string{"abc", "cba", "acb", "bac", "cab", "bca"},
			map[string][]string{
				"abc": {"abc", "acb", "bac", "bca", "cab", "cba"},
			},
		},
		{
			[]string{"молоко", "колом", "ломик", "около", "лимон", "милок"},
			map[string][]string{
				"иклмо": {"ломик", "милок"},
			},
		},
		{
			[]string{""},
			map[string][]string{},
		},
	}

	for _, test := range tests {
		output := findAnagramSets(test.input)

		// проверяем, что полученный результат соответствует ожидаемому
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("findAnagramSets(%v) = %v; expected %v", test.input, output, test.expected)
		}

		// проверяем, что все массивы значений отсортированы по возрастанию
		for _, value := range output {
			if !sort.StringsAreSorted(value) {
				t.Errorf("Values in the anagram set %v are not sorted", value)
			}
		}
	}
}
