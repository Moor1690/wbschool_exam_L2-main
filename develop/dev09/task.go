/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"launchpad.net/gnuflag"
)

func main() {
	var url string
	gnuflag.Parse(true)
	if gnuflag.NArg() != 1 {
		fmt.Println("invalid argument")
	} else {
		url = gnuflag.Arg(0)
	}

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	outFile, err := os.Create("downloaded_page.html")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		panic(err)
	}

	println("Скачивание завершено.")
}
