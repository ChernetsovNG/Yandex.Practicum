package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split()

	var line string

	// читаем количество элементов в дереве
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// Ответ - соответствующее число Каталана

	// массив для вычисления
	c := make([]int64, n+1)

	c[0] = 1
	c[1] = 1

	for k := 2; k <= n; k++ {
		var sum int64 = 0
		for i := 0; i < k; i++ {
			sum += c[i] * c[k-1-i]
		}
		c[k] = sum
	}

	fmt.Print(c[n])
}
