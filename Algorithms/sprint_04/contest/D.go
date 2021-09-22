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

	var line string

	// читаем количество записей в логе
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем записи лога
	dict := make(map[string]int)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		_, ok := dict[line]
		if !ok {
			dict[line] = 1
			fmt.Println(line)
		}
	}
}
