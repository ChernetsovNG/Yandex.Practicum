package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	isTerminal bool
	// массив переходов:
	// на позиции с индексом i хранится ссылка на ребро, отвечающее i-му символу алфавита
	edges []*Node
	// в терминальных узлах будем хранить слова, которые их образуют
	s string
}

func newNode() *Node {
	edges := make([]*Node, 26)
	return &Node{false, edges, ""}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем текст, который надо разбить на слова
	scanner.Scan()
	text := scanner.Text()

	// читаем число допустимых к использованию слов
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем допустимые к использованию слова и добавляем их в префиксное дерево
	prefixTree := newNode()

	// запоминаем длину самого длинного слова из словаря
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		addString(prefixTree, line)
	}

	/*
	 dp[i] — возможно ли разбить на слова из заданного словаря текст, оканчивающийся в i-ом символе
	*/
	dp := make([]bool, len(text)+1)

	// базовый случай:
	// dp[0] = true, пустой текст можно набрать, если не брать никакое слово из словаря
	dp[0] = true

	// переход динамики:
	// dp[i] = true, если символы от 1 до i являются словом из словаря,
	// или если для любого j от 1 до i-1 dp[j] = true и символы [j+1 ... i] являются словом из словаря
	textSymbols := []rune(text)

	for l := 1; l <= len(text); l++ {
		for i := 0; i < l; i++ {
			if dp[i] {
				currentNode := prefixTree
				for j := i; j < l; j++ {
					pos := position(textSymbols[j])
					currentNode = currentNode.edges[pos]
					if currentNode == nil {
						break
					}
				}
				if currentNode != nil && currentNode.isTerminal {
					dp[l] = true
					break
				}
			}
		}
	}

	if dp[len(text)] == true {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// функция построения префиксного дерева - добавления нового символа в дерево
func addString(root *Node, s string) {
	currentNode := root
	for _, symbol := range s {
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

	// пометить currentNode как терминальный символ
	currentNode.isTerminal = true
	currentNode.s = s
}

// ord(a) = 97
func position(r rune) int {
	return int(r) - 97
}
