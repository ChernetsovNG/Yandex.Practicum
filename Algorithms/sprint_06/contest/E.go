package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack struct {
	array []int
}

func newStack() Stack {
	return Stack{[]int{}}
}

func (stack *Stack) push(x int) {
	stack.array = append(stack.array, x)
}

func (stack *Stack) pop() int {
	valueToReturn := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return valueToReturn
}

func (stack *Stack) top() int {
	return stack.array[len(stack.array)-1]
}

func (stack *Stack) size() int {
	return len(stack.array)
}

func (stack *Stack) isEmpty() bool {
	return stack.size() == 0
}

type kv struct {
	Key   int
	Value int
}

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
		adjacencyListMap[v] = append(adjacencyListMap[v], u)
	}

	// цвета вершин: 1 - white, 2 - gray
	color := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		color[i] = 1
	}

	componentColor := 2
	// будем сохранять узлы, относящиеся к каждой компоненте связности
	connectivityComponentsMap := make(map[int][]int)

	dfsStack := newStack()
	for i := 1; i <= n; i++ {
		if color[i] == 1 { // white
			componentColor += 1
			dfsStack.push(i) // начинаем обход очередной компоненты связности
			for {
				if dfsStack.isEmpty() {
					break
				}
				// пока стек не пуст
				// получаем из стека очередную вершину
				v := dfsStack.pop()

				if color[v] == 1 { // color[v] == white
					// красим вершину в серый и кладём её обратно в стек
					color[v] = 2
					dfsStack.push(v)

					// добавляем в стек все не посещённые соседние вершины
					adjacencyList, _ := adjacencyListMap[v]
					// для каждого исходящего из v ребра
					for _, w := range adjacencyList {
						if color[w] == 1 { // white
							dfsStack.push(w)
						}
					}
				} else if color[v] == 2 { // gray
					// серую вершину мы могли получить из стека только на обратном пути => красим её в цвет
					// компоненты связности
					color[v] = componentColor
					connectivityComponentsMap[componentColor] = append(connectivityComponentsMap[componentColor], v)
				}
			}
		}
	}

	for _, connectivityComponent := range connectivityComponentsMap {
		sort.Ints(connectivityComponent)
	}

	// выводим количество компонент связности
	fmt.Println(len(connectivityComponentsMap))

	// выводим компоненты связности, упорядоченные по номеру первой вершины
	var ss []kv
	for k, v := range connectivityComponentsMap {
		ss = append(ss, kv{k, v[0]})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		printArray(connectivityComponentsMap[kv.Key])
	}
}

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
