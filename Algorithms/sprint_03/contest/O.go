package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mergeSort(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	l, r, k := 0, 0, 0
	for true {
		if l >= len(left) || r >= len(right) {
			break
		}
		if left[l] <= right[r] {
			result[k] = left[l]
			l += 1
		} else {
			result[k] = right[r]
			r += 1
		}
		k += 1
	}

	for true {
		if l >= len(left) {
			break
		}
		result[k] = left[l]
		l += 1
		k += 1
	}

	for true {
		if r >= len(right) {
			break
		}
		result[k] = right[r]
		r += 1
		k += 1
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество островов
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем площади островов
	squares := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		squares[i] = value
	}

	// читаем число k
	var k int

	scanner.Scan()
	line = scanner.Text()
	k, _ = strconv.Atoi(line)

	// сортируем массив площадей
	sort.Ints(squares)

	// массив расстояний между минимальной площадью и всеми остальными
	distances := make([]int, n-1)
	for i := 1; i < n; i++ {
		distances[i-1] = abs(squares[0] - squares[i])
	}

	fmt.Println(distances)
	fmt.Println(k)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}
