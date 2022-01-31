package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	// читаем первую и вторую строки
	scanner.Scan()
	first := scanner.Text()

	scanner.Scan()
	second := scanner.Text()

	// оставляем символы на чётных позициях в английском алфавите
	firstFilter := filter(first)
	secondFilter := filter(second)

	// сравниваем строки
	if firstFilter < secondFilter {
		fmt.Print(-1)
	} else if firstFilter == secondFilter {
		fmt.Print(0)
	} else {
		fmt.Print(1)
	}
}

func filter(s string) string {
	runes := []rune(s)
	var result []rune
	for _, r := range runes {
		if position(r)%2 == 0 {
			result = append(result, r)
		}
	}
	return string(result)
}

func position(r rune) int {
	return int(r) - 96
}
