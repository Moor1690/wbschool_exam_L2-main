package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	w := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	a := findAnagrams(w)
	fmt.Println(a)
}

func findAnagrams(w []string) map[string][]string {

	anagrams := make(map[string][]string)

	for i, val := range w {
		w[i] = strings.ToLower(w[i])
		sl := []rune{}
		for _, char := range val {
			sl = append(sl, char)
		}

		sort.Slice(sl, func(i, j int) bool {
			return sl[i] < sl[j]
		})

		if _, exists := anagrams[string(sl)]; !exists {
			anagrams[string(sl)] = []string{}
		}
		anagrams[string(sl)] = append(anagrams[string(sl)], w[i])

	}
	anagramsResult := make(map[string][]string)
	for i, val := range anagrams {
		if len(val) < 2 {
			delete(anagrams, i)
		} else {
			sort.Slice(val, func(i, j int) bool {
				return val[i] < val[j]
			})
			anagramsResult[anagrams[i][0]] = anagrams[i]
		}
	}
	return anagramsResult
}
