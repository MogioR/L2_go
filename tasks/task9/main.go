package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

func main() {
	// Обработка флагов
	downloadAllPagePtr := flag.Bool("p", false, "Скачать всю страницу")
	flag.Parse()
	urlString := flag.Arg(0)

	url, err := url.Parse(urlString)
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}

	// Скачиваем страницу
	err = downloadFile(url, true)
	if err != nil {
		log.Fatal("Error downloading page:", err)
	}

	// И ресурсы страницы, при необходимости
	if *downloadAllPagePtr {
		err = downloadResources(url, url.Hostname()+"/index.html")
		if err != nil {
			log.Fatal("Error downloading resources:", err)
		}
	}
}

func downloadResources(baseURL *url.URL, fileName string) error {
	// Открываем файл скачанной страницы
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Считываем содержимое файла построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ищем ссылки на ресурсы
		relativeURLs := getRelativeURLsFromLine(line)
		// Скачиваем ресурсы
		for _, relativeURL := range relativeURLs {
			absoluteURL, _ := relativeURLToAbsolute(baseURL, relativeURL)

			err = downloadFile(absoluteURL, false)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

func downloadFile(url *url.URL, isPage bool) error {
	// Создаем HTTP-клиента
	client := http.Client{}

	// Отправляем GET-запрос и получаем ответ
	response, err := client.Get(url.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Создаем директорию
	var dirName string
	if isPage {
		dirName = url.Hostname()
		err = os.Mkdir(dirName, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			return err
		}
	} else {
		dirName = url.Hostname() + filepath.Dir(url.Path)
		err = os.MkdirAll(dirName, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}
	// Создаем файл, в который будем сохранять файл
	var fileName string
	if isPage {
		fileName = path.Join(dirName, "index.html")
	} else {
		fileName = path.Join(dirName, filepath.Base(url.Path))
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func getRelativeURLsFromLine(line string) []string {
	relativeURLRegexp := regexp.MustCompile("\"\\/([\\w-]*\\/)*([\\w-]*\\.)+\\w*\"")
	all := relativeURLRegexp.FindAllString(line, -1)

	result := make([]string, len(all))
	for i, a := range all {
		result[i] = strings.ReplaceAll(a, "\"", "")
	}

	return result
}

func relativeURLToAbsolute(baseURL *url.URL, relativeUrl string) (*url.URL, error) {
	return url.Parse(baseURL.Scheme + "://" + baseURL.Host + relativeUrl)
}
