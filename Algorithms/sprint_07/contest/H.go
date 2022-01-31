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

	// читаем размеры поля n и m
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// создаём матрицу
	matrix := make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, m)
	}

	// заполняем матрицу значениями
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		for j, char := range line {
			if char == '0' {
				matrix[i][j] = false
			} else if char == '1' {
				matrix[i][j] = true
			}
		}
	}

	/*
	 dp[i][j] — максимально возможное число цветков, которое можно собрать,
	 добравшись до ячейки (i, j).
	*/
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// базовый случай
	if matrix[n-1][0] == true {
		dp[n-1][0] = 1
	} else {
		dp[n-1][0] = 0
	}

	// формула перехода:
	// dp[i][j] = max(dp[i+1][j], dp[i][j-1]) + points[i][j]
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if i < n-1 {
				if j > 0 {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1]) + boolValue(matrix[i][j])
				} else {
					dp[i][j] = dp[i+1][j] + boolValue(matrix[i][j])
				}
			} else { // i == n-1
				if j > 0 {
					dp[i][j] = dp[i][j-1] + boolValue(matrix[i][j])
				} else {
					dp[i][j] = boolValue(matrix[i][j]) // базовый случай - левая нижняя ячейка
				}
			}
		}
	}

	fmt.Print(dp[0][m-1])
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func boolValue(val bool) int {
	if val == true {
		return 1
	} else {
		return 0
	}
}
