package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	a, _ = strconv.Atoi(values[0])
	b, _ = strconv.Atoi(values[1])
	c, _ = strconv.Atoi(values[2])

	a1 := parity(a)
	b1 := parity(b)
	c1 := parity(c)

	var resultBool bool
	if a1 && b1 && c1 {
		resultBool = true
	} else if !a1 && !b1 && !c1 {
		resultBool = true
	} else {
		resultBool = false
	}

	var result string
	if resultBool {
		result = "WIN"
	} else {
		result = "FAIL"
	}

	writer.WriteString(result)
	writer.Flush()
}

func parity(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}
