package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	// читаем строку
	var s string

	scanner.Scan()
	s = scanner.Text()
	symbols := []rune(s)
	n := len(s)

	if n == 0 {
		fmt.Print("0")
		return
	} else if n == 1 {
		fmt.Print("1")
		return
	}
	maxLength := -1

	previousSymbols := make(map[rune]int)
	leftIndex := 0
	length := 0
	for index := 0; index < n; index++ {
		symbol := symbols[index]

		prevIndex, ok := previousSymbols[symbol]
		if ok { // если такой символ уже встречался
			// и он находится в пределах рассматриваемой подстроки
			if prevIndex >= leftIndex {
				leftIndex = prevIndex + 1
			}
		}
		length = index - leftIndex + 1
		if length > maxLength {
			maxLength = length
		}

		// сохраняем индекс, по которому встречается символ
		previousSymbols[symbol] = index
	}

	fmt.Print(maxLength)
}
