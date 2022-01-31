package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string

	// читаем число слитков n и вместимость рюкзака m
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// читаем массы слитков
	w := make([]int, n+1)

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")

	var value int
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		w[i+1] = value
	}

	/*
	 dp[i][j] — наибольшая масса слитков, которую мы можем взять, если
	 брать только слитки с номерами от 0 до i и иметь рюкзак вместимостью j.
	 Для экономии памяти будем хранить не весь массив, а только текущую
	 и предыдущую строки
	*/
	dpPrev := make([]int, m+1)
	dpCurr := make([]int, m+1)

	// вычисляем данные в массиве dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// берём максимум из двух величин:
			// 1. наибольшая масса рюкзака вместимости j, если мы не берём i-й слиток
			// 2. сумма массы i-го слитка, который мы кладём в рюкзак, и наибольшей массы рюкзака,
			// размер которого соответствует оставшемуся свободному месту в рюкзаке
			if j-w[i] < 0 {
				dpCurr[j] = dpPrev[j]
			} else {
				dpCurr[j] = max(dpPrev[j], w[i]+dpPrev[j-w[i]])
			}
		}
		for j := 0; j <= m; j++ {
			dpPrev[j] = dpCurr[j]
		}
	}

	fmt.Print(dpCurr[m])
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
