package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
https://contest.yandex.ru/contest/24414/run-report/53763938/

-- ПРИНЦИП РАБОТЫ --
Каждый документ мы разбиваем на слова, и для каждого слова сохраняем словарь, в котором ключом является
номер документа, а значением - количество раз, сколько данное слово встречается в документе с этим номером.

Далее, когда выполняется поисковый запрос, то мы выделяем в нём уникальные слова, и для каждого слова
находим, в каких документах и сколько раз оно встречалось. Далее выполняется суммирование по всем
словам запроса, и таким образом вычисляется релевантность всех документов. Полученный результат сортируется
по релевантности, и берутся первые 5 элементов массива.

Для оптимизации решения результат обработки каждого поискового запроса кешируется.

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
Мы сохраняем поисковый индекс размера O(w1 * f1). Дополнительно сохраняется кеш поисковых запросов, размер
которого зависит от количества повторяющихся запросов
*/
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
	searchCache := newSearchCache()

	for i := 0; i < m; i++ {
		// 1. считываем запрос
		scanner.Scan()
		line = scanner.Text()

		// если результат был закеширован, то используем его
		cachedResult, ok := searchCache.cacheElements[line]
		if ok {
			topFiveDocuments := cachedResult.topFiveDocuments
			printArray(writer, topFiveDocuments)
			continue
		}

		// 2. по индексу считаем релевантность всех документов
		uniqueWords := make(map[string]bool) // в запросе нужно рассматривать только уникальные слова
		words = strings.Split(line, " ")
		for j := 0; j < len(words); j++ {
			uniqueWords[words[j]] = true
		}

		// для каждого документа - его релевантность
		documentsRelevance := make(map[int]int)
		for word := range uniqueWords {
			documentsCounts, ok := searchIndex[word]
			if ok { // такое слово есть среди документов
				for documentIndex, count := range documentsCounts {
					documentsRelevance[documentIndex] += count
				}
			}
		}

		// 3. сортируем документы по релевантности
		relevanceArray := make([]kv, len(documentsRelevance))
		arrIndex := 0
		for documentIndex, relevance := range documentsRelevance {
			relevanceArray[arrIndex] = kv{documentIndex + 1, relevance}
			arrIndex++
		}

		sort.Slice(relevanceArray, func(i, j int) bool {
			rel1 := relevanceArray[i]
			rel2 := relevanceArray[j]
			if rel1.Value > rel2.Value {
				return true
			} else if rel1.Value < rel2.Value {
				return false
			} else { // релевантности совпадают => сортируем по возрастанию порядкового номер документа
				return rel1.Key < rel2.Key
			}
		})

		// выводим лучшие 5 документов
		printArrayOfPairs(writer, relevanceArray)

		// кешируем результат
		resultForCaching := make([]int, len(relevanceArray))
		for k := 0; k < len(relevanceArray); k++ {
			pair := relevanceArray[k]
			resultForCaching[k] = pair.Key
		}
		searchCache.cacheElements[line] = newCacheElement(resultForCaching)
	}

	writer.Flush()
}

type kv struct {
	Key   int
	Value int
}

// Для ускорения поиска кешируем результаты: для каждого запроса сохраняем топ-5 документов
type SearchCache struct {
	cacheElements map[string]CacheElement
}

func newSearchCache() SearchCache {
	return SearchCache{make(map[string]CacheElement)}
}

type CacheElement struct {
	topFiveDocuments []int
}

func newCacheElement(topFiveDocuments []int) CacheElement {
	return CacheElement{topFiveDocuments}
}

func printArray(writer *bufio.Writer, array []int) {
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
		writer.WriteString(strconv.Itoa(array[i]))
		writer.WriteString(" ")
	}
	writer.WriteString(strconv.Itoa(array[k-1]))
	writer.WriteString("\n")
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
