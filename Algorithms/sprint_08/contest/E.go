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
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем строку s
	scanner.Scan()
	s := scanner.Text()

	// читаем количество подаренных строк
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем информацию о подаренных строках - пары t_i и k_i
	kt := make(map[int]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		tk := strings.Split(line, " ")
		ti := tk[0]
		ki, _ := strconv.Atoi(tk[1])

		kt[ki] = ti
	}

	// сортируем позиции вставляемых строк по возрастанию
	keys := make([]int, 0, len(kt))
	for k := range kt {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// вставляем строки

	// строка ti вставляется в исходную строку s сразу после символа с номером ki
	resultLen := len(s)
	for _, t := range kt {
		resultLen += len(t)
	}
	resultArr := make([]rune, resultLen)

	sRunes := []rune(s)
	fromIndexS := 0
	fromIndexResult := 0
	for _, k := range keys {
		t := kt[k]

		for i := 0; i < k-fromIndexS; i++ {
			resultArr[fromIndexResult] = sRunes[fromIndexS+i]
			fromIndexResult++
		}
		for i := 0; i < len(t); i++ {
			resultArr[fromIndexResult] = rune(t[i])
			fromIndexResult++
		}

		fromIndexS += k - fromIndexS
	}

	if fromIndexS < len(s) {
		for i := 0; i < len(s)-fromIndexS; i++ {
			resultArr[fromIndexResult] = sRunes[fromIndexS+i]
			fromIndexResult++
		}
	}

	fmt.Print(string(resultArr))
}
