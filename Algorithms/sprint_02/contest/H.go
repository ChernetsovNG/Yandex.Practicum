package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	array []rune
}

func (stack *Stack) push(x rune) {
	stack.array = append(stack.array, x)
}

func (stack *Stack) pop() {
	stack.array = stack.array[:len(stack.array)-1]
}

func (stack *Stack) top() rune {
	return stack.array[len(stack.array)-1]
}

func (stack *Stack) size() int {
	return len(stack.array)
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()

	var stack Stack

	// идём по скобкам в строке
	symbols := []rune(line)

	for i := 0; i < len(symbols); i++ {
		char := symbols[i]
		if char == ')' {
			if stack.size() > 0 && stack.top() == '(' {
				stack.pop()
			} else {
				stack.push(char)
			}
		} else if char == ']' {
			if stack.size() > 0 && stack.top() == '[' {
				stack.pop()
			} else {
				stack.push(char)
			}
		} else if char == '}' {
			if stack.size() > 0 && stack.top() == '{' {
				stack.pop()
			} else {
				stack.push(char)
			}
		} else {
			stack.push(char)
		}
	}

	if stack.size() == 0 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}
