package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	key   int
	value int
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

	deltasCountMap := make(map[int]int)
	var delta int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s1 := squares[i]
			s2 := squares[j]
			delta = abs(s1 - s2)
			deltasCountMap[delta] += 1
		}
	}

	// сортируем словарь по возрастанию delta
	var ss []kv
	for k, v := range deltasCountMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].key < ss[j].key
	})

	var keyValue kv
	var deltaKey, count, overall, result int
	overall = 0
	for i := 0; i < len(ss); i++ {
		keyValue = ss[i]
		deltaKey = keyValue.key
		count = keyValue.value
		overall += count
		if overall >= k {
			result = deltaKey
			break
		}
	}

	fmt.Println(result)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}
