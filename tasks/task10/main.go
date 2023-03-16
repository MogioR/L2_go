package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализовать простейший telnet-клиент.

// Примеры вызовов:
// go-telnet --timeout=10s host port go-telnet mysite.ru 8080
// go-telnet --timeout=3s 1.1.1.1 123

// Требования:
// 1. 	Программа должна подключаться к указанному хосту (ip или доменное
// 		имя + порт) по протоколу TCP. После подключения STDIN программы
// 		должен записываться в сокет, а данные полученные и сокета должны
// 		выводиться в STDOUT

// 2.	Опционально в программу можно передать таймаут на подключение

// 		к серверу (через аргумент --timeout, по умолчанию 10s)
// 3.	При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
// 		Если сокет закрывается со стороны сервера, программа должна также
// 		завершаться. При подключении к несуществующему сервер, программа
// 		должна завершаться через timeout

func main() {
	// Обрабатываем аргументы запуска
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Use: go-telnet [--timeout timeout] host port")
	}
	host := args[0]
	port := args[1]

	// Производим подключение
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	// Горутина обработки принятых сообщений
	go func() {
		buf := make([]byte, 1024) // Создаём буфер
		for {
			n, err := conn.Read(buf) // Записываем в буфер входящие сообщения
			if err != nil {          // В случае ошибки вызодим
				fmt.Println("Connection closed:", err)
				os.Exit(0)
			}
			fmt.Fprint(os.Stdout, string(buf[:n])) // Пишем в Stdout
		}
	}()

	// Обработка закрытия
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func(chan os.Signal) {
		<-sig
		fmt.Println("Close conection...")
		conn.Close()
		os.Exit(0)
	}(sig)

	// Обрабатываем действия с клиента
	buf := make([]byte, 1024) // Буфер для отправки
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal("Error reading from STDIN:", err)
		}
		if n == 0 { // Ctrl+D отправляет EOF в Stdin
			fmt.Println("Close conection...")
			conn.Close()
			os.Exit(0)
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Fatal("Error writing to server:", err)
		}
	}
}
