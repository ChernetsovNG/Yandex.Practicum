package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string

	// читаем число различных туристических мест
	var n int

	scanner.Scan()
	line = scanner.Text()

	n, _ = strconv.Atoi(line)

	// читаем рейтинги достопримечательностей
	r := make([]int, n)

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")

	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(values[i])
		r[i] = value
	}
}
