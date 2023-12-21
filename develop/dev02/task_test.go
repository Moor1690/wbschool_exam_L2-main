package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("некорректная строка")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}

	for _, tc := range testCases {
		result, err := unpackString(tc.input)
		if tc.err != nil && err == nil {
			t.Errorf("Expected error for input %v", tc.input)
		} else if err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Unexpected error for input %v: got %v want %v", tc.input, err, tc.err)
		} else if result != tc.expected {
			t.Errorf("For input %v, expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}
