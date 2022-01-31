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
	symbols := []rune(s)

	// делаем предвычисление хеша для всех подстрок исходной строки
	length := len(s)
	hashes := make([]int, length)
	hashes[0] = int(symbols[0]) % m

	for i := 1; i < length; i++ {
		hashes[i] = (hashes[i-1]*a + (int(symbols[i]))) % m
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

		// вычисляем хеш для подстроки s[l-1...r-1]
		// h[L...R] = (h[0...R] - h[0...(L-1)]*q^(R-L+1)) mod m
		var hash int
		if l <= 1 {
			hash = hashes[r-1]
		} else {
			hash = module(hashes[r-1]-hashes[l-2]*powerByModule(a, r-l+1, m), m)
		}
		fmt.Println(hash)
	}
}

func polynomialHash(a, m int, str string) int {
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
	for {
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
