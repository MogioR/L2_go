package main

import (
	"fmt"
	"strings"
)

// Шаблон проектирования "Команда" (Command) позволяет инкапсулировать запросы
// или операции как объекты, что позволяет параметризовать клиентов с различными
// запросами, очередями или журналами запросов, а также поддерживать отмену
//  операций и их повторное выполнение.

// Интерфейс команды
type Command interface {
	Execute()
	Cancel()
}

// Конкретная команда 1
type PasteHubaBuba struct {
	receiver *Notepad
}

func (cc1 *PasteHubaBuba) Execute() {
	cc1.receiver.textBox += "HubaBuba"
}

func (cc1 *PasteHubaBuba) Cancel() {
	cc1.receiver.textBox = strings.ReplaceAll(cc1.receiver.textBox, "HubaBuba", "")
}

// Конкретная команда 2
type Close struct {
	receiver *Notepad
}

func (cc2 *Close) Execute() {
	cc2.receiver.Close()
}

func (cc1 *Close) Cancel() {
	cc1.receiver.Open()
}

// Получатель команды
type Notepad struct {
	textBox string
}

func (r *Notepad) Open() {
	fmt.Println("notepad open")
}

func (r *Notepad) Close() {
	fmt.Println("notepad close")
}

// Инициатор команды
type Invoker struct {
	command Command
	history []Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
	i.history = append(i.history, i.command)
}

func (i *Invoker) UndoCommand() {
	if len(i.history) > 0 {
		i.history[len(i.history)-1].Cancel()
		i.history = i.history[:len(i.history)-1]
	}
}

// Клиентский код
func main() {
	// Создание объектов команд и получателя
	receiver := &Notepad{
		"",
	}
	paste := &PasteHubaBuba{receiver: receiver}
	close := &Close{receiver: receiver}

	// Создание объекта инициатора команды
	invoker := &Invoker{
		history: make([]Command, 0),
	}

	// Установка команды и выполнение
	invoker.SetCommand(paste)
	invoker.ExecuteCommand()

	fmt.Println(receiver.textBox)

	// Установка другой команды и выполнение
	invoker.SetCommand(close)
	invoker.ExecuteCommand()

	invoker.UndoCommand()
	invoker.UndoCommand()

	fmt.Println(receiver.textBox)

}
