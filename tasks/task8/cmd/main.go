package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task8/internal/shell"
	"task8/internal/shell/shellcommands"
)

// Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

// - cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
// - pwd - показать путь до текущего каталога
// - echo <args> - вывод аргумента в STDOUT
// - kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
// - ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

// Так же требуется поддерживать функционал fork/exec-команд

// Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

func main() {
	myShell := shell.NewShell()

	// Регистрируем команды
	myShell.RegisterComand("pwd", shellcommands.PWDCommand{})
	myShell.RegisterComand("cd", shellcommands.CDCommand{})
	myShell.RegisterComand("echo", shellcommands.EchoCommand{})
	myShell.RegisterComand("ps", shellcommands.PSCommand{})
	myShell.RegisterComand("kill", shellcommands.KillCommand{})
	myShell.RegisterComand("exec", shellcommands.ExecCommand{})
	myShell.RegisterComand("exit", shellcommands.ExitCommand{})

	// Обрабатываем ввод
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, "|") // Делим на команды
		for _, rawCommand := range commands {
			command := strings.Fields(rawCommand)       // Делим на аргументы
			myShell.ExecComand(command[0], command[1:]) // Вызывае команду
			fmt.Print(">")
		}
	}
}
