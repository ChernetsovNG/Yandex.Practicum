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

	// читаем число вершин и рёбер
	var n, m int

	scanner.Scan()
	line = scanner.Text()

	nm := strings.Split(line, " ")
	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	// читаем информацию о рёбрах (в виде вершин, соединяемых ребром)
	adjacencyListMap := make(map[int]*Set)

	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uv := strings.Split(line, " ")
		u, _ := strconv.Atoi(uv[0])
		v, _ := strconv.Atoi(uv[1])

		adjacencyList, contains := adjacencyListMap[u]
		if !contains {
			set := makeSet()
			set.Add(v)
			adjacencyListMap[u] = set
		} else {
			adjacencyList.Add(v)
		}

		adjacencyList, contains = adjacencyListMap[v]
		if !contains {
			set := makeSet()
			set.Add(u)
			adjacencyListMap[v] = set
		} else {
			adjacencyList.Add(u)
		}
	}

	// граф из одной вершины без рёбер будем считать связным
	if n == 1 && m == 0 {
		fmt.Print("YES")
		return
	}

	// проверяем граф на полноту
	var isFullyConnected = true

	// для каждой вершины в списке смежности должны содержаться все соседние вершины
out:
	for i := 1; i <= n; i++ {
		adjacencyList, contains := adjacencyListMap[i]
		if !contains {
			isFullyConnected = false
			break
		} else {
			for j := 1; j < n; j++ {
				if j == i {
					continue
				}
				if !adjacencyList.Exists(j) {
					isFullyConnected = false
					break out
				}
			}
		}
	}

	if isFullyConnected {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// set
func makeSet() *Set {
	return &Set{
		container: make(map[int]bool),
	}
}

type Set struct {
	container map[int]bool
}

func (c *Set) Exists(key int) bool {
	_, exists := c.container[key]
	return exists
}

func (c *Set) Add(key int) {
	c.container[key] = true
}
