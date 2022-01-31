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

	// читаем число n - количество элементов в первой последовательности
	var n int

	scanner.Scan()
	line = scanner.Text()

	n, _ = strconv.Atoi(line)

	// читаем элементы первой последовательности
	a := make([]int, n+1)

	scanner.Scan()
	line = scanner.Text()
	values := strings.Split(line, " ")

	a[0] = -1
	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(values[i])
		a[i+1] = value
	}

	// читаем число m - количество элементов во второй последовательности
	var m int

	scanner.Scan()
	line = scanner.Text()

	m, _ = strconv.Atoi(line)

	// читаем элементы второй последовательности
	b := make([]int, m+1)

	scanner.Scan()
	line = scanner.Text()
	values = strings.Split(line, " ")

	b[0] = -1
	for i := 0; i < m; i++ {
		value, _ = strconv.Atoi(values[i])
		b[i+1] = value
	}

	// находим наибольшую общую подпоследовательность (НОП)

	/*
	 в dp[i][j] будем хранить длину НОП для подстрок a[1:i] и b[1:j]
	*/
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// в процессе перехода динамики мы удлиняем одну из строк на один символ
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// если до удлинения крайние символы в двух строках были равны и включены в НОП,
			// то добавление нового символа не изменит длину НОП
			if a[i] == b[j-1] {
				dp[i][j] = dp[i][j-1]
			} else {
				// иначе надо сравнить, не совпал ли новый символ с символом на конце второй строки.
				// Если совпал, то увеличиваем НОП
				if a[i] == b[j] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}

	// длина наибольшей общей подпоследовательности
	length := dp[n][m]

	// восстанавливаем НОП

	// в массиве будем хранить НОП, записанную от конца к началу
	var answerA []int
	var answerB []int

	// начинаем с клетки dp[n][m]
	i := n
	j := m

	for {
		if dp[i][j] == 0 {
			break
		}
		if a[i] == b[j] {
			answerA = append(answerA, i)
			answerB = append(answerB, j)
			i -= 1
			j -= 1
		} else {
			if dp[i][j] == dp[i-1][j] {
				// перемещаемся вверх, в соседнюю ячейку
				i -= 1
			} else if dp[i][j] == dp[i][j-1] {
				// перемещаемся в ячейку левее текущей
				j -= 1
			}
		}
	}

	// выводим длину НОП
	fmt.Println(length)

	if len(answerA) > 0 && len(answerB) > 0 {
		// индексы элементов первой последовательности, которые участвуют в НОП
		for i := len(answerA) - 1; i >= 1; i-- {
			fmt.Printf("%d ", answerA[i])
		}
		fmt.Printf("%d\n", answerA[0])

		// индексы элементов второй последовательности, которые участвуют в НОП
		for i := len(answerB) - 1; i >= 1; i-- {
			fmt.Printf("%d ", answerB[i])
		}
		fmt.Printf("%d\n", answerB[0])
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
