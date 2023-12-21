package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"launchpad.net/gnuflag"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type mFlags struct {
	fields    []int
	delimiter string
	separated bool
}

func main() {
	mf := mFlags{}
	fieldFlag := gnuflag.String("f", "", `-f - "fields" - выбрать поля (колонки)`)
	gnuflag.StringVar(&mf.delimiter, "d", "\t", `-d - "delimiter" - использовать другой разделитель`)
	gnuflag.BoolVar(&mf.separated, "s", false, `-s - "separated" - только строки с разделителем`)
	gnuflag.Parse(true)

	if *fieldFlag == "" {
		fmt.Println("Error: Please provide the -f flag with the fields to extract")
		os.Exit(1)
	}

	var err error
	mf.fields, err = parseColumns(*fieldFlag)
	if err != nil {
		fmt.Printf("Ошибка при разборе флага -f: %v\n", err)
		os.Exit(1)
	}
	var result string

	if gnuflag.NArg() < 1 {
		f := os.Stdin
		result, err = mCut(f, mf)
	} else {
		filename := gnuflag.Arg(0)
		var f *os.File
		f, err = os.Open(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "err:\t", err)
			return
		}
		result, err = mCut(f, mf)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func mCut(f *os.File, mf mFlags) (string, error) {

	scanner := bufio.NewScanner(f)
	columns := []string{}
	rColumns := []string{}
	var result string

	for scanner.Scan() {
		str := scanner.Text()
		columns = strings.Split(str, mf.delimiter)
		if mf.separated && !strings.Contains(str, mf.delimiter) {
			continue
		}
		for i, _ := range mf.fields {
			if mf.fields[i] > len(columns) {
				errMsg := fmt.Sprintf("отсутствует колонка номер %d", mf.fields[i])
				return "", fmt.Errorf(errMsg)
			}
			rColumns = append(rColumns, columns[mf.fields[i]-1])
		}
		result += strings.Join(rColumns, mf.delimiter)
		rColumns = rColumns[:0]
		result += "\n"
	}
	return result, nil
}

func parseColumns(input string) ([]int, error) {
	var columns []int
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		if strings.Contains(r, "-") {
			bounds := strings.Split(r, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("неверный формат диапазона: %s", r)
			}

			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return nil, fmt.Errorf("неверный формат начала диапазона: %s", bounds[0])
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return nil, fmt.Errorf("неверный формат конца диапазона: %s", bounds[1])
			}

			for i := start; i <= end; i++ {
				columns = append(columns, i)
			}
		} else {
			col, err := strconv.Atoi(r)
			if err != nil {
				return nil, fmt.Errorf("неверный формат колонки: %s", r)
			}
			columns = append(columns, col)
		}
	}
	return columns, nil
}
