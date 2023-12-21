// /*
// === Взаимодействие с ОС ===

// Необходимо реализовать собственный шелл

// встроенные команды: cd/pwd/echo/kill/ps
// поддержать fork/exec команды
// конвеер на пайпах

// Реализовать утилиту netcat (nc) клиент
// принимать данные из stdin и отправлять в соединение (tcp/udp)
// Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
// */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Получение текущей директории
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Ошибка при получении текущей директории: %v\n", err)
			continue
		}

		// Вывод текущей директории и приглашения к вводу
		fmt.Printf("%s> ", cwd)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("cd: expected target directory")
				continue
			}
			if err := os.Chdir(args[1]); err != nil {
				fmt.Println("cd:", err)
			}
		case "pwd":
			if dir, err := os.Getwd(); err != nil {
				fmt.Println("pwd:", err)
			} else {
				fmt.Println(dir)
			}
		case "ls":
			files, err := os.ReadDir(cwd)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("kill: expected process ID")
				continue
			}
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("kill: invalid process ID: %s\n", args[1])
				continue
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Printf("kill: process not found: %d\n", pid)
				continue
			}
			if err := process.Kill(); err != nil {
				fmt.Printf("kill: failed to kill process: %d\n", pid)
			} else {
				fmt.Printf("kill: process killed: %d\n", pid)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("ps: ошибка: %v\n", err)
				continue
			}
			fmt.Print(string(output))
		case `\quit`:
			return
		default:
			// Выполнение внешней команды
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Ошибка при выполнении команды: %v\n", err)
			}
		}
	}
}
