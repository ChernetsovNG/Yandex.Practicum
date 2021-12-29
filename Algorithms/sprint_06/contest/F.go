package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = math.MaxInt32

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

	// читаем информацию о рёбрах

	// для каждой вершины сохраняем связанные с ней вершины
	// ключ - номер вершины, значение - смежные с данной вершиной вершины
	// и расстояние до каждой из них
	adjacencyListMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uv := strings.Split(line, " ")
		u, _ := strconv.Atoi(uv[0])
		v, _ := strconv.Atoi(uv[1])

		adjacencyListMap[u] = append(adjacencyListMap[u], v)
		adjacencyListMap[v] = append(adjacencyListMap[v], u)
	}

	// читаем номер стартовой вершины s и конечной вершины t
	scanner.Scan()
	line = scanner.Text()
	st := strings.Split(line, " ")
	s, _ := strconv.Atoi(st[0])
	t, _ := strconv.Atoi(st[1])

	// вычисляем массив кратчайших расстояний от вершины s до всех вершин в графе
	minDistance := Dijkstra(n, adjacencyListMap, s, t)

	// выводим минимальное расстояние до конечной вершины t
	fmt.Print(minDistance)
}

func Dijkstra(n int, graph map[int][]int, s, t int) int {
	// массив кратчайших расстояний до всех вершин графа
	dist := make([]int, n+1)
	visited := make([]bool, n+1)

	// для каждой вершины из графа
	for i := 1; i <= n; i++ {
		// задаём расстояние по умолчанию
		dist[i] = INF
		// список статусов посещённости вершин
		visited[i] = false
	}

	// расстояние от вершины до самой себя 0
	dist[s] = 0

	// пока существуют не посещённые вершины с расстоянием,
	// не равным бесконечности
	for true {
		exists := false
		for i := 1; i <= n; i++ {
			if !visited[i] && dist[i] < INF {
				exists = true
				break
			}
		}
		if !exists {
			break
		}

		// находим ещё не посещённую вершину с минимальным
		// расстоянием от s
		currentMinimum := INF
		currentMinimumVertex := -1

		// для каждой вершины v из graph.vertices
		for i := 1; i <= n; i++ {
			if !visited[i] && (dist[i] < currentMinimum) {
				currentMinimum = dist[i]
				currentMinimumVertex = i
			}
		}

		u := currentMinimumVertex
		visited[u] = true

		// в любой момент работы программы, в том числе и по её завершении,
		// для всех посещённых вершин будет записано оптимальное расстояние
		if u == t {
			// если посещённая вершина совпадает с конечной, то значит для неё уже
			// вычислено оптимальное расстояние, которое мы возвращаем
			return dist[u]
		}

		// из множества рёбер графа выбираем те, которые исходят из u
		adjacencyList, _ := graph[u]
		// для каждого ребра (u, v) среди рёбер к соседним вершинам
		for _, v := range adjacencyList {
			// проверяем, не получился ли путь короче найденного ранее
			if dist[v] > dist[u]+1 {
				dist[v] = dist[u] + 1
			}
		}
	}

	return -1
}
