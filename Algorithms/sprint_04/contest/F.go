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
	const maxCapacity = 32 * 60_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество строк
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	symbolsCodes := make(map[rune]int)
	symbolsCodes['a'] = 2
	symbolsCodes['b'] = 3
	symbolsCodes['c'] = 5
	symbolsCodes['d'] = 7
	symbolsCodes['e'] = 11
	symbolsCodes['f'] = 13
	symbolsCodes['g'] = 17
	symbolsCodes['h'] = 19
	symbolsCodes['i'] = 23
	symbolsCodes['j'] = 29
	symbolsCodes['k'] = 31
	symbolsCodes['l'] = 37
	symbolsCodes['m'] = 41
	symbolsCodes['n'] = 43
	symbolsCodes['o'] = 47
	symbolsCodes['p'] = 53
	symbolsCodes['q'] = 59
	symbolsCodes['r'] = 61
	symbolsCodes['s'] = 67
	symbolsCodes['t'] = 71
	symbolsCodes['u'] = 73
	symbolsCodes['v'] = 79
	symbolsCodes['w'] = 83
	symbolsCodes['x'] = 89
	symbolsCodes['y'] = 97
	symbolsCodes['z'] = 101

	// ключ - хеш код анаграммы, значение - список индексов анаграмм с таким хеш кодом
	anagramsDict := make(map[int][]int)

	// читаем строки
	scanner.Scan()
	line = scanner.Text()
	var s string
	var hashCode int
	values := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		s = values[i]
		hashCode = hash(s, symbolsCodes)
		anagrams, ok := anagramsDict[hashCode]
		if !ok {
			anagramsDict[hashCode] = []int{i}
		} else {
			anagramsDict[hashCode] = append(anagrams, i)
		}
	}

	// выводим результат
	valuesArray := make([][]int, 0, len(anagramsDict))

	for _, value := range anagramsDict {
		valuesArray = append(valuesArray, value)
	}
	// сортируем массивы по первому индесу
	sort.Slice(valuesArray, func(i, j int) bool {
		value1 := valuesArray[i]
		value2 := valuesArray[j]

		f1 := value1[0]
		f2 := value2[0]

		if f1 <= f2 {
			return true
		} else {
			return false
		}
	})

	for i := 0; i < len(valuesArray); i++ {
		printArray(valuesArray[i])
	}
}

const bigPrime = 100_000_064_207

func hash(s string, symbolsCodes map[rune]int) int {
	symbols := []rune(s)
	var symbol rune
	var code int
	hash := 1
	for i := 0; i < len(symbols); i++ {
		symbol = symbols[i]
		code = symbolsCodes[symbol]
		hash = (hash * code) % bigPrime
	}
	return hash
}

func printArray(array []int) {
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d", array[i])
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
