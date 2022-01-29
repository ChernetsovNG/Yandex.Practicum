package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 1000000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var s string

	// читаем входную строку
	scanner.Scan()
	s = scanner.Text()

	writer := bufio.NewWriter(os.Stdout)

	// вычисляем (и сразу печатаем) префикс функцию
	n := len(s)
	pi := make([]int, n)

	pi[0] = 0

	if n == 1 {
		writer.WriteString(strconv.Itoa(pi[0]))
	} else {
		writer.WriteString(strconv.Itoa(pi[0]))
		writer.WriteString(" ")
	}

	for i := 1; i < n; i++ {
		k := pi[i-1]
		for {
			if k <= 0 || s[k] == s[i] { // пока (k > 0) и (s[k] ≠ s[i])
				break
			}
			k = pi[k-1]
		}
		if s[k] == s[i] {
			k += 1
		}
		pi[i] = k
		if i == n-1 {
			writer.WriteString(strconv.Itoa(pi[i]))
		} else {
			writer.WriteString(strconv.Itoa(pi[i]))
			writer.WriteString(" ")
		}
	}

	writer.Flush()
}
