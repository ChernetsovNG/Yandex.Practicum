package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

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
		uniqueWords := make(map[string]bool) // в запросе нужно рассматривать только уникальные слова
		words = strings.Split(line, " ")
		for j := 0; j < len(words); j++ {
			uniqueWords[words[j]] = true
		}

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
		i := 0
		for documentIndex, relevance := range documentsRelevance {
			relevanceArray[i] = kv{documentIndex + 1, relevance}
			i++
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
		printArray(relevanceArray)
	}
}

type kv struct {
	Key   int
	Value int
}

func printArray(array []kv) {
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
		fmt.Printf("%d", array[i].Key)
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", array[k-1].Key)
}
