package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество чисел
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем числа
	numbers := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		numbers[i] = value
	}

	/*
	  AB ABY ==> ABABY vs ABYAB
	  AB > ABY if A > Y
	*/
	sort.Slice(numbers, func(i, j int) bool {
		str1 := strconv.Itoa(numbers[i])
		str2 := strconv.Itoa(numbers[j])

		len1 := len(str1)
		len2 := len(str2)

		var toLen int
		if len1 == len2 {
			toLen = len1
		} else {
			toLen = len1 + len2
		}

		var ac, bc uint8
		for i := 0; i < toLen; i++ {
			if i < len1 {
				ac = str1[i]
			} else {
				ac = str2[i-len1]
			}

			if i < len2 {
				bc = str2[i]
			} else {
				bc = str1[i-len2]
			}

			if ac == bc {
				continue
			}
			if ac > bc {
				return true
			} else {
				return false
			}
		}

		return false
	})

	var b bytes.Buffer
	for _, i := range numbers {
		fmt.Fprintf(&b, "%d", i)
	}
	fmt.Print(b.String())
}
