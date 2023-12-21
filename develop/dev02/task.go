package main

import (
	"errors"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	examples := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	for _, ex := range examples {
		unpacked, err := unpackString(ex)
		if err != nil {
			println("Ошибка:", err.Error())
		} else {
			println(unpacked)
		}
	}
}

func unpackString(str string) (string, error) {
	result := []rune{}
	runes := []rune(str)
	checkEscape := false
	for i, r := range runes {
		if checkEscape {
			result = append(result, r)
			checkEscape = false
			continue
		}
		if r == '\\' {
			checkEscape = true
			continue
		}
		if unicode.IsDigit(r) {

			if !checkEscape && i > 0 {
				run, _ := strconv.Atoi(string(runes[i]))
				for j := 0; j < run-1; j++ {
					result = append(result, runes[i-1])
				}
			} else {
				return "", errors.New("некорректная строка")
			}

		}

		if !unicode.IsDigit(r) {
			result = append(result, r)
		}
	}

	return string(result), nil
}
