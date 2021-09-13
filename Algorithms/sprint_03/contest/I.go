package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type kv struct {
	key   int
	value int
}

func mergeSort(array []kv) []kv {
	if len(array) == 1 {
		return array
	}
	left := mergeSort(array[0 : len(array)/2])
	right := mergeSort(array[len(array)/2 : len(array)])

	result := make([]kv, len(array))

	l, r, k := 0, 0, 0
	for true {
		if l >= len(left) || r >= len(right) {
			break
		}
		if left[l].value >= right[r].value {
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

	// читаем количество студентов
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем ID вузов для каждого студента
	ids := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		ids[i] = value
	}

	// читаем число k
	var k int

	scanner.Scan()
	line = scanner.Text()
	k, _ = strconv.Atoi(line)

	var id int
	idCountMap := make(map[int]int)
	for i := 0; i < n; i++ {
		id = ids[i]
		idCountMap[id] += 1
	}

	// сортируем словарь по убыванию значений
	var ss []kv
	for k, v := range idCountMap {
		ss = append(ss, kv{k, v})
	}

	ss = mergeSort(ss)

	// выводим первые k значений
	var b bytes.Buffer
	var pair kv
	for i := 0; i < k-1; i++ {
		pair = ss[i]
		fmt.Fprintf(&b, "%d ", pair.key)
	}
	fmt.Fprintf(&b, "%d", ss[k-1].key)

	fmt.Print(b.String())
}
