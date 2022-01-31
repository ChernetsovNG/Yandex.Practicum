package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	var line string

	// читаем входное число
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	if n <= 4 {
		if n == 1 || n == 4 {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
		return
	}

	if n%4 != 0 {
		fmt.Println("False")
		return
	}

	var rem int
	for {
		rem = n % 4
		n /= 4
		if n == 4 && rem == 0 {
			fmt.Println("True")
			break
		}
		if n < 4 {
			fmt.Println("False")
			break
		}
	}
}
