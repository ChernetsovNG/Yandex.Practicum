package main

import (
	"bufio"
	"fmt"
	"os"
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

	// читаем числа островов n и m
	var n, m int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем численности населения на островах в северной части
	northCount := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		northCount[i] = value
	}

	// читаем численности населения на островах в южной части
	southCount := make([]int, m)

	scanner.Scan()
	row = scanner.Text()
	values = strings.Split(row, " ")
	for i := 0; i < m; i++ {
		value, _ = strconv.Atoi(values[i])
		southCount[i] = value
	}

	result := mergeSort(northCount, southCount)

	count := n + m
	var median float64
	if count%2 == 0 { // четное количество элементов в массиве
		median = (float64(result[count/2-1]) + float64(result[count/2])) / 2.0
	} else { // нечётное количество элементов в массиве
		median = float64(result[count/2])
	}

	fmt.Print(median)
}
