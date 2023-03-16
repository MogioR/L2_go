package main

import "fmt"

// Паттерн "Стратегия" (Strategy)
// позволяет определить набор алгоритмов, инкапсулировать их и
// сделать их взаимозаменяемыми.

// Интерфейс стратегии
type Strategy interface {
	Execute()
}

// Конкретная реализация стратегии #1
type MagicSortOne struct{}

func (csa *MagicSortOne) Execute() {
	fmt.Println("Magic sort one")
}

// Конкретная реализация стратегии #2
type MagicSortTwo struct{}

func (csb *MagicSortTwo) Execute() {
	fmt.Println("Magic sort two")
}

// Контекст
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

// Клиентский код
func main() {
	// Создание объектов стратегий
	magicSortA := &MagicSortOne{}
	magicSortB := &MagicSortTwo{}

	// Создание объектов контекста с разными стратегиями
	contextA := NewContext(magicSortA)
	contextB := NewContext(magicSortB)

	// Использование контекста с разными стратегиями
	contextA.ExecuteStrategy()
	contextB.ExecuteStrategy()
}
