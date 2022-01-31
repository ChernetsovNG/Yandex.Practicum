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

	// сортируем списки смежности вершин в порядке убывания
	for _, adjacencyList := range adjacencyListMap {
		sort.Sort(sort.Reverse(sort.IntSlice(adjacencyList)))
	}

	// читаем информацию о стартовой вершине
	scanner.Scan()
	line = scanner.Text()
	s, _ := strconv.Atoi(line)

	// цвета вершин: 1 - white, 2 - gray, 3 - black
	color := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		color[i] = 1
	}

	// выполняем обход вершин
	stack := newStack()

	stack.push(s)

	// массив вершин в порядке обхода
	var result []int

	for true {
		if stack.isEmpty() {
			break
		}
		// пока стек не пуст
		// получаем из стека очередную вершину
		v := stack.pop()

		if color[v] == 1 { // color[v] == white
			result = append(result, v)

			// красим вершину в серый и кладём её обратно в стек
			color[v] = 2
			stack.push(v)

			// добавляем в стек все не посещённые соседние вершины
			adjacencyList, _ := adjacencyListMap[v]
			// для каждого исходящего из v ребра
			for _, w := range adjacencyList {
				if color[w] == 1 { // white
					stack.push(w)
				}
			}
		} else if color[v] == 2 { // gray
			// серую вершину мы могли получить из стека только на обратном пути => красим её в чёрный
			color[v] = 3
		}
	}

	printArray(result)
}

func printArray(array []int) {
	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d", array[i])
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", array[len(array)-1])
}
