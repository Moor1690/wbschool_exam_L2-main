package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "стандартный тест",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "пустой ввод",
			words:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "с одним элементом в группе",
			words:    []string{"слово", "слово1", "слово2"},
			expected: map[string][]string{
				// Ожидается, что группы с одним словом не попадут в результат
			},
		},
		// Дополнительные тестовые случаи
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.words)

			// Сортировка результатов для стабильного сравнения
			for _, v := range result {
				sort.Strings(v)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagrams(%v) = %v, want %v", tt.words, result, tt.expected)
			}
		})
	}
}
