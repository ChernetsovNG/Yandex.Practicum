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

	var a, x, b, c int
	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	a, _ = strconv.Atoi(values[0])
	x, _ = strconv.Atoi(values[1])
	b, _ = strconv.Atoi(values[2])
	c, _ = strconv.Atoi(values[3])

	result := a*x*x + b*x + c

	outputString := strconv.Itoa(result)
	writer.WriteString(outputString)
	writer.Flush()
}
