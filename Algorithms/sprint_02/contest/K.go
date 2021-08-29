package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	var n int

	// читаем количество команд
	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	fmt.Print(fibonacci(n))
}
