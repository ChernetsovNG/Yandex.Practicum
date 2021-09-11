package main

import (
	"bufio"
	"fmt"
	"os"
)

func generate(input []rune, n int, index int, prefix string, result *[]string) {
	if n == 0 {
		*result = append(*result, prefix)
		return
	}
	symbol := input[index]
	if symbol == '2' {
		generate(input, n-1, index+1, prefix+"a", result)
		generate(input, n-1, index+1, prefix+"b", result)
		generate(input, n-1, index+1, prefix+"c", result)
	} else if symbol == '3' {
		generate(input, n-1, index+1, prefix+"d", result)
		generate(input, n-1, index+1, prefix+"e", result)
		generate(input, n-1, index+1, prefix+"f", result)
	} else if symbol == '4' {
		generate(input, n-1, index+1, prefix+"g", result)
		generate(input, n-1, index+1, prefix+"h", result)
		generate(input, n-1, index+1, prefix+"i", result)
	} else if symbol == '5' {
		generate(input, n-1, index+1, prefix+"j", result)
		generate(input, n-1, index+1, prefix+"k", result)
		generate(input, n-1, index+1, prefix+"l", result)
	} else if symbol == '6' {
		generate(input, n-1, index+1, prefix+"m", result)
		generate(input, n-1, index+1, prefix+"n", result)
		generate(input, n-1, index+1, prefix+"o", result)
	} else if symbol == '7' {
		generate(input, n-1, index+1, prefix+"p", result)
		generate(input, n-1, index+1, prefix+"q", result)
		generate(input, n-1, index+1, prefix+"r", result)
		generate(input, n-1, index+1, prefix+"s", result)
	} else if symbol == '8' {
		generate(input, n-1, index+1, prefix+"t", result)
		generate(input, n-1, index+1, prefix+"u", result)
		generate(input, n-1, index+1, prefix+"v", result)
	} else if symbol == '9' {
		generate(input, n-1, index+1, prefix+"w", result)
		generate(input, n-1, index+1, prefix+"x", result)
		generate(input, n-1, index+1, prefix+"y", result)
		generate(input, n-1, index+1, prefix+"z", result)
	}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	// читаем входную строку из цифр 2-9
	scanner.Scan()
	line := scanner.Text()

	n := len(line)
	symbols := []rune(line)

	var result []string

	generate(symbols, n, 0, "", &result)

	count := len(result)
	for i := 0; i < count-1; i++ {
		fmt.Print(result[i] + " ")
	}
	fmt.Print(result[count-1])
}
