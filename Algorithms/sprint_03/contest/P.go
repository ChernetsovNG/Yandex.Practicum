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
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество чисел
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем массив
	numbers := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		numbers[i] = value
	}

	// вычисляем решение
	minimums := make([]int, n)
	maximums := make([]int, n)

	currentMax := numbers[0]
	currentMin := numbers[n-1]

	maximums[0] = currentMax
	minimums[n-1] = currentMin

	for i := 1; i < n; i++ {
		if numbers[i] > currentMax {
			currentMax = numbers[i]
		}
		maximums[i] = currentMax
	}
	for i := n - 2; i >= 0; i-- {
		if numbers[i] < currentMin {
			currentMin = numbers[i]
		}
		minimums[i] = currentMin
	}

	// находим границы блоков и их количество
	count := 0

	for i := 0; i < n-1; i++ {
		if maximums[i] < minimums[i+1] {
			count += 1
		}
	}

	fmt.Print(count + 1)
}
