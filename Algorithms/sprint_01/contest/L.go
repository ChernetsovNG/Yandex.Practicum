package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()

	scanner.Scan()
	t := scanner.Text()

	sSymbols := []rune(s)
	tSymbols := []rune(t)

	tMap := make(map[rune]int)

	var symbol rune
	var contains bool

	for i := 0; i < len(tSymbols); i++ {
		symbol = tSymbols[i]
		_, contains = tMap[symbol]
		if !contains {
			tMap[symbol] = 1
		} else {
			tMap[symbol] = tMap[symbol] + 1
		}
	}

	for i := 0; i < len(sSymbols); i++ {
		symbol = sSymbols[i]
		_, contains = tMap[symbol]
		if !contains {
			fmt.Printf("%c", symbol)
			return
		} else {
			tMap[symbol] = tMap[symbol] - 1
			if tMap[symbol] == 0 {
				delete(tMap, symbol)
			}
		}
	}

	for k, _ := range tMap {
		fmt.Printf("%c", k)
	}
}
