package main

/*
https://contest.yandex.ru/contest/24414/run-report/53771527/

-- ПРИНЦИП РАБОТЫ --
Каждый документ мы разбиваем на слова, и для каждого слова сохраняем словарь, в котором ключом является
номер документа, а значением - количество раз, сколько данное слово встречается в документе с этим номером.

Далее, когда выполняется поисковый запрос, то мы выделяем в нём уникальные слова, и для каждого слова
находим, в каких документах и сколько раз оно встречалось. Далее выполняется суммирование по всем
словам запроса, и таким образом вычисляется релевантность всех документов. Полученный результат сортируется
по релевантности, и берутся первые 5 элементов массива.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Решение соответствует условию задания в части определения релевантности документа

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Пусть n - количество документов, а w1 - среднее количество слов в каждом документе. Пусть каждое слово из
документов в среднем содержится в f1 документах.
Тогда построение поискового индекса выполняется за время O(n * w).

Пусть m - количество запросов, а w2 - среднее количество слов в каждом запросе. Пусть каждое слово поискового
запроса в среднем содержится в f2 документах из индекса.
Тогда получение топ-5 документов по релевантности выполняется за время O(m * (w2 + w2*f2 + f2*log(f2)))

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Мы сохраняем поисковый индекс размера O(w1 * f1).
*/

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(os.Stdout)

	var line string

	// читаем количество документов в базе
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем документы и сразу же рассчитываем "поисковый индекс"
	// для каждого слова сохраняем индексы документов, в которых оно встречается, и сколько раз
	// ключ - слово; значение - словарь "номер документа - количество раз"
	searchIndex := make(map[string]map[int]int)
	var words []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		words = strings.Split(line, " ")
		for j := 0; j < len(words); j++ {
			documentsCounts, ok := searchIndex[words[j]]
			if !ok {
				documentsCountsMap := make(map[int]int)
				documentsCountsMap[i] = 1
				searchIndex[words[j]] = documentsCountsMap
			} else {
				documentsCounts[i] += 1
			}
		}
	}

	// читаем число запросов
	var m int

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем и обрабатываем запросы
	for i := 0; i < m; i++ {
		// 1. считываем запрос
		scanner.Scan()
		line = scanner.Text()

		// 2. по индексу считаем релевантность всех документов
		// в запросе нужно рассматривать только уникальные слова
		uniqueWords := getUniqueWords(strings.Split(line, " "))

		// для каждого документа - его релевантность
		documentsRelevance := make([]int, n)
		for j := 0; j < len(uniqueWords); j++ {
			documentsCounts, ok := searchIndex[uniqueWords[j]]
			if ok { // такое слово есть среди документов
				for documentIndex, count := range documentsCounts {
					documentsRelevance[documentIndex] += count
				}
			}
		}

		// 3. сортируем документы по релевантности
		var relevanceArray []kv
		for documentIndex, relevance := range documentsRelevance {
			if relevance > 0 {
				relevanceArray = append(relevanceArray, kv{documentIndex + 1, relevance})
			}
		}

		sort.Slice(relevanceArray, func(i, j int) bool {
			rel1 := relevanceArray[i]
			rel2 := relevanceArray[j]
			if rel1.Value == rel2.Value { // релевантности совпадают => сортируем по возрастанию порядкового номер документа
				return rel1.Key < rel2.Key
			} else {
				return rel1.Value > rel2.Value
			}
		})

		// выводим лучшие 5 документов
		printArrayOfPairs(writer, relevanceArray)
	}

	writer.Flush()
}

func getUniqueWords(words []string) []string {
	uniqueWords := make(map[string]bool)
	for j := 0; j < len(words); j++ {
		uniqueWords[words[j]] = true
	}
	result := make([]string, len(uniqueWords))
	index := 0
	for word := range uniqueWords {
		result[index] = word
		index++
	}
	return result
}

type kv struct {
	Key   int
	Value int
}

func printArrayOfPairs(writer *bufio.Writer, array []kv) {
	if len(array) == 0 {
		return
	}
	var k int
	if len(array) >= 5 {
		k = 5
	} else {
		k = len(array)
	}
	for i := 0; i < k-1; i++ {
		writer.WriteString(strconv.Itoa(array[i].Key))
		writer.WriteString(" ")
	}
	writer.WriteString(strconv.Itoa(array[k-1].Key))
	writer.WriteString("\n")
}
