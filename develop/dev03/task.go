package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"launchpad.net/gnuflag"
)

// ColumnSorter структура для сортировки по столбцу
type ColumnSorter struct {
	OriginalString string // Исходная строка
	SortColumn     string // Столбец для сортировки
}

// mySort структура, содержащая параметры для сортировки
type mySort struct {
	filename      string         // Имя файла для сортировки
	k             int            // Индекс колонки для сортировки
	n             bool           // Флаг числовой сортировки
	r             bool           // Флаг сортировки в обратном порядке
	u             bool           // Флаг для удаления дубликатов
	columnSorters []ColumnSorter // Срез структур ColumnSorter для сортировки
}

// newMySort создаёт новый экземпляр mySort
func newMySort() *mySort {
	return &mySort{}
}

func main() {
	// Инициализация структуры mySort
	ms := newMySort()

	// Парсинг флагов командной строки
	gnuflag.IntVar(&ms.k, "k", 0, "указание колонки для сортировки")
	gnuflag.BoolVar(&ms.n, "n", false, "сортировать по числовому значению")
	gnuflag.BoolVar(&ms.r, "r", false, "сортировать в обратном порядке")
	gnuflag.BoolVar(&ms.u, "u", false, "не выводить повторяющиеся строки")
	gnuflag.StringVar(&ms.filename, "f", "", "имя файла")
	gnuflag.Parse(true)

	// Чтение данных из файла или стандартного ввода
	if ms.filename == "" {
		readLines(os.Stdin, ms)
	} else {
		file, err := os.Open(ms.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		defer file.Close()
		readLines(file, ms)
	}

	// Сортировка строк
	sortStringsByColumn(ms)
}

// readLines читает строки из файла или стандартного ввода
func readLines(f *os.File, ms *mySort) {
	scanner := bufio.NewScanner(f)
	uniqueMap := make(map[string]bool)

	for scanner.Scan() {
		str := scanner.Text()
		columns := strings.Split(str, " ")
		sortColumn := str
		if ms.k < len(columns) {
			sortColumn = columns[ms.k]
		}

		if ms.u {
			if _, exists := uniqueMap[sortColumn]; exists {
				continue
			}
			uniqueMap[sortColumn] = true
		}

		ms.columnSorters = append(ms.columnSorters, ColumnSorter{OriginalString: str, SortColumn: sortColumn})
	}
}

// sortStringsByColumn сортирует строки с учётом заданных параметров
func sortStringsByColumn(ms *mySort) {
	// Определение функции сортировки
	sortFunc := func(i, j int) bool {
		if ms.n {
			ival, ierr := strconv.Atoi(ms.columnSorters[i].SortColumn)
			jval, jerr := strconv.Atoi(ms.columnSorters[j].SortColumn)
			if ierr == nil && jerr == nil {
				return ival < jval
			}
		}
		return ms.columnSorters[i].SortColumn < ms.columnSorters[j].SortColumn
	}

	// Применение сортировки
	if ms.r {
		sort.Slice(ms.columnSorters, func(i, j int) bool {
			return !sortFunc(i, j)
		})
	} else {
		sort.Slice(ms.columnSorters, sortFunc)
	}

	// Вывод отсортированных строк
	for _, cs := range ms.columnSorters {
		fmt.Println(cs.OriginalString)
	}
}
