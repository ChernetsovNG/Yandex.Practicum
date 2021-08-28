package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackMax struct {
	array    []int
	arrayMax []int
}

func (stack *StackMax) push(x int) {
	stack.array = append(stack.array, x)
	if len(stack.arrayMax) == 0 {
		stack.arrayMax = append(stack.arrayMax, x)
	} else {
		currentMax := stack.arrayMax[len(stack.arrayMax)-1]
		if x > currentMax { // обновляем текущий максимум
			stack.arrayMax = append(stack.arrayMax, x)
		} else {
			stack.arrayMax = append(stack.arrayMax, currentMax)
		}
	}
}

func (stack *StackMax) pop() {
	if len(stack.array) == 0 {
		fmt.Println("error")
	} else {
		stack.array = stack.array[:len(stack.array)-1]
		stack.arrayMax = stack.arrayMax[:len(stack.arrayMax)-1]
	}
}

func (stack *StackMax) getMax() {
	if len(stack.arrayMax) == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(stack.arrayMax[len(stack.arrayMax)-1])
	}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество команд
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	var stack StackMax
	// читаем и выполняем команды
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "push") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			stack.push(x)
		} else if strings.Contains(line, "pop") {
			stack.pop()
		} else if strings.Contains(line, "get_max") {
			stack.getMax()
		}
	}
}
