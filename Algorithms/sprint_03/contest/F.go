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
	const maxCapacity = 32 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество отрезков
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем длины отрезков
	lengths := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		lengths[i] = value
	}

	sort.Ints(lengths)

	// идём по массиву с конца, и проверяем неравенство треугольника
	var a, b, c, result int
	for i := n - 1; i >= 2; i-- {
		a = lengths[i-2]
		b = lengths[i-1]
		c = lengths[i]
		if c < a+b {
			result = a + b + c
			break
		}
	}

	fmt.Print(strconv.Itoa(result))
}
