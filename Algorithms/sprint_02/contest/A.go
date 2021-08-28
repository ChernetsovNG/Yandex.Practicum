package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем размеры матрицы
	var n, m int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// создаём матрицу
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	// заполняем матрицу значениями
	var value int
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := scanner.Text()
		values := strings.Split(row, " ")
		for j := 0; j < m; j++ {
			value, _ = strconv.Atoi(values[j])
			matrix[i][j] = value
		}
	}

	// транспонируем матрицу
	transpose := make([][]int, m)
	for i := range transpose {
		transpose[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	// выводим результат
	var text string
	var resultRow []string
	var result string

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			text = strconv.Itoa(transpose[i][j])
			resultRow = append(resultRow, text)
		}
		result = strings.Join(resultRow, " ")
		fmt.Println(result)
		resultRow = resultRow[:0]
	}
}
