package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
		adjacencyListMap[v] = append(adjacencyListMap[v], u)
	}

	color := make([]int, n+1)
	for i := 0; i <= n; i++ {
		color[i] = -1
	}

	isBipartite := true
out:
	for i := 1; i <= n; i++ {
		if color[i] == -1 {
			color[i] = 1

			_, contains := adjacencyListMap[i]
			if !contains {
				continue
			}

			queue := newQueueSized(n)
			queue.push(i)

			for {
				if queue.isEmpty() {
					break
				}
				u, _ := queue.pop()
				adjacencyList, _ := adjacencyListMap[u]
				for _, v := range adjacencyList {
					// находим все неокрашенные смежные вершины
					if color[v] == -1 {
						// присваиваем вершине противоположный цвет (0 => 1, 1 => 0)
						color[v] = 1 - color[u]
						queue.push(v)
					} else if color[v] == color[u] {
						// иначе если вершины окрашены в один и тот же цвет, то граф не двудольный
						isBipartite = false
						break out
					}
				}
			}
		}
	}

	if isBipartite {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
