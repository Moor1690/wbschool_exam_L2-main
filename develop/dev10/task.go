/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		return
	}

	hostPort := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	conn, err := net.DialTimeout("tcp", hostPort, *timeout)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println(hostPort)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			_, err := conn.Write([]byte(text + "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		conn.Close()
	}()

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buff[:n]))
	}

	fmt.Println("Disconnected")
}
