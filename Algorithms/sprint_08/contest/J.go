package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	isTerminal bool
	// массив переходов:
	// на позиции с индексом i хранится ссылка на ребро, отвечающее i-му символу алфавита
	edges []*Node
	// в терминальных узлах будем хранить слова, которые их образуют
	strings []string
}

func newNode() *Node {
	edges := make([]*Node, 26)
	for i := 0; i < 26; i++ {
		edges[i] = nil
	}
	return &Node{false, edges, []string{}}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество названий классов в исходном наборе
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем названия классов и строим префиксное дерево
	prefixTree := newNode()
	// запоминаем в массиве все названия классов
	var classes []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		addString(prefixTree, line)
		classes = append(classes, line)
	}
	sort.Strings(classes)

	// читаем количество запросов m
	var m int

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем запросы (шаблоны)
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		// шаблону из нуля прописных букв удовлетворяет любое название
		if len(line) == 0 {
			for i := 0; i < n; i++ {
				fmt.Println(classes[i])
			}
		} else {
			// ищем удовлетворяющие шаблону названия классов в префиксном дереве
			strings := findNode(prefixTree, line)
			// Если ни одна из строк не подходит под шаблон, то выведите для данного запроса пустую строку
			if len(strings) == 0 {
				fmt.Println("")
			} else {
				sort.Strings(strings)
				for _, s := range strings {
					fmt.Println(s)
				}
			}
		}
	}

}

// функция построения префиксного дерева - добавления новой строки в дерево
func addString(root *Node, s string) {
	currentNode := root

	// в дереве храним только uppercase-символы
	for _, symbol := range s {
		if isUppercase(symbol) {
			pos := position(symbol)
			// если из currentNode нет перехода по символу:
			if currentNode.edges[pos] == nil {
				// создать узел newNode
				// создать ребро symbol из currentNode в newNode
				currentNode.edges[pos] = newNode()
			}
			// сдвинуться на следующий символ
			currentNode = currentNode.edges[pos]
		}
	}

	// пометить currentNode как терминальный символ
	currentNode.isTerminal = true
	currentNode.strings = append(currentNode.strings, s)
}

// функция поиска слов (названий классов) в префиксном дереве по шаблону
func findNode(root *Node, s string) []string {
	currentNode := root
	var result []string
	for _, symbol := range s {
		// если из current_node нет перехода по symbol
		pos := position(symbol)
		if currentNode.edges[pos] == nil {
			return result
		}
		// сдвинуться на следующий символ
		currentNode = currentNode.edges[pos]
	}
	// в качестве ответа выводим строки в текущем узле (если он терминальный),
	// и во всех его терминальных узлах-потомках
	findAllTerminalNodesWords(currentNode, &result)

	return result
}

func findAllTerminalNodesWords(root *Node, strings *[]string) {
	if root.isTerminal {
		for _, s := range root.strings {
			*strings = append(*strings, s)
		}
	}
	for _, edge := range root.edges {
		if edge != nil {
			findAllTerminalNodesWords(edge, strings)
		}
	}
}

// ord(A) = 65
func position(r rune) int {
	return int(r) - 65
}

func isUppercase(r rune) bool {
	ord := int(r)
	return 65 <= ord && ord <= 90
}
