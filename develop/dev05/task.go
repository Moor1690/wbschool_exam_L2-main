package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"launchpad.net/gnuflag"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type mFlags struct {
	after     int  //-A - "after" печатать +N строк после совпадения
	before    int  //-B - "before" печатать +N строк до совпадения
	context   int  //-C - "context" (A+B) печатать ±N строк вокруг совпадения
	count     bool //-c - "count" (количество строк)
	gnoreCase bool //-i - "ignore-case" (игнорировать регистр)
	invert    bool //-v - "invert" (вместо совпадения, исключать)
	fixed     bool //-F - "fixed", точное совпадение со строкой, не паттерн
	lineNum   bool //-n - "line num", печатать номер строки
}

func main() {
	mf := mFlags{}
	lines := make([]string, 0)
	gnuflag.IntVar(&mf.after, "A", 0, `- "after" печатать +N строк после совпадения`)
	gnuflag.IntVar(&mf.before, "B", 0, `- "- "before" печатать +N строк до совпадения`)
	gnuflag.IntVar(&mf.context, "C", 0, `- "context" (A+B) печатать ±N строк вокруг совпадения`)
	gnuflag.BoolVar(&mf.count, "c", false, `- "count" (количество строк)`)
	gnuflag.BoolVar(&mf.gnoreCase, "i", false, `- ignore-case" (игнорировать регистр)`)
	gnuflag.BoolVar(&mf.invert, "v", false, `- "invert" (вместо совпадения, исключать)`)
	gnuflag.BoolVar(&mf.fixed, "F", false, `- "fixed", точное совпадение со строкой, не паттерн`)
	gnuflag.BoolVar(&mf.lineNum, "n", false, `- "line num", печатать номер строки`)
	gnuflag.Parse(true)

	if gnuflag.NArg() < 2 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN FILE")
		os.Exit(1)
	}

	pattern := gnuflag.Arg(0)
	filename := gnuflag.Arg(1)

	var re *regexp.Regexp
	var err error
	if mf.fixed {
		re, err = regexp.Compile(regexp.QuoteMeta(pattern))
	} else {
		if mf.gnoreCase {
			pattern = "(?i)" + pattern
		}
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		fmt.Println("Invalid pattern:", err)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	readLines(file, &lines)

	grep(mf, lines, re)

}
func readLines(f *os.File, lines *[]string) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
}

func grep(mf mFlags, lines []string, pattern *regexp.Regexp) []string {
	resultLines := []string{}
	count := 0
	lastBeforeIndex := -1

	for index := 0; index < len(lines); index++ {
		match := (pattern.MatchString(lines[index]) && !mf.invert) || (!pattern.MatchString(lines[index]) && mf.invert)
		if mf.count && match {
			count++
		} else if match {
			if mf.context > 0 {
				start := max(0, index-mf.context)
				if start <= lastBeforeIndex {
					start = lastBeforeIndex + 1
				}
				for i := start; i < min(len(lines), index+mf.context+1); i++ {
					resultLines = append(resultLines, addToResult(mf.lineNum, lines[i], i))
				}
				lastBeforeIndex = index + mf.context
			} else if mf.before > 0 {
				start := max(0, index-mf.before)
				if start <= lastBeforeIndex {
					start = lastBeforeIndex + 1
				}
				for i := start; i <= index; i++ {
					resultLines = append(resultLines, addToResult(mf.lineNum, lines[i], i))
				}
				lastBeforeIndex = index
			} else if mf.after > 0 {
				i := index
				for ; i < min(len(lines), index+mf.after+1); i++ {
					resultLines = append(resultLines, addToResult(mf.lineNum, lines[i], i))
				}
				index = i - 1
			} else {
				resultLines = append(resultLines, addToResult(mf.lineNum, lines[index], index))
			}

		}
	}

	if mf.count {
		resultLines = append(resultLines, strconv.Itoa(count))
	}
	for i := 0; i < len(resultLines); i++ {
		fmt.Println(resultLines[i])
	}

	return resultLines
}

func addToResult(lineNum bool, line string, index int) string {
	if lineNum {
		return strconv.Itoa(index+1) + ":" + line
	}
	return line
}
