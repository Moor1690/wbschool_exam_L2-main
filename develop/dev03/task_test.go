package main

import (
	"reflect"
	"testing"
)

// TestSortStringsByColumn тестирует функцию сортировки строк
func TestSortStringsByColumn(t *testing.T) {
	tests := []struct {
		name     string
		input    []ColumnSorter
		expected []ColumnSorter
		ms       *mySort
	}{
		{
			name: "Тест сортировки по строкам",
			input: []ColumnSorter{
				{"cherry", "cherry"},
				{"banana", "banana"},
				{"apple", "apple"},
			},
			expected: []ColumnSorter{
				{"apple", "apple"},
				{"banana", "banana"},
				{"cherry", "cherry"},
			},
			ms: &mySort{k: 0, n: false, r: false, u: false},
		},
		{
			name: "Тест сортировки по числам",
			input: []ColumnSorter{
				{"2", "2"},
				{"1", "1"},
				{"3", "3"},
			},
			expected: []ColumnSorter{
				{"1", "1"},
				{"2", "2"},
				{"3", "3"},
			},
			ms: &mySort{k: 0, n: true, r: false, u: false},
		},
		// Дополнительные тесты для других случаев
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ms.columnSorters = tt.input
			sortStringsByColumn(tt.ms)
			if !reflect.DeepEqual(tt.ms.columnSorters, tt.expected) {
				t.Errorf("sortStringsByColumn() = %v, want %v", tt.ms.columnSorters, tt.expected)
			}
		})
	}
}
