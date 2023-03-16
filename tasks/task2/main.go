package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
	Создать Go-функцию, осуществляющую примитивную распаковку строки,
	содержащую повторяющиеся символы/руны, например:
	"a4bc2d5e" => "aaaabccddddde"
	"abcd" => "abcd"
	"45" => "" (некорректная строка)
	"" => ""

	Дополнительно
	Реализовать поддержку escape-последовательностей.
	Например:
	qwe\4\5 => qwe45 (*)
	qwe\45 => qwe44444 (*)
	qwe\\5 => qwe\\\\\ (*)


	В случае если была передана некорректная строка, функция должна
	возвращать ошибку. Написать unit-тесты.
*/

// Функция чтения числа с позиции i, i перемнщается на следующий символ после числа
func readNum(runes []rune, i *int) int {
	n := 0
	// Читаем первую цифру
	buf, _ := strconv.Atoi(string(runes[*i]))
	n += buf
	*i += 1
	// Читаем оставшиеся
	for ; *i < len(runes) && unicode.IsDigit(runes[*i]); *i++ {
		buf, _ := strconv.Atoi(string(runes[*i]))
		n = n*10 + buf
	}

	return n
}

// Функция распаковки
func unpack(str string) string {
	resultBuilder := strings.Builder{}

	// Прохидимся по рунам
	runes := []rune(str)
	for i := 0; i < len(runes); {
		// Если находим цифру до буквы это ошибка
		if unicode.IsDigit(runes[i]) {
			return ""
		}
		// Если находим слеш, пропускаем его
		if runes[i] == '\\' {
			i += 1
			// Если в конце одинарный слеш, это ошибка
			if i >= len(runes) {
				return ""
			}
		}

		// Сохраняем символ
		selectedRune := runes[i]
		resultBuilder.WriteRune(runes[i])
		i += 1
		// Если после идет число
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			// Читаем число
			n := readNum(runes, &i)

			// Добовляем предыдущий символ n-1 раз
			for j := 0; j < n-1; j++ {
				resultBuilder.WriteRune(selectedRune)
			}
		}
	}
	return resultBuilder.String()
}

func main() {
	fmt.Println(unpack(`a\`))
	fmt.Println(unpack(`a4bc2d5e`))
	fmt.Println(unpack(`abcd`))
	fmt.Println(unpack(`a11b`))
	fmt.Println(unpack(`a12`))
	fmt.Println(unpack(`45`))
	fmt.Println(unpack(``))

	fmt.Println(unpack(`qwe\4\5`))
	fmt.Println(unpack(`qwe\45`))
	fmt.Println(unpack(`qwe\\5`))
	fmt.Println(unpack(`qwe\\`))
}
