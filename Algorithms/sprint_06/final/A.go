package main

/*
https://contest.yandex.ru/contest/25070/run-report/61890495/

-- ПРИНЦИП РАБОТЫ --
Мы используем алгоритм Прима поиска остовного дерева с максимальным весом рёбер.
Алгоритму не важно, с какой вершины начинать, так как в итоге все вершины попадут в остовное дерево. Мы берём
вершину с номером 1.
Далее рассматриваем все рёбра, выходящие из этой вершины. Берём ребро с максимальным весом, и добавляем в остовное
дерево вершину, в которую оно ведёт.
Добавляем ко множеству потенциально добавляемых рёбер все, которые исходят из новой вершины и входят в вершины, ещё
не включённые в остовное дерево.
Повторяем эти действия до тех пор, пока в остовное дерево не будут добавлены все вершины графа.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
На самом деле неочевидно :) В лекции оно больше апеллирует к интуиции читателя

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Первоначально добавляем все вершины в множество не добавленных, за O(n).
Далее добавляем каждую вершину в остовное дерево, и на каждом шаге выбираем исходящее из неё ребро
максимального веса. Для хранения рёбер мы используем пирамиду с поддержанием максимума.
В лекции говорится, что для такого случая сложность алгоритма составляет O(E * log(V)), где E - количество рёбер,
а V - количество вершин в графе. Доказательство этого выражения также неочевидно.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Мы храним граф в виде списка смежности. Для каждой из n вершин храним исходящие из неё рёбра. Для этого хранения
требуется количество памяти, пропорциональное O(V + E). Затем мы храним множества добавленных и не добавленных вершин,
сумма размеров этих множеств равна O(V). В пирамиде с поддержанием максимума мы храним рёбра, исходящие из вершин,
добавленных в остовное дерево. Пирамида реализована на основе массива. Размер массива задан перед началом вычислений,
и равен O(E). Наконец, в массиве мы храним рёбра, составляющие само остовное дерево. Размер его не определён, но
ограничен сверху значением O(E). Итоговая пространственная сложность составляет O(V + E + V + E + E) = O(2*V + 3*E) =
O(V + E).
*/

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

	// читаем число достопримечательностей и мостов
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
	adjacencyListMap := make(map[int][]Edge)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		uvw := strings.Split(line, " ")
		u, _ := strconv.Atoi(uvw[0])
		v, _ := strconv.Atoi(uvw[1])
		w, _ := strconv.Atoi(uvw[2])

		adjacencyList, contains := adjacencyListMap[u]
		if !contains {
			adjacencyListMap[u] = []Edge{newEdge(v, w)}
		} else {
			adjacencyListMap[u] = append(adjacencyList, newEdge(v, w))
		}

		adjacencyList, contains = adjacencyListMap[v]
		if !contains {
			adjacencyListMap[v] = []Edge{newEdge(u, w)}
		} else {
			adjacencyListMap[v] = append(adjacencyList, newEdge(u, w))
		}
	}

	// находим максимальное остовное дерево (точнее, сумму весов составляющих его рёбер)
	var maximumSpanningTree []Edge

	// Множество вершин, уже добавленных в остов
	added := makeSet()

	// Множество вершины, ещё не добавленных в остов
	notAdded := makeSet()

	// Массив рёбер, исходящих из остовного дерева
	edges := newHeap(m)

	// первоначально все вершины не добавлены
	for i := 1; i <= n; i++ {
		notAdded.Add(i)
	}

	// берём первую попавшуюся вершину (допустим, первую) и добавляем её в остовное дерево
	v := 1

	addVertex(v, &adjacencyListMap, added, notAdded, &edges)

	// пока notAdded не пуст и edges не пуст
	for {
		if notAdded.IsEmpty() || edges.isEmpty() {
			break
		}
		// извлекаем ребро с максимальным весом из массива рёбер
		maxEdge := edges.popMax()

		if notAdded.Exists(maxEdge.to) {
			maximumSpanningTree = append(maximumSpanningTree, maxEdge)
			addVertex(maxEdge.to, &adjacencyListMap, added, notAdded, &edges)
		}
	}

	if !notAdded.IsEmpty() {
		fmt.Print("Oops! I did it again")
	} else {
		// подсчитываем сумму длин рёбер максимального остовного дерева
		weight := 0
		for _, edge := range maximumSpanningTree {
			weight += edge.weight
		}
		fmt.Print(weight)
	}
}

func addVertex(v int, graph *map[int][]Edge, added *Set, notAdded *Set, edges *MaxHeap) {
	added.Add(v)
	notAdded.Remove(v)

	// Добавляем все рёбра, которые инцидентны v, но их конец ещё не в остовном дереве
	adjacencyList, contains := (*graph)[v]
	if contains {
		for _, e := range adjacencyList {
			to := e.to
			if notAdded.Exists(to) {
				(*edges).heapAdd(newEdge(to, e.weight))
			}
		}
	}
}

// Тип для ребра
type Edge struct {
	to     int
	weight int
}

func newEdge(to, weight int) Edge {
	return Edge{to, weight}
}

// Тип для множества
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

func (c *Set) Remove(key int) {
	delete(c.container, key)
}

func (c *Set) IsEmpty() bool {
	return len(c.container) == 0
}

// Пирамида с поддержанием максимума
type MaxHeap struct {
	array   []Edge
	maxSize int // максимальный размер пирамиды
	size    int // текущий размер пирамиды
}

func newHeap(maxSize int) MaxHeap {
	return MaxHeap{make([]Edge, maxSize), maxSize, 0}
}

func (heap *MaxHeap) swap(from int, to int) {
	tmp := heap.array[from]
	heap.array[from] = heap.array[to]
	heap.array[to] = tmp
}

func (heap *MaxHeap) siftUp(idx int) int {
	if idx == 1 {
		return 1
	}
	parentIndex := idx / 2
	if heap.array[parentIndex].weight < heap.array[idx].weight {
		heap.swap(parentIndex, idx)
		return heap.siftUp(parentIndex)
	} else {
		return idx
	}
}

func (heap *MaxHeap) siftDown(idx int) int {
	left := 2 * idx
	right := 2*idx + 1

	// нет дочерних узлов
	if heap.size < left+1 {
		return idx
	}

	var indexLargest int
	// right <= heap.size проверяет, что есть оба дочерних узла
	if right <= heap.size-1 && heap.array[left].weight < heap.array[right].weight {
		indexLargest = right
	} else {
		indexLargest = left
	}

	if heap.array[idx].weight < heap.array[indexLargest].weight {
		heap.swap(idx, indexLargest)
		return heap.siftDown(indexLargest)
	} else {
		return idx
	}
}

func (heap *MaxHeap) heapAdd(key Edge) {
	index := heap.size + 1
	heap.array[index] = key
	heap.siftUp(index)
	heap.size += 1
}

func (heap *MaxHeap) popMax() Edge {
	result := heap.array[1]
	heap.array[1] = heap.array[heap.size]
	heap.siftDown(1)
	heap.size -= 1
	return result
}

func (heap *MaxHeap) isEmpty() bool {
	return heap.size == 0
}
