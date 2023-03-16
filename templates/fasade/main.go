package main

import "fmt"

/*
	Шаблон Фассад

	Позволяет скрыть сложность системы, путём создания структуры агригирующей
	взаимодействие с множеством других модулей, упрощает построеение бизнес
	логики.
*/

type CodeEditor struct{}

func (t *CodeEditor) Save() {
	fmt.Println("Code saved")
}

type Compiller struct{}

func (t *Compiller) Compile() {
	fmt.Println("Code compiled")
}

type CLI struct{}

func (t *CLI) Run() {
	fmt.Println("Program run")
}

type IDEFacade struct {
	codeEditor *CodeEditor
	compiller  *Compiller
	cli        *CLI
}

func newIDEFacade() *IDEFacade {
	ideFacacde := &IDEFacade{
		codeEditor: &CodeEditor{},
		compiller:  &Compiller{},
		cli:        &CLI{},
	}
	return ideFacacde
}

func (ide *IDEFacade) BuildAndRun() {
	ide.codeEditor.Save()
	ide.compiller.Compile()
	ide.cli.Run()
}

func main() {
	ide := newIDEFacade()
	ide.BuildAndRun()
}
