package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 4 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	// читаем два числа
	scanner.Scan()
	number1 := scanner.Text()

	scanner.Scan()
	number2 := scanner.Text()

	n1Len := len(number1)
	n2Len := len(number2)

	var sb strings.Builder
	// дополняем более короткое число слева нулями
	if n1Len < n2Len {
		delta := n2Len - n1Len
		for i := 0; i < delta; i++ {
			sb.WriteRune('0')
		}
		zerosString := sb.String()
		number1 = zerosString + number1
	} else if n2Len < n1Len {
		delta := n1Len - n2Len
		for i := 0; i < delta; i++ {
			sb.WriteRune('0')
		}
		zerosString := sb.String()
		number2 = zerosString + number2
	}

	// Теперь числа одинаковой длины. Идём справа налево и складываем их
	number1Array := []rune(number1)
	number2Array := []rune(number2)

	var result []string
	addition := false // признак переноса из предыдущего разряда
	for i := len(number1) - 1; i >= 0; i-- {
		char1 := number1Array[i]
		char2 := number2Array[i]

		if char1 == '0' && char2 == '0' {
			if !addition {
				result = append(result, "0")
			} else if addition {
				result = append(result, "1")
				addition = false
			}
		} else if char1 == '0' && char2 == '1' {
			if !addition {
				result = append(result, "1")
			} else if addition {
				result = append(result, "0")
			}
		} else if char1 == '1' && char2 == '0' {
			if !addition {
				result = append(result, "1")
			} else if addition {
				result = append(result, "0")
			}
		} else if char1 == '1' && char2 == '1' {
			if !addition {
				result = append(result, "0")
				addition = true
			} else if addition {
				result = append(result, "1")
			}
		}
	}
	if addition {
		result = append(result, "1")
	}

	// Печатаем результат
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}
