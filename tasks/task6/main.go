package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

func main() {
	// Парсим флаги
	fieldsPtr := flag.String("f", "", "выбрать поля (колонки)")
	delimiterPtr := flag.String("d", "\t", "использовать другой разделитель")
	separatedPtr := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	//Получаем значения флагов
	fields := *fieldsPtr
	delimiter := *delimiterPtr
	separated := *separatedPtr

	if fields == "" {
		log.Fatal("-f must be defined for work")
	}

	columns, err := parseFields(fields)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}

	// Считываем строки из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, delimiter) // Делим на колонки

		// Если нет колонок и только строки с разделителем выходим
		if len(columns) == 1 && separated {
			continue
		}

		//Проходимся по ограничениям
		var output []string
		for i, field := range fields {
			if columns[i] {
				output = append(output, field)
			}
		}

		fmt.Println(strings.Join(output, delimiter))
	}
	fmt.Println("weqarsdfsdf")
}

// Функеция без учета интервальных значений
func parseFields(fieldsStr string) (map[int]bool, error) {
	fields := strings.Split(fieldsStr, ",")
	res := make(map[int]bool, len(fields))
	for _, intStr := range fields {
		num, err := strconv.Atoi(intStr)
		if err != nil || num < 0 {
			return nil, errors.New("bad field list")
		}
		res[num] = true
	}
	return res, nil
}
