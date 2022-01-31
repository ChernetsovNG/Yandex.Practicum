package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

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
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := scanner.Text()
		values := strings.Split(row, " ")
		for j := 0; j < m; j++ {
			value, _ := strconv.Atoi(values[j])
			matrix[i][j] = value
		}
	}

	// читаем координаты элемента
	var row, col int

	scanner.Scan()
	line = scanner.Text()
	row, _ = strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()
	col, _ = strconv.Atoi(line)

	// ищем соседей
	var neighbors []int

	if row-1 >= 0 {
		neighbors = append(neighbors, matrix[row-1][col])
	}
	if row+1 <= n-1 {
		neighbors = append(neighbors, matrix[row+1][col])
	}
	if col-1 >= 0 {
		neighbors = append(neighbors, matrix[row][col-1])
	}
	if col+1 <= m-1 {
		neighbors = append(neighbors, matrix[row][col+1])
	}

	// сортируем соседние элементы по возрастанию
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i] < neighbors[j]
	})

	// выводим результат
	var neighborsText []string
	for _, neighbor := range neighbors {
		text := strconv.Itoa(neighbor)
		neighborsText = append(neighborsText, text)
	}

	result := strings.Join(neighborsText, " ")
	fmt.Println(result)
}
