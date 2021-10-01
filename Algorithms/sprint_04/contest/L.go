package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1_000_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var s string

	// читаем числа n и k
	var n, k int

	scanner.Scan()
	s = scanner.Text()
	values := strings.Split(s, " ")

	n, _ = strconv.Atoi(values[0])
	k, _ = strconv.Atoi(values[1])

	// читаем строку
	scanner.Scan()
	s = scanner.Text()
	symbols := []rune(s)

	// делаем предвычисление хеша для всех подстрок исходной строки длины n
	var symbolG int
	length := len(s)
	hashes := make([]int, length-n+1)
	hashes[0] = polynomialHash(s[0:n])
	for i := 1; i < length-n+1; i++ {
		symbolA := int(symbols[i-1])
		if i == length-n {
			symbolG = 0
		} else {
			symbolG = int(symbols[i+n-1])
		}
		x := module(hashes[i-1]-symbolA*powerByModule(a, n-1, m), m)
		x = module(x*a, m)
		x = module(x+symbolG, m)
		hashes[i] = x
	}

	// определяем подстроки с одинаковыми хешами
	hashesIndexes := make(map[int][]int)
	var hash int
	for i := 0; i < length-n+1; i++ {
		hash = hashes[i]
		indexes, ok := hashesIndexes[hash]
		if !ok {
			hashesIndexes[hash] = []int{i}
		} else {
			hashesIndexes[hash] = append(indexes, i)
		}
	}

	var result []int
	// ищем подстроки, которые встречаются k и более раз
	for _, indexes := range hashesIndexes {
		count := 0
		if len(indexes) >= k {
			index := indexes[0]
			// проверяем, что строки по этим индексам действительно совпадают
			substring := s[index:(index + n)]
			count += 1
			for i := 1; i < len(indexes); i++ {
				index = indexes[i]
				if substring == s[index:(index+n)] {
					count += 1
				}
			}
			if count >= k {
				result = append(result, indexes[0])
			}
		}
	}

	printArray(result)
}

// основание, по которому вычисляется хеш
const a = 1000

// модуль m
const m = 123_987_123

func polynomialHash(str string) int {
	var hash, symbolCode int
	var symbol rune
	n := len(str)
	symbols := []rune(str)
	hash = 0
	for i := 0; i < n; i++ {
		symbol = symbols[i]
		symbolCode = int(symbol)
		hash = (hash + symbolCode*powerByModule(a, n-1-i, m)) % m
	}
	return hash
}

func powerByModule(b, e, m int) int {
	result := 1
	if 1&e == 1 {
		result = b
	}
	for true {
		if e == 0 {
			break
		}
		e >>= 1
		b = (b * b) % m
		if e&1 == 1 {
			result = (result * b) % m
		}
	}
	return result
}

func module(x, y int) int {
	if x >= 0 {
		return x % y
	} else {
		return y - ((-x) % y)
	}
}

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d", array[i])
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
