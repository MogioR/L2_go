package main

import "fmt"

// Паттерн "Состояние" (State)
// Позволяет объекту изменять свое поведение в зависимости от
// внутреннего состояния.

// С его помощью возможно инкапсулировать поведение, связанное
// с определенным состоянием, и делает его взаимозаменяемым.

// Интерфейс состояния
type State interface {
	Handle(context *Context)
}

// Конкретная реализация состояния #1
type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle(context *Context) {
	fmt.Print("Хуба")
	context.SetState(&ConcreteStateB{})
}

// Конкретная реализация состояния #2
type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Буба")
	context.SetState(&ConcreteStateA{})
}

// Контекст
type Context struct {
	state State
}

func NewContext(state State) *Context {
	return &Context{state}
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

// Клиентский код
func main() {
	// Создание объектов состояний
	stateA := &ConcreteStateA{}

	// Создание объекта контекста с начальным состоянием
	context := NewContext(stateA)

	// Использование контекста для выполнения действий в зависимости от состояния
	context.Request()
	context.Request()
	context.Request()
}
