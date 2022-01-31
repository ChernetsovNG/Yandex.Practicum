package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type QueueSized struct {
	head  int
	tail  int
	size  int
	maxN  int
	queue []int
}

func newQueueSized(n int) QueueSized {
	return QueueSized{0, 0, 0, n, make([]int, n)}
}

func (queue *QueueSized) push(x int) error {
	if queue.size != queue.maxN {
		queue.queue[queue.tail] = x
		queue.tail = (queue.tail + 1) % queue.maxN
		queue.size += 1
		return nil
	} else {
		return errors.New("queue is full")
	}
}

func (queue *QueueSized) pop() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	x := queue.queue[queue.head]
	queue.head = (queue.head + 1) % queue.maxN
	queue.size -= 1
	return x, nil
}

func (queue *QueueSized) peek() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	return queue.queue[queue.head], nil
}

func (queue *QueueSized) isEmpty() bool {
	return queue.size == 0
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

	// сортируем списки смежности вершин в порядке возрастания
	for _, adjacencyList := range adjacencyListMap {
		sort.Ints(adjacencyList)
	}

	// читаем информацию о стартовой вершине
	scanner.Scan()
	line = scanner.Text()
	s, _ := strconv.Atoi(line)

	// цвета вершин: 1 - white, 2 - gray, 3 - black
	color := make([]int, n+1)
	previous := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		color[i] = 1
		previous[i] = -1
	}

	// создадим очередь вершин и положим туда стартовую вершину
	planned := newQueueSized(n)
	planned.push(s)
	color[s] = 2 // gray

	// массив вершин в порядке обхода
	var result []int

	for {
		if planned.isEmpty() {
			break
		}
		// пока очередь не пуста
		u, _ := planned.pop() // возьмём вершину из очереди
		result = append(result, u)

		adjacencyList, _ := adjacencyListMap[u]
		// для каждого ребра (u,v), исходящего из u
		// возьмём вершину v
		for _, v := range adjacencyList {
			if color[v] == 1 { // white
				// серые и чёрные вершины уже либо в очереди, либо обработаны
				previous[v] = u
				color[v] = 2    // grey
				planned.push(v) // запланируем посещение вершины
			}
		}
		color[u] = 3 // black, теперь вершина считается обработанной
	}

	printArray(result)
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
