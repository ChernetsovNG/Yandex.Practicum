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

	var line string

	// читаем число вершин и рёбер
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// матрица смежности
	adjacencyMatrix := make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}

	// читаем информацию о рёбрах (в виде вершин, соединяемых ребром)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uv := strings.Split(line, " ")
		u, _ := strconv.Atoi(uv[0])
		v, _ := strconv.Atoi(uv[1])

		adjacencyMatrix[u-1][v-1] = 1
	}

	// выводим матрицу смежности
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			fmt.Print(strconv.Itoa(adjacencyMatrix[i][j]) + " ")
		}
		fmt.Print(strconv.Itoa(adjacencyMatrix[i][n-1]))
		if i < n-1 {
			fmt.Print("\n")
		}
	}
}
