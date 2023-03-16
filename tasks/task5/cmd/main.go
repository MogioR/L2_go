package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"task5/internal/params"
	"task5/pkg/circularqueue"
)

// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

// Реализовать поддержку утилитой следующих ключей:
// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", напечатать номер строки

func printLine(line string, lineNum int, prtintLineNum bool) {
	if prtintLineNum {
		fmt.Println(strconv.Itoa(lineNum) + ":" + line)
	} else {
		fmt.Println(line)
	}
}

func main() {
	// Парсим параметры
	params := params.Params{}
	err := params.ParseArguments()
	if err != nil {
		log.Fatal("Can't parce arguments: " + string(err.Error()))
	}

	if flag.NArg() < 2 {
		log.Fatal("No pattern or file name found")
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	// Открываем файл для чтения
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Can't open file: " + string(err.Error()))
	}
	defer file.Close()

	//Пременные для вывода информации
	beforeLines := circularqueue.NewCircularQueue[string](params.Before) // Очередь предыдущих строк
	currentLineNum := 0
	afterLinesCount := 0
	matchedCount := 0

	// Компиляция регулярного выражения
	var re *regexp.Regexp
	if params.Fixed { // Если необходимо точное значение, преобразуем шаблон
		pattern = regexp.QuoteMeta(pattern)
	}
	if params.IgnoreCase {
		re, err = regexp.Compile("(?i)" + pattern) // Игнорирование регистра
	} else {
		re, err = regexp.Compile(pattern) // Обычное выражение
	}
	if err != nil {
		log.Fatal("Can't compile regexp: " + string(err.Error()))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()          // Читаем строку
		matched := re.MatchString(currentLine) // Есть ли совпадение в строке?

		// Если строка необходимая
		if (matched && !params.Invert) || (!matched && params.Invert) {
			// Если считаем строки
			if params.Counting {
				matchedCount += 1
				continue
			}

			// Выводим строки до
			for !beforeLines.IsEmpty() {
				printLine(beforeLines.Pull(), currentLineNum-beforeLines.Len()-1, params.LineNum)
			}

			// Выводим данную строку
			printLine(currentLine, currentLineNum, params.LineNum)

			// Сообщаем что нужно писать контекст
			afterLinesCount = params.After
		} else {
			if afterLinesCount <= 0 { // Сохраняем стороку на случай нахождения
				beforeLines.Push(currentLine)
			} else { // Или печатаем контекст если необходимо
				afterLinesCount -= 1
				printLine(currentLine, currentLineNum, params.LineNum)
			}
		}
		currentLineNum += 1
	}

	// Вывод количества строк
	if params.Counting {
		fmt.Println(matchedCount)
	}
}
