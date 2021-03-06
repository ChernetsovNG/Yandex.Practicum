package main

/*
https://contest.yandex.ru/contest/26133/run-report/64693406/

-- ПРИНЦИП РАБОТЫ --

Для определения возможности "сегментации" текста на слова из заданного словаря используем динамическое
программирование со следующей динамикой:
dp[i] — возможно ли разбить на слова из заданного словаря текст, оканчивающийся в i-ом символе. Заполнив
этот массив значениями true/false для всех символов, определим окончательный ответ как dp[n], где n - количество
символов в тексте.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Базовый случай:
dp[0] = true, т.к. пустой текст можно набрать, если не брать ни одно слово из словаря

Переход динамики:
dp[i] = true, если для какого-нибудь j от 1 до i-1 dp[j] = true и при этом символы [j+1 ... i]
являются словом из словаря. dp[j] = true означает, что из словаря можно набрать текст с символами [1...j],
а, так как символы [j+1 ... i] являются словом из словаря, то текст из символов [1 ... i] можно набрать из
слов словаря.

Находясь в i-ом символе, можно быстро проверить, не содержится ли в предыдущих символах слово из словаря,
с использованием префиксного дерева, построенного на словах из словаря, но записанных в обратном порядке.
Отступая назад от символа i, и подавая очередной символ в префиксное дерево, мы проверяем текущий узел. Если он
оказывается терминальным, то значит одно из слов словаря предшествует символу i.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

Во внешнем цикле мы движемся по символам текста (которых n штук). На каждой итерации этого цикла мы обходим
префиксное дерево, построенное на словах из словаря. В худшем случае время этого обхода не будет больше, чем
размер самого длинного слова из словаря. Тогда временная сложность не превысит O(n*k), где k - длина самого
длинного слова из словаря.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

Мы храним массив символов размера n, а также префиксное дерево, количество узлов которого в худшем случае равно
количество символов в словах словаря (обозначим его через m). Кроме того, для реализации алгоритма динамического
программирования, хранится массив промежуточных ответов размера n. Суммарная пространственная сложность составляет
O(n + m + n) = O(n + m).
*/

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
}

func newNode() *Node {
	return &Node{false, make([]*Node, 26)}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 8 * 100000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	// читаем текст, который надо разбить на слова
	scanner.Scan()
	text := scanner.Text()

	// читаем число допустимых к использованию слов
	var n int

	scanner.Scan()
	line := scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем допустимые к использованию слова и добавляем их в префиксное дерево
	prefixTree := newNode()
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		addStringReverse(prefixTree, line)
	}

	/*
	 dp[i] — возможно ли разбить на слова из заданного словаря текст, оканчивающийся в i-ом символе
	*/
	dp := make([]bool, len(text)+1)

	// базовый случай:
	// dp[0] = true, пустой текст можно набрать, если не брать ни одно слово из словаря
	dp[0] = true

	// переход динамики:
	// dp[i] = true, если для какого-нибудь j от 0 до i-1 dp[j] = true
	// и символы [j+1 ... i-1] являются словом из словаря
	textSymbols := []rune(text)

	for i := 1; i <= len(text); i++ {
		j := i
		currentNode := prefixTree

		for {
			j -= 1
			if j < 0 {
				break
			}
			pos := position(textSymbols[j])

			currentNode = currentNode.edges[pos]
			if currentNode == nil {
				break
			}

			if dp[j] && currentNode != nil && currentNode.isTerminal {
				dp[i] = true
				break
			}
		}
	}

	if dp[len(text)] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// функция построения префиксного дерева:
// добавляем новую строку в дерево посимвольно, обходя её в обратном порядке
func addStringReverse(root *Node, s string) {
	currentNode := root
	runes := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		symbol := runes[i]
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
}

// ord(a) = 97
func position(r rune) int {
	return int(r) - 97
}
