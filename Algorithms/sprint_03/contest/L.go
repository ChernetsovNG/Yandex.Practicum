package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func positionBinarySearch(arr []int, x, left, right int) int {
	if right <= left {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] >= x {
		if mid > 0 {
			if arr[mid-1] < x {
				return mid
			} else if arr[mid-1] == x {
				return mid - 1
			} else {
				return positionBinarySearch(arr, x, left, mid)
			}
		} else {
			return mid
		}
	} else if x < arr[mid] {
		return positionBinarySearch(arr, x, left, mid)
	} else {
		return positionBinarySearch(arr, x, mid+1, right)
	}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем количество дней
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем количество денег в копилке
	arr := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		arr[i] = value
	}

	// читаем стоимость велосипеда
	var s int

	scanner.Scan()
	line = scanner.Text()
	s, _ = strconv.Atoi(line)

	// вычисляем ответ
	var k1, k2 int

	k1 = positionBinarySearch(arr, s, 0, n)
	k2 = positionBinarySearch(arr, s*2, 0, n)

	if k1 != -1 {
		k1 += 1
	}
	if k2 != -1 {
		k2 += 1
	}

	fmt.Print(strconv.Itoa(k1) + " " + strconv.Itoa(k2))
}
