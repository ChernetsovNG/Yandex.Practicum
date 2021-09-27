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

	var line string

	// читаем основание, по которому вычисляется хеш
	var a int

	scanner.Scan()
	line = scanner.Text()
	a, _ = strconv.Atoi(line)

	// читаем модуль m
	var m int

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем строку
	var s string

	scanner.Scan()
	s = scanner.Text()

	// делаем предвычисление хеша для всех подстрок исходной строки
	length := len(s)
	hashes := make([]int, length)
	for i := 1; i <= length; i++ {
		hashes[i-1] = polynomialHash(a, m, s, i)
	}

	// читаем число запросов t
	var t int

	scanner.Scan()
	line = scanner.Text()
	t, _ = strconv.Atoi(line)

	// читаем индексы начала и конца подстроки для запросов
	var l, r int

	for i := 0; i < t; i++ {
		scanner.Scan()
		line = scanner.Text()
		split := strings.Split(line, " ")
		l, _ = strconv.Atoi(split[0])
		r, _ = strconv.Atoi(split[1])

		// печатаем хеш для подстроки
		hashR := hashes[r-1]
		hashL := hashes[l-1]
		hashL = (hashL * powerByModule(a, l, m)) % m

		hashSubstring := (hashR - hashL) % m
		fmt.Println(hashSubstring)
	}
}

// хеш для подстроки строки str до символа с индексом indexTo включительно
func polynomialHash(a, m int, str string, indexTo int) int {
	var hash, symbolCode int
	symbols := []rune(str)
	hash = 0
	for i := 0; i < indexTo; i++ {
		symbolCode = int(symbols[i])
		hash = (hash + symbolCode*powerByModule(a, indexTo-1-i, m)) % m
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
