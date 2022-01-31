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
	const maxCapacity = 32 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество дней
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем значения температуры в каждый из дней
	temperatures := make([]int, n)

	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")

	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(values[i])
		temperatures[i] = value
	}

	if n == 1 {
		fmt.Println(1)
		return
	}

	result := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			if temperatures[i] > temperatures[i+1] {
				result++
			}
		} else if i == n-1 {
			if temperatures[i] > temperatures[i-1] {
				result++
			}
		} else {
			if (temperatures[i] > temperatures[i-1]) && (temperatures[i] > temperatures[i+1]) {
				result++
			}
		}
	}

	fmt.Println(result)
}
