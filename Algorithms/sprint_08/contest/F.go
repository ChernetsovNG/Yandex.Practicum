package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем число строк
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем строки
	counter := make(map[string]int)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		count, contains := counter[line]
		if !contains {
			counter[line] = 1
		} else {
			counter[line] = count + 1
		}
	}

	// находим самое часто встречающееся слово
	var ss []kv
	for k, v := range counter {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	// находим слова, которые встречаются одинаково часто
	var mostFrequencyWords []string

	mostFrequency := ss[0].value
	for i := 0; i < len(ss); i++ {
		kv := ss[i]
		if kv.value == mostFrequency {
			mostFrequencyWords = append(mostFrequencyWords, kv.key)
		} else if kv.value < mostFrequency {
			break
		}
	}

	sort.Strings(mostFrequencyWords)

	fmt.Print(mostFrequencyWords[0])
}

type kv struct {
	key   string
	value int
}
