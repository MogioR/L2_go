package main

import "fmt"

// Шаблон проектирования "Фабричный метод" (Factory Method) позволяет определить интерфейс
// для создания объектов, но позволяет подклассам решать, какой класс создавать. То есть,
// вместо того чтобы явно создавать объекты в коде, мы создаем фабричный метод, который
// создает объекты для нас.

// Интерфейс продукта
type Product interface {
	Use()
}

// Конкретная реализация продукта
type ConcreteProduct struct{}

func (cp *ConcreteProduct) Use() {
	fmt.Println("Using ConcreteProduct")
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретная реализация фабрики
type ConcreteFactory struct{}

func (cf *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{}
}

// Клиентский код
func main() {
	// Создание объекта фабрики
	factory := &ConcreteFactory{}

	// Создание продукта с помощью фабричного метода
	product := factory.CreateProduct()

	// Использование продукта
	product.Use()
}
