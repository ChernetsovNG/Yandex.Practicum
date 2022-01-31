package main

/*
https://contest.yandex.ru/contest/25070/run-report/63195056/

-- ПРИНЦИП РАБОТЫ --
Рассмотрим города и связывающие их железные дороги как направленный граф. По условию задачи карта железных дорог
не будет оптимальной, если найдутся какие-нибудь два города, между которыми можно будет проехать по дорогам типа R и
по дорогам типа B. Рассмотрим рёбра графа типа R как направленные "вперёд", а рёбра типа B - как направленные "назад".
Тогда между двумя вершинами в не оптимальной карте можно будет проехать слева-направо по дорогам типа R, и затем
вернуться назад по дорогам типа B. А это значит, что в построенном таким образом графе будет цикл. Таким образом,
наличие в графе цикла эквивалентно не оптимальности карты железных дорог.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Выше показано, что если между двумя городами можно проехать как по дорогам типа R, так и по дорогам типа B, то в
направленном графе с рёбрами типа R, направленными "вперёд", и с рёбрами типа B, направленными "назад", будет цикл.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Для поиска цикла мы используем алгоритм поиска в глубину. В этом алгоритме мы просматриваем все вершины графа, а для
каждой вершины - все исходящие из неё рёбра. Поэтому общее количество операций составляет O(V + E), где V - количество
вершин, а E - количество рёбер в графе.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Мы храним граф в виде списка смежности. Для каждой из n вершин храним исходящие из неё рёбра. Для этого хранения
требуется количество памяти, пропорциональное O(V + E). Для реализации алгоритма поиска в глубину мы используем
стек на основе массива. Количество элементов в стеке точно не превышает V. Таким образом, суммарная пространственная
сложность не превышает O(V + E).
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var line string

	// читаем количество городов в стране (узлов графа)
	var n int

	scanner.Scan()
	line = scanner.Text()

	n, _ = strconv.Atoi(line)

	// читаем информацию о дорогах (рёбрах графа)
	// для каждой вершины сохраняем связанные с ней вершины
	// ключ - номер вершины, значение - смежные с данной вершиной вершины
	// и тип дороги до каждой из них
	adjacencyListMap := make(map[int][]int)

	for i := 1; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		symbols := []rune(line)
		for j := 1; j <= n-i; j++ {
			roadType := symbols[j-1] // R или B
			if roadType == 'R' {     // ребро направлено "вперёд", из i в i+j
				adjacencyListMap[i] = append(adjacencyListMap[i], i+j)
			} else if roadType == 'B' { // ребро направлено "назад", из i+j в i
				adjacencyListMap[i+j] = append(adjacencyListMap[i+j], i)
			}
		}
	}

	// цвета вершин: 1 - white, 2 - gray, 3 - black
	color := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		color[i] = 1
	}

	// выполняем обход графа в глубину в поисках цикла
	hasCycle := false

	// пока остались не посещённые вершины
	dfsStack := newStack()
out:
	for i := 1; i <= n; i++ {
		if color[i] == 1 { // white
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
					adjacencyList, contains := adjacencyListMap[v]
					if contains {
						// для каждого исходящего из v ребра
						for _, w := range adjacencyList {
							if color[w] == 1 { // white
								dfsStack.push(w)
							} else if color[w] == 2 { // grey => обнаружили цикл
								hasCycle = true
								break out
							}
						}
					}
				} else if color[v] == 2 { // gray
					// серую вершину мы могли получить из стека только на обратном пути => красим её в чёрный
					color[v] = 3
				}
			}
		}
	}

	if hasCycle {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}

// Стек
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
