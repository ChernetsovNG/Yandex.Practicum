package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	// читаем вход
	var n int

	scanner.Scan()
	line := scanner.Text()
	n, _ = strconv.Atoi(line)

	if n == 0 {
		fmt.Print("0")
		return
	}

	var result []string

	for true {
		if n <= 0 {
			break
		}
		result = append(result, strconv.Itoa(n%2))
		n /= 2
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}
