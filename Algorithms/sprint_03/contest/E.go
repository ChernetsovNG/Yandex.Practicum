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
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество домов и общий бюджет
	var n, k int

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")
	n, _ = strconv.Atoi(values[0])
	k, _ = strconv.Atoi(values[1])

	// читаем стоимости домов
	costs := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values = strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		costs[i] = value
	}

	// считаем, сколько максимум можно купить домов

	// сортируем дома по стоимости
	sort.Ints(costs)

	result := 0
	for i := 0; i < n; i++ {
		cost := costs[i]
		k -= cost
		if k >= 0 {
			result += 1
		} else {
			break
		}
	}

	fmt.Print(result)
}
