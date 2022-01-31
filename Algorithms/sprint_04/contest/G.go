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
	const maxCapacity = 32 * 100_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество раундов
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем результаты раундов
	results := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		results[i] = value
	}

	// массив, в котором будем накапливать сумму баллов
	summs := make([]int, n)

	var result int
	for i := 0; i < n; i++ {
		result = results[i]
		if result == 0 {
			if i == 0 {
				summs[i] = -1
			} else {
				summs[i] = summs[i-1] - 1
			}
		} else if result == 1 {
			if i == 0 {
				summs[i] = 1
			} else {
				summs[i] = summs[i-1] + 1
			}
		}
	}

	// словарь, в котором для каждой суммы хранятся индексы ячеек, в которых она встречается
	summsIndexes := make(map[int][]int)
	summsIndexes[0] = []int{-1}

	var sum int
	for i := 0; i < n; i++ {
		sum = summs[i]
		indexes, ok := summsIndexes[sum]
		if !ok {
			summsIndexes[sum] = []int{i}
		} else {
			summsIndexes[sum] = append(indexes, i)
		}
	}

	// определяем наибольшую разность индексов для каждой суммы
	maxDistance := -1

	var min, max, dist int
	for _, indexes := range summsIndexes {
		min = indexes[0]
		max = indexes[len(indexes)-1]
		dist = max - min
		if dist > maxDistance {
			maxDistance = dist
		}
	}

	fmt.Println(maxDistance)
}
