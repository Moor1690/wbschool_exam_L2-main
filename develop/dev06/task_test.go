package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseColumns(t *testing.T) {
	testCases := []struct {
		input    string
		expected []int
		hasError bool
	}{
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-3", []int{1, 2, 3}, false},
		{"1,2-4", []int{1, 2, 3, 4}, false},
		{"a-3", nil, true},
		{"1-", nil, true},
		{"-3", nil, true},
	}

	for _, tc := range testCases {
		output, err := parseColumns(tc.input)
		if (err != nil) != tc.hasError {
			t.Errorf("parseColumns(%s) unexpected error status: %v", tc.input, err)
		}
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("parseColumns(%s) = %v, expected %v", tc.input, output, tc.expected)
		}
	}
}

func TestMCut(t *testing.T) {
	// Создание временного файла для тестирования
	content := "a\tb\tc\nd\te\tf\n"
	tmpfile, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // очищаем после завершения теста

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name     string
		mf       mFlags
		expected string
		hasError bool
	}{
		{"TestOnlyFirstColumn", mFlags{fields: []int{1}, delimiter: "\t", separated: false}, "a\nd\n", false},
		{"TestSecondColumn", mFlags{fields: []int{2}, delimiter: "\t", separated: false}, "b\ne\n", false},
		{"TestNonExistentColumn", mFlags{fields: []int{4}, delimiter: "\t", separated: false}, "", true},
		{"TestSeparatedFlag", mFlags{fields: []int{1}, delimiter: "\t", separated: true}, "a\nd\n", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			result, err := mCut(f, tc.mf)
			if (err != nil) != tc.hasError {
				t.Errorf("unexpected error status: %v", err)
			}
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}
