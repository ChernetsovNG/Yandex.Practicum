package main

/*
https://contest.yandex.ru/contest/24810/run-report/59993200/

-- ПРИНЦИП РАБОТЫ --
Мы создаём бинарную пирамиду, в которой максимальный элемент хранится в вершине, и которая обладает
свойством, что каждый дочерний элемент не больше родительского. При добавлении и удалении новых
элементов этот инвариант поддерживается при помощи "просеивания" элементов вверх и вниз.

При добавлении нового элемента он добавляется в конец пирамиды, а затем "всплывает" вверх до тех пор, пока
свойство пирамиды не будет восстановлено.

При удалении максимального элемента из вершины пирамиды последний элемент помещается на его место, и затем
он "тонет" вниз до тех пор, пока свойство пирамиды не будет восстановлено.

Мы помещаем всех участников соревнования в пирамиду одного за другим. Затем удаляем текущего максимального
участника из вершины пирамиды и восстанавливаем её свойство. Каждый раз мы извлекаем из вершины текущего
максимального участника, таким образом получая всех участников, отсортированных по убыванию.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Свойство пирамиды гарантирует нам, что на каждом шаге в её вершине находится лучший участник из всех, добавленных
в неё. Первоначально в неё добавлены все участники. Каждый раз извлекая очередного участника, мы получаем всех
участников, отсортированных по убыванию.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Вначале мы выделяем массив для n участников за O(1).

Затем вставляем n элементов подряд в пирамиду. Сложность этого этапа:
O(log 1) + O(log 2) + ... + O(log n)
Это значение ограничено сверху:
O(log n) + O(log n) + ... + O(log n) = O(n * log n)

Затем мы извлекаем по очереди n элементов. Сложность этой операции
также ограничена сверху значением O(n * log n).

Таким образом, общая временная сложность алгоритма составляет O(n * log n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В алгоритме используется массив для хранения n элементов пирамиды.
Пространственная сложность составляет O(n).
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type participant struct {
	login string
	p     int
	f     int
}

/*
при сравнении двух участников выше будет идти тот, у которого решено больше задач. При равенстве
числа решённых задач первым идёт участник с меньшим штрафом. Если же и штрафы совпадают, то первым
будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке
*/
func cmp(p1 participant, p2 participant) int {
	if p1.p > p2.p {
		return 1
	} else if p1.p < p2.p {
		return -1
	} else { // число решённых задач одинаково
		if p1.f < p2.f {
			return 1
		} else if p1.f > p2.f {
			return -1
		} else { // штрафы одинаковые
			if p1.login < p2.login {
				return 1
			} else if p1.login > p2.login {
				return -1
			} else { // равны все параметры участников
				return 0
			}
		}
	}
}

type MaxHeap struct {
	array   []participant
	maxSize int // максимальный размер пирамиды
	size    int // текущий размер пирамиды
}

func newHeap(maxSize int) MaxHeap {
	return MaxHeap{make([]participant, maxSize), maxSize, 0}
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
	if cmp(heap.array[parentIndex], heap.array[idx]) < 0 {
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
	if right <= heap.size-1 && cmp(heap.array[left], heap.array[right]) < 0 {
		indexLargest = right
	} else {
		indexLargest = left
	}

	if cmp(heap.array[idx], heap.array[indexLargest]) < 0 {
		heap.swap(idx, indexLargest)
		return heap.siftDown(indexLargest)
	} else {
		return idx
	}
}

func (heap *MaxHeap) heapAdd(key participant) {
	index := heap.size + 1
	heap.array[index] = key
	heap.siftUp(index)
	heap.size += 1
}

func (heap *MaxHeap) popMax() participant {
	result := heap.array[1]
	heap.array[1] = heap.array[heap.size]
	heap.siftDown(1)
	heap.size -= 1
	return result
}

func heapsort(a []participant) []participant {
	// Создадим пустую пирамиду
	heap := newHeap(len(a) + 1)

	// Вставим в неё по одному все элементы массива, сохраняя свойства пирамиды
	for i := 0; i < len(a); i++ {
		participant := a[i]
		heap.heapAdd(participant)
	}

	sortedArray := make([]participant, 0, len(a))

	// Будем извлекать из пирамиды наиболее приоритетные элементы
	for {
		if heap.size == 0 {
			break
		}
		sortedArray = append(sortedArray, heap.popMax())
	}

	return sortedArray
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество участников
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем данные по участникам
	participants := make([]participant, n)

	var login string
	var p, f int

	for i := 0; i < n; i++ {
		scanner.Scan()
		row := scanner.Text()
		values := strings.Split(row, " ")
		login = values[0]
		p, _ = strconv.Atoi(values[1])
		f, _ = strconv.Atoi(values[2])
		participants[i] = participant{login, p, f}
	}

	// сортируем участников по убыванию при помощи max-пирамиды
	sorted := heapsort(participants)

	// выводим участников: от лучшего к худшему
	for i := 0; i < len(sorted); i++ {
		fmt.Println(sorted[i].login)
	}
}
