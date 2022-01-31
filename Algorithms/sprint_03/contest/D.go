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

	var line string

	// читаем количество детей
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем факторы жадности детей
	greedyFactors := make([]int, n)

	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(values[i])
		greedyFactors[i] = value
	}

	// читаем количество печенек
	var m int

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем размеры печенек
	sizes := make([]int, m)

	scanner.Scan()
	row = scanner.Text()
	values = strings.Split(row, " ")
	for i := 0; i < m; i++ {
		value, _ = strconv.Atoi(values[i])
		sizes[i] = value
	}

	// считаем, сколько детей останутся довольными
	sort.Ints(greedyFactors)
	sort.Ints(sizes)

	i, j := 0, 0
	var greedy, size int
	result := 0
	for {
		if i >= n || j >= m {
			break
		}

		greedy = greedyFactors[i]
		size = sizes[j]

		if size >= greedy { // ребёнок может взять эту печеньку
			result += 1

			i += 1
			j += 1
		} else { // иначе рассматриваем следующую печеньку
			j += 1
		}
	}

	fmt.Println(result)
}
