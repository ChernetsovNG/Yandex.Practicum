package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countingSort(array []int, k int) {
	countedValues := make([]int, k+1)
	for _, value := range array {
		countedValues[value] += 1
	}
	index := 0
	for value := 0; value <= k; value++ {
		for amount := 1; amount <= countedValues[value]; amount++ {
			array[index] = value
			index += 1
		}
	}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1_000_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество предметов в гардеробе
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем массив цветов предметов
	colors := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		colors[i] = value
	}

	countingSort(colors, 2)

	var b bytes.Buffer
	var color int
	for i := 0; i < n-1; i++ {
		color = colors[i]
		fmt.Fprintf(&b, "%d ", color)
	}
	if n > 0 {
		fmt.Fprintf(&b, "%d", colors[n-1])
	}

	fmt.Print(b.String())
}
