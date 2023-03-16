package main

import (
	"fmt"
	"log"
	"os"
	"task3/internal/params"
	"task3/pkg/myfileutils"
	"task3/pkg/stringsorter"
)

/*
	Отсортировать строки в файле по аналогии с консольной утилитой sort
	(man sort — смотрим описание и основные параметры): на входе подается
	файл из несортированными строками, на выходе — файл с отсортированными.

	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать
		 в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

func main() {
	// Обрабатываем параметры запуска
	params := params.Params{}
	params.ParseArguments()

	// Читаем файл
	fileStrings, err := myfileutils.ReadFileStrings(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("Can't read file")
	}

	// Сортируем
	stringSorter := stringsorter.NewStringSorter(
		fileStrings,
		params.SortColumn,
		params.SortByNums,
	)
	stringSorter.Sort()

	//Выводим отсортированный файл на экран (Как делает это linux sort)
	fileStrings = stringSorter.Get(
		params.SortReverse,
		params.RemoveDuplicates,
	)
	for i := 0; i < len(fileStrings); i++ {
		fmt.Println(fileStrings[i])
	}

	// Можно и писать в файл
	// stringsorter.WriteStringsToFile("output.txt", fileStrings)
}
