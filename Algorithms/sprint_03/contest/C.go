package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 16 * 150_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var s, t string

	// читаем строки s, t
	scanner.Scan()
	s = scanner.Text()

	scanner.Scan()
	t = scanner.Text()

	result := false

	// вычисляем ответ
	symbolsS := []rune(s)
	symbolsT := []rune(t)

	lenS := len(s)
	lenT := len(t)

	if lenS == 0 || lenT == 0 {
		fmt.Print("True")
		return
	}

	i, j := 0, 0
	for true {
		if i == lenS && j <= lenT {
			result = true
			break
		}

		if j == lenT {
			break
		}
		symbolS := symbolsS[i]
		symbolT := symbolsT[j]

		if symbolS == symbolT {
			i += 1
		}
		j += 1
	}

	if result {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
