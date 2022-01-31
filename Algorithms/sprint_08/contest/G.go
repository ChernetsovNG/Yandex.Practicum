package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество измерений
	var n int

	scanner.Scan()
	line = scanner.Text()

	n, _ = strconv.Atoi(line)

	// читаем результаты измерений
	x := make([]int, n)

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")

	var value int
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		x[i] = value
	}

	// читаем длину шаблона
	var m int

	scanner.Scan()
	line = scanner.Text()

	m, _ = strconv.Atoi(line)

	// читаем элементы шаблона
	a := make([]int, m)

	scanner.Scan()
	line = scanner.Text()
	values = strings.Split(line, " ")

	for i := 0; i < m; i++ {
		value, _ = strconv.Atoi(values[i])
		a[i] = value
	}

	printArray(find(x, a))
}

// ищем все вхождения, но со сдвигом на константу. Для этого все символы строки изменяем так,
// чтобы первый символ совпадал с первым символом шаблона
func find(text []int, pattern []int) []int {
	if len(pattern) > len(text) { // длинный шаблон не может содержаться в короткой строке
		return []int{}
	}
	var occurrences []int
	for pos := 0; pos <= len(text)-len(pattern); pos++ {
		delta := pattern[0] - text[pos]
		match := true
		for offset := 0; offset <= len(pattern)-1; offset++ {
			if text[pos+offset]+delta != pattern[offset] {
				// одного несовпадения достаточно, чтобы не проверять дальше текущее расположение шаблона
				match = false
				break
			}
		}
		// как только нашлось совпадение шаблона, возвращаем его
		// это первое вхождение шаблона в строку
		if match == true {
			occurrences = append(occurrences, pos+1)
		}
	}
	return occurrences
}

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
