package main

import "fmt"

// Шаблон проектирования "Цепочка вызовов" (Chain of Responsibility)

// Позволяет построить цепочку объектов-обработчиков,
// где каждый объект может обработать запрос или передать
// его следующему объекту в цепочке.

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	nextHandler Handler
}

func (bh *BaseHandler) SetNext(handler Handler) Handler {
	bh.nextHandler = handler
	return handler
}

func (bh *BaseHandler) Handle(request string) {
	if bh.nextHandler != nil {
		bh.nextHandler.Handle(request)
	}
}

// Конкретные реализации обработчиков
type FirstHandler struct {
	BaseHandler
}

func (fh *FirstHandler) Handle(request string) {
	if request == "first" {
		fmt.Println("FirstHandler handled the request")
		return
	}
	fh.BaseHandler.Handle(request)
}

type SecondHandler struct {
	BaseHandler
}

func (sh *SecondHandler) Handle(request string) {
	if request == "second" {
		fmt.Println("SecondHandler handled the request")
		return
	}
	sh.BaseHandler.Handle(request)
}

type ThirdHandler struct {
	BaseHandler
}

func (th *ThirdHandler) Handle(request string) {
	if request == "third" {
		fmt.Println("ThirdHandler handled the request")
		return
	}
	th.BaseHandler.Handle(request)
}

// Клиентский код
func main() {
	// Создание объектов обработчиков
	firstHandler := &FirstHandler{}
	secondHandler := &SecondHandler{}
	thirdHandler := &ThirdHandler{}

	// Установка цепочки вызовов
	firstHandler.SetNext(secondHandler).SetNext(thirdHandler)

	// Выполнение запросов
	firstHandler.Handle("second")
	firstHandler.Handle("third")
	firstHandler.Handle("fourth")
}
