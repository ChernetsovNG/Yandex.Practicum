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
	scanner.Split(bufio.ScanLines)

	var line string

	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	var factors []string

	divisor := 2
	for true {
		if divisor*divisor > n {
			break
		}
		if n%divisor == 0 {
			n /= divisor
			factors = append(factors, strconv.Itoa(divisor))
		} else {
			divisor += 1
		}
	}
	if n != 1 {
		factors = append(factors, strconv.Itoa(n))
	}

	result := strings.Join(factors, " ")
	fmt.Println(result)
}
