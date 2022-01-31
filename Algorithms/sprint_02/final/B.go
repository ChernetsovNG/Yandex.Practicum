package main

/*
// https://contest.yandex.ru/contest/22781/run-report/52645127/
-- ПРИНЦИП РАБОТЫ --
-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Используется алгоритм из задания

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Операции со стеком (добавление и извлечение элемента) выполняются за амортизированное константное время,
т.к. для реализации стека используется динамический массив

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Для хранения элементов (операторов и операндов) в стеке используется массив. Максимальный размер массива в стеке
будет порядка количества символов (за исключением пробелов) во входной строке
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	array []string
}

func (stack *Stack) push(x string) {
	stack.array = append(stack.array, x)
}

func (stack *Stack) pop() string {
	valueToReturn := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return valueToReturn
}

func (stack *Stack) top() string {
	return stack.array[len(stack.array)-1]
}

func (stack *Stack) size() int {
	return len(stack.array)
}

// математическое целочисленное деление (взято из java: Math.floorDiv)
func floorDiv(x, y int) int {
	r := x / y
	if x^y < 0 && r*y != x {
		r--
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	scanner.Scan()
	line := scanner.Text()

	var stack Stack

	symbols := strings.Split(line, " ")
	n := len(symbols)

	for i := 0; i < n; i++ {
		symbol := symbols[i]
		// на вход подан знак операции
		if symbol == "+" || symbol == "-" || symbol == "*" || symbol == "/" {
			// в нашем случае все операции бинарные, поэтому из стека извлекаются два операнда
			operand2String := stack.pop() // берём элементы в порядке добавления
			operand1String := stack.pop()
			operand1, _ := strconv.Atoi(operand1String)
			operand2, _ := strconv.Atoi(operand2String)
			var result int
			if symbol == "+" {
				result = operand1 + operand2
			} else if symbol == "-" {
				result = operand1 - operand2
			} else if symbol == "*" {
				result = operand1 * operand2
			} else if symbol == "/" {
				result = floorDiv(operand1, operand2)
			}
			stack.push(strconv.Itoa(result))
		} else { // на вход подан операнд (число)
			stack.push(symbol)
		}
	}

	fmt.Println(stack.top())
}
