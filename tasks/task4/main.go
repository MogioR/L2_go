package main

import (
	"fmt"
	"sort"
	"strings"
)

// Написать функцию поиска всех множеств анаграмм по словарю.

// Например:
// 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
// 'листок', 'слиток' и 'столик' - другому.

// Требования:
// * Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
// * Выходные данные: ссылка на мапу множеств анаграмм
// * Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
//   слово из множества.
// * Массив должен быть отсортирован по возрастанию.
// * Множества из одного элемента не должны попасть в результат.
// * Все слова должны быть приведены к нижнему регистру.
// * В результате каждое слово должно встречаться только один раз.

func sortWord(word string) string {
	// Переводим сторку в руны и сортируем
	sortedWord := []rune(word)
	sort.Slice(sortedWord, func(i, j int) bool {
		return sortedWord[i] < sortedWord[j]
	})
	// Возвращаем строкой
	return string(sortedWord)
}

func findAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к общему виду
		lowerWord := strings.ToLower(word)
		sortedWord := sortWord(lowerWord)

		// Проверяем, есть ли уже множество для данного слова
		if _, ok := anagramSets[sortedWord]; !ok {
			// Если множества еще нет, создаем новый массив и добавляем текущее слово в нижнем регистре
			anagramSets[sortedWord] = []string{lowerWord}
		} else {
			// Если множество уже есть, добавляем текущее слово в нижнем регистре
			anagramSets[sortedWord] = append(anagramSets[sortedWord], lowerWord)
		}
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		} else {
			// Сортируем массив слов в множестве по возрастанию
			sort.Strings(value)
			anagramSets[key] = value
		}
	}
	return anagramSets
}

func main() {
	wordsArr := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слитока"}
	anagramSets := findAnagramSets(wordsArr)
	fmt.Println(anagramSets)
}
