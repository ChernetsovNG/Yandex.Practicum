package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const module = 1000000007

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем число вершин и рёбер
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// читаем информацию о рёбрах (в виде вершин, соединяемых ребром)

	// для каждой вершины сохраняем список вершин, из которых в неё входят рёбра
	// ключ - номер вершины, значение - смежные с данной вершиной вершины, из которых
	// в вершину входят рёбра (своего рода "обратный" список смежности)
	invAdjacencyList := make(map[int][]int)

	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uv := strings.Split(line, " ")
		u, _ := strconv.Atoi(uv[0])
		v, _ := strconv.Atoi(uv[1])
		invAdjacencyList[v] = append(invAdjacencyList[v], u)
	}

	// читаем вершины a и b, между которыми нужно посчитать количество путей
	var a, b int

	scanner.Scan()
	line = scanner.Text()

	ab := strings.Split(line, " ")
	a, _ = strconv.Atoi(ab[0])
	b, _ = strconv.Atoi(ab[1])

	// вычисляем количество путей (по модулю module)

	// dp[i] - количество путей от вершины a до i
	dp := make([]int, n+1)

	// w - массив такой, что w[v] = true, если ответ для вершины v уже посчитан,
	// и w[v] = false в противном случае
	w := make([]bool, n+1)

	dp[a] = 1
	w[a] = true
	answer := count(a, b, &w, &dp, invAdjacencyList)

	fmt.Print(answer)
}

// количество путей между вершинами g и v
func count(g, v int, w *[]bool, d *[]int, invAdjacencyList map[int][]int) int {
	if (*w)[v] == true { // ответ уже посчитан
		return (*d)[v] % module
	} else {
		sum := 0
		(*w)[v] = true
		// для всех вершин, из которых в вершину v ведут рёбра
		for _, c := range invAdjacencyList[v] {
			sum += count(g, c, w, d, invAdjacencyList) % module
		}
		(*d)[v] = sum % module
		return sum % module
	}
}
