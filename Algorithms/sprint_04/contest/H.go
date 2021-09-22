package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1_000_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var s, t string

	// читаем строки s и t
	scanner.Scan()
	s = scanner.Text()

	scanner.Scan()
	t = scanner.Text()

	if len(s) != len(t) {
		fmt.Print("NO")
		return
	}

	// здесь длины строк равны
	n := len(s)

	symbolsS := []rune(s)
	symbolsT := []rune(t)

	symbolSToSymbolT := make(map[rune]rune)
	symbolTToSymbolS := make(map[rune]rune)

	var symbolS, symbolT rune
	for i := 0; i < n; i++ {
		symbolS = symbolsS[i]
		symbolT = symbolsT[i]

		relation, ok := symbolSToSymbolT[symbolS]
		if !ok { // такое соответствие ещё не встречалось
			symbolSToSymbolT[symbolS] = symbolT
		} else { // такое соответствие уже встречалось
			if relation != symbolT {
				fmt.Print("NO")
				return
			}
		}

		relation, ok = symbolTToSymbolS[symbolT]
		if !ok { // такое соответствие ещё не встречалось
			symbolTToSymbolS[symbolT] = symbolS
		} else { // такое соответствие уже встречалось
			if relation != symbolS {
				fmt.Print("NO")
				return
			}
		}
	}

	fmt.Print("YES")
	return
}
