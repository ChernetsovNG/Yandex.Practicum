package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var str string

	scanner.Scan()
	str = scanner.Text()

	// вычисляем полиномиальный хеш
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

	fmt.Println(hash)
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
