package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	// читаем допустимые к использованию слова и добавляем их в множество (для быстрой проверки вхождения)
	words := makeSet()

	// запоминаем длину самого длинного слова из словаря
	maxLen := -1

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		words.add(line)
		if len(line) > maxLen {
			maxLen = len(line)
		}
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

	for i := 1; i <= len(text); i++ {
		if words.contains(string(textSymbols[0:i])) {
			dp[i] = true
			continue
		}
		for j := i - maxLen; j <= i-1; j++ {
			if j >= 1 && dp[j] == true && words.contains(string(textSymbols[j:i])) {
				dp[i] = true
				break
			}
		}
	}

	if dp[len(text)] == true {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// множество строк
func makeSet() *Set {
	return &Set{
		container: make(map[string]bool),
	}
}

type Set struct {
	container map[string]bool
}

func (c *Set) contains(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *Set) add(key string) {
	c.container[key] = true
}
