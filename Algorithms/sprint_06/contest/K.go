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

type Pair struct {
	to   int
	dist int
}

func newPair(to, dist int) Pair {
	return Pair{to, dist}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем число достопримечательностей и мостов
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// читаем информацию о мостах

	// для каждой вершины сохраняем связанные с ней вершины
	// ключ - номер вершины, значение - смежные с данной вершиной вершины
	// и расстояние до каждой из них
	adjacencyListMap := make(map[int][]Pair)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uvl := strings.Split(line, " ")
		u, _ := strconv.Atoi(uvl[0])
		v, _ := strconv.Atoi(uvl[1])
		l, _ := strconv.Atoi(uvl[2])

		adjacencyListMap[u] = append(adjacencyListMap[u], newPair(v, l))
		adjacencyListMap[v] = append(adjacencyListMap[v], newPair(u, l))
	}

	// для каждой вершины запустим алгоритм Дейкстры, который найдёт кратчайшие расстояния
	// от неё до остальных вершин в графе
	for i := 1; i <= n; i++ {
		minDistances := Dijkstra(n, adjacencyListMap, i)
		printArray(minDistances)
	}
}

func Dijkstra(n int, graph map[int][]Pair, s int) []int {
	// массив кратчайших расстояний до всех вершин графа
	dist := make([]int, n+1)
	previous := make([]int, n+1)
	visited := make([]bool, n+1)

	// для каждой вершины из графа
	for i := 1; i <= n; i++ {
		// задаём расстояние по умолчанию
		dist[i] = INF
		// задаём предшественника для восстановления SPT
		previous[i] = -1
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
		// из множества рёбер графа выбираем те, которые исходят из u

		adjacencyList, _ := graph[u]
		// для каждого ребра (u, v) среди рёбер к соседним вершинам
		for _, neighbor := range adjacencyList {
			// проверяем, не получился ли путь короче найденного ранее
			v := neighbor.to
			weight := neighbor.dist
			if dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
				previous[v] = u
			}
		}
	}

	// все расстояния, которые остались равными бесконечности, заменяем на -1,
	// т.к. вершина недостижима
	for i := 1; i <= n; i++ {
		if dist[i] == INF {
			dist[i] = -1
		}
	}

	return dist
}

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 1; i < len(array)-1; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
