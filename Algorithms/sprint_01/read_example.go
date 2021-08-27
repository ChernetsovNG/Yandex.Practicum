package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		var value1, value2 int
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		value1, _ = strconv.Atoi(values[0])
		value2, _ = strconv.Atoi(values[1])
		result := value1 + value2
		outputString := strconv.Itoa(result)
		_, err := writer.WriteString(outputString)
		if err != nil {
			return
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
