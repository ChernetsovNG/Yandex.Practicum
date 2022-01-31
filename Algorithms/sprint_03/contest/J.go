package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	var line string

	// читаем длину массива
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем значения в массиве
	arr := make([]int, n)

	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(values[i])
		arr[i] = value
	}

	// выполняем пузырьковую сортировку массива
	changesCount := 0 // количество обменов
	cycleCount := 0   // общее количество сортировок

	var tmp int
out:
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				changesCount++
			}
		}
		if changesCount == 0 { // массив уже отсортирован
			break out
		} else {
			fmt.Println(strSlice(arr))
			changesCount = 0
			cycleCount += 1
		}
	}

	if cycleCount == 0 { // значит массив уже был отсортирован
		fmt.Println(strSlice(arr))
	}
}

func strSlice(arr []int) string {
	var b bytes.Buffer
	for _, i := range arr {
		fmt.Fprintf(&b, "%d ", i)
	}
	return b.String()
}
