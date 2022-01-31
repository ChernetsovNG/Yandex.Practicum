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

	// читаем число предметов и грузоподъёмность рюкзака
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// читаем информацию о предметах: масса и стоимость
	mass := make([]int, n+1)
	cost := make([]int, n+1)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		massCost := strings.Split(line, " ")
		readMass, _ := strconv.Atoi(massCost[0])
		readCost, _ := strconv.Atoi(massCost[1])

		mass[i+1] = readMass
		cost[i+1] = readCost
	}

	/*
	 dp[i][j] — максимальная значимость предметов, которую мы можем взять, если
	 брать только предметы с номерами от 1 до i и иметь рюкзак вместимостью j.
	*/
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// вычисляем данные в массиве dp
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// берём максимум из двух величин:
			// 1. наибольшая ценность рюкзака вместимости j, если мы не берём i-й предмет
			// 2. сумма ценности i-го предмета, который мы кладём в рюкзак, и наибольшей
			// ценности рюкзака, размер которого соответствует оставшемуся свободному месту в рюкзаке
			// dp[i][j] = max(dp[i - 1][j], cost[i] + dp[i - 1][j - weight[i]])
			if j-mass[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], cost[i]+dp[i-1][j-mass[i]])
			}
		}
	}

	// восстанавливаем номера предметов, которые мы берём для достижения максимальной ценности рюкзака
	i := n
	j := m

	var answer []int

	for {
		if i == 0 {
			break
		}
		if j-mass[i] < 0 {
			// не могли взять i-ый предмет
			i -= 1
		} else {
			if cost[i]+dp[i-1][j-mass[i]] >= dp[i-1][j] {
				// мы брали i-ый предмет
				answer = append(answer, i)
				j -= mass[i]
				i -= 1
			} else {
				// мы не брали i-ый предмет
				i -= 1
			}
		}
	}

	// выводим ответ:

	// количество предметов
	fmt.Println(len(answer))

	// номера предметов
	if len(answer) > 0 {
		for i := 0; i < len(answer)-1; i++ {
			fmt.Printf("%d ", answer[i])
		}
		fmt.Print(answer[len(answer)-1])
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
