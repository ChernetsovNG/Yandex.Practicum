package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string

	// читаем сумму, которую нужно получить из банкомата
	var x int

	scanner.Scan()
	line = scanner.Text()

	x, _ = strconv.Atoi(line)

	// читаем число различных номиналов банкнот
	var k int

	scanner.Scan()
	line = scanner.Text()

	k, _ = strconv.Atoi(line)

	// читаем номиналы купюр
	nominals := make([]int, k)

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")

	var value int
	for i := 0; i < k; i++ {
		value, _ = strconv.Atoi(values[i])
		nominals[i] = value
	}

	/*
	 dp[i] — минимальное число купюр, которыми можно набрать сумму i
	*/
	dp := make([]int, x+1)

	// Динамика: dp[i] = min(dp[i-aj]) + 1 по всем номиналам j

	// Базовый случай: сумму 0 мы можем набрать 0 купюр
	dp[0] = 0

	for i := 1; i <= x; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < k; j++ {
			if i-nominals[j] >= 0 {
				dp[i] = min(dp[i], dp[i-nominals[j]]+1)
			}
		}
	}

	if dp[x] < math.MaxInt32 {
		fmt.Print(dp[x])
	} else {
		fmt.Print(-1)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
