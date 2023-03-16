package main

import "fmt"

// Паттерн Visitor
// определяет операцию, выполняемую на каждом элементе из некоторой структуры.
// Позволяет, не изменяя классы этих объектов, добавлять в них новые операции.

// Интерфейс посетителя
type Visitor interface {
	VisitHubaBuba(element *HubaBuba)
	VisitTakSak(element *TakSak)
}

// Конкретный посетитель 1
type JSONEncoder struct{}

func (cv1 *JSONEncoder) VisitHubaBuba(element *HubaBuba) {
	fmt.Println("JSON HubaBuba")
}

func (cv1 *JSONEncoder) VisitTakSak(element *TakSak) {
	fmt.Println("JSON TakSak")
}

// Конкретный посетитель 2
type XMLEncoder struct{}

func (cv2 *XMLEncoder) VisitHubaBuba(element *HubaBuba) {
	fmt.Println("XML HubaBuba")
}

func (cv2 *XMLEncoder) VisitTakSak(element *TakSak) {
	fmt.Println("XML TakSak")
}

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type HubaBuba struct{}

func (ea *HubaBuba) Accept(visitor Visitor) {
	visitor.VisitHubaBuba(ea)
}

// Конкретный элемент B
type TakSak struct{}

func (eb *TakSak) Accept(visitor Visitor) {
	visitor.VisitTakSak(eb)
}

// Клиентский код
func main() {
	// Создание объектов элементов
	hubaBuba := &HubaBuba{}
	elementB := &TakSak{}

	// Создание объектов посетителей
	jsonEncoder := &JSONEncoder{}
	xmlEncoder := &XMLEncoder{}

	// Применение посетителя к элементам
	hubaBuba.Accept(jsonEncoder)
	elementB.Accept(xmlEncoder)

	hubaBuba.Accept(jsonEncoder)
	elementB.Accept(xmlEncoder)
}
