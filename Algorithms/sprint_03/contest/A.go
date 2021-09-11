package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func correctBracketsSequence(n, openCount, closeCount int, result string) {
	if openCount+closeCount == 2*n {
		fmt.Println(result)
		return
	}
	if openCount < n {
		correctBracketsSequence(n, openCount+1, closeCount, result+"(")
	}
	if openCount > closeCount {
		correctBracketsSequence(n, openCount, closeCount+1, result+")")
	}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	correctBracketsSequence(n, 0, 0, "")
}
