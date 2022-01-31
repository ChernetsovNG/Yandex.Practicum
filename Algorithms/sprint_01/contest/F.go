package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 4 * 20_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	// читаем входной текст
	scanner.Scan()
	line := strings.ToLower(scanner.Text())

	reg, _ := regexp.Compile("[^a-z0-9]+")
	line = reg.ReplaceAllString(line, "")

	var isPalindrome = true

	var i, j int
	i = 0
	j = len(line) - 1

	for i = 0; i < j; {
		if line[i] != line[j] {
			isPalindrome = false
			break
		}
		i++
		j--
	}

	if isPalindrome {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
