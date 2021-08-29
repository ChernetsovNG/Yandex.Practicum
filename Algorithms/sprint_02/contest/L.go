package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func fibonacci(n, k int) int {
	if n == 0 || n == 1 {
		return 1
	}
	f1 := 1
	f2 := 1
	mod := int(math.Pow(10, float64(k)))
	var f3 int
	for i := 2; i <= n; i++ {
		f3 = (f1 + f2) % mod
		f1 = f2
		f2 = f3
	}
	return f3
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	var n, k int

	// читаем количество команд
	scanner.Scan()
	line = scanner.Text()
	nkString := strings.Split(line, " ")

	n, _ = strconv.Atoi(nkString[0])
	k, _ = strconv.Atoi(nkString[1])

	fmt.Print(fibonacci(n, k))
}
