package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/22450/run-report/52379715/
func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line, gameField string

	// читаем k - на сколько клавиш может нажать каждый игрок в некоторый момент времени
	var k int

	scanner.Scan()
	line = scanner.Text()
	k, _ = strconv.Atoi(line)

	// читаем вид тренажёра - поле размера 4 х 4
	for i := 1; i <= 4; i++ {
		scanner.Scan()
		gameField += scanner.Text()
	}
	lineSymbols := []rune(gameField)

	// сохраним в виде словаря, сколько на игровом поле каждой цифры
	numbersCountMap := make(map[rune]int)
	for i := 0; i < len(lineSymbols); i++ {
		symbol := lineSymbols[i]
		if symbol == '.' {
			continue
		}
		numbersCountMap[symbol] += 1
	}

	// идём по времени от 1 до 9
	result := 0
	for t := '1'; t <= '9'; t++ {
		symbolsCount, contains := numbersCountMap[t]
		if contains && symbolsCount <= 2*k {
			result++
		}
	}

	// выводим результат
	fmt.Print(result)
}
