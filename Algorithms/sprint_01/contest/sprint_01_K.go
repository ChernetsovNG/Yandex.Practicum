package main

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

	// длина списочной формы первого числа
	var _ int

	scanner.Scan()
	line = scanner.Text()
	_, _ = strconv.Atoi(line)

	// списочная форма первого числа
	// читаем значения температуры в каждый из дней
	scanner.Scan()
	line = scanner.Text()
	number1Digits := strings.Split(line, " ")

	// второе число
	var k int

	scanner.Scan()
	line = scanner.Text()
	k, _ = strconv.Atoi(line)

	number1String := strings.Join(number1Digits, "")
	number1, _ := strconv.Atoi(number1String)

	result := number1 + k

	// преобразуем число в списочную форму
	resultString := strconv.Itoa(result)
	symbols := []rune(resultString)

	for i := 0; i < len(symbols)-1; i++ {
		symbol := symbols[i]
		fmt.Printf("%c", symbol)
		fmt.Print(" ")
	}
	fmt.Printf("%c", symbols[len(symbols)-1])
}
