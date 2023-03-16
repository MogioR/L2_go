package myfileutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileStrings(fileName string) (result []string, err error) {
	file, err := os.Open(fileName) // Открываем файл

	if err != nil {
		log.Fatal("Can't open file: " + fileName)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Can't close file")
		}
	}(file)

	// Читаем построчно
	result = make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return result, sc.Err()
}

func WriteStringsToFile(fileName string, strings []string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Can't create file")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Can't close file")
		}
	}(file)

	for i := 0; i < len(strings); i++ {
		_, err := fmt.Fprintln(file, strings[i])
		if err != nil {
			log.Fatal("Can't write to file")
		}
	}
}
