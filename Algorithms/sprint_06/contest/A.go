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

	// для каждой вершины сохраняем связанные с ней вершины
	// ключ - номер вершины, значение - смежные с данной вершиной вершины
	adjacencyListMap := make(map[int][]int)

	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uv := strings.Split(line, " ")
		u, _ := strconv.Atoi(uv[0])
		v, _ := strconv.Atoi(uv[1])

		adjacencyListMap[u] = append(adjacencyListMap[u], v)
	}

	// выводим ответ:
	// для каждой вершины - число рёбер, исходящих из неё,
	// и вершины, в которые ведут эти рёбра, в порядке возрастания их номеров
	for i := 1; i <= n; i++ {
		adjacencyList, contains := adjacencyListMap[i]
		if !contains {
			fmt.Println(0)
		} else {
			count := len(adjacencyList)
			sort.Ints(adjacencyList)
			fmt.Print(strconv.Itoa(count) + " ")
			for j := 0; j < count-1; j++ {
				fmt.Print(strconv.Itoa(adjacencyList[j]) + " ")
			}
			fmt.Print(strconv.Itoa(adjacencyList[count-1]) + "\n")
		}
	}

}
