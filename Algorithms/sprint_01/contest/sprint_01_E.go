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
	const maxCapacity = 4 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество символов во входной строке
	var _ int

	scanner.Scan()
	line = scanner.Text()
	_, _ = strconv.Atoi(line)

	// читаем саму строку с текстом
	scanner.Scan()
	row := scanner.Text()

	// разбиваем на отдельные слова
	words := strings.Split(row, " ")

	var maxLen = 0
	var maxWord string

	for _, word := range words {
		wordLen := len(word)
		if wordLen > maxLen {
			maxWord = word
			maxLen = wordLen
		}
	}

	fmt.Println(maxWord)
	fmt.Println(maxLen)
}
