package main

/*
https://contest.yandex.ru/contest/25597/run-report/63633502/
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// читаем строки s и t
	var s, t string

	scanner.Scan()
	s = scanner.Text()

	scanner.Scan()
	t = scanner.Text()

	/*
	 dp[i][j] — расстояние редактирования между префиксами строк:
	 a[:i] (до символа i) и b[:j] (до символа j)
	*/
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	// Базовый случай:
	// если вторая строка пустая, то нужно удалить все символы из первой строки, и наоборот
	for i := 0; i < len(s)+1; i++ {
		dp[i][0] = i
	}

	for j := 0; j < len(t)+1; j++ {
		dp[0][j] = j
	}

	// Рассматриваем различные варианты, как префикс a[:j] мог быть получен из префикса b[:i]
	// при помощи разрешённых операций
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			// 1. Если последние символы префиксов совпадают, то в этом случае можно не менять эти последние символы
			if s[i-1] == t[j-1] {
				// 3. Мы можем взять префикс a[:i-1] и превратить его в b[:j] за dp[i-1][j] операций, и затем
				// добавить в конец символ a[i-1]

				// 4. Мы можем взять префикс a[:i] и превратить его в b[:j-1] за dp[i][j-1] операций, и затем
				// добавить в конец символ b[j-1]

				// берём минимальное расстояние редактирования из всех вариантов
				dp[i][j] = minOf3(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
			} else {
				// 2. Если последние символы не совпадают, то тогда можно потратить 1 операцию на замену
				// символа a[i-1] на b[j-1]
				dp[i][j] = minOf3(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}

	fmt.Print(dp[len(s)][len(t)])
}

func minOf3(x, y, z int) int {
	return min(min(x, y), z)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
