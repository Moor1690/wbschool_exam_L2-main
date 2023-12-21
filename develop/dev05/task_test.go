package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Создание временного файла
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Очистка

	content := "line1\nline2\nline3"
	if _, err = tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err = tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Повторное открытие файла для чтения
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	readLines(file, &lines)

	expectedLines := 3
	if len(lines) != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, len(lines))
	}

	// Проверка содержимого строк
	for i, line := range lines {
		expectedLine := fmt.Sprintf("line%d", i+1)
		if line != expectedLine {
			t.Errorf("Expected %s, got %s", expectedLine, line)
		}
	}
}

func TestGrepCount(t *testing.T) {
	lines := []string{"test line", "another line", "test"}
	pattern := regexp.MustCompile("test")
	mf := mFlags{count: true}

	result := grep(mf, lines, pattern)

	if len(result) != 1 || result[0] != "2" {
		t.Errorf("Expected count 2, got %v", result)
	}
}
func TestGrepIgnoreCaseInvert(t *testing.T) {
	lines := []string{"Test", "test", "another"}
	pattern := regexp.MustCompile("(?i)test")
	mf := mFlags{gnoreCase: true, invert: true}

	result := grep(mf, lines, pattern)

	if len(result) != 1 || result[0] != "another" {
		t.Errorf("Expected 1 non-matching line, got %v", result)
	}
}
func TestGrepEmptyFile(t *testing.T) {
	lines := []string{}
	pattern := regexp.MustCompile("test")
	mf := mFlags{}

	result := grep(mf, lines, pattern)

	if len(result) != 0 {
		t.Errorf("Expected no output for empty file, got %v", result)
	}
}
