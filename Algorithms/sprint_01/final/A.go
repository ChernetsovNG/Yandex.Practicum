package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/run-report/52379682/
func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1_000_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем длину улицы
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	numbers := make([]int, n) // номера домов
	result := make([]int, n)  // массив для результата

	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")

	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(values[i])
	}

	// два указателя
	left := 0
	right := n - 1

	var distLeft, distRight int
	for {
		if left >= n && right < 0 {
			break
		}

		if numbers[left] != 0 {
			distLeft = distance(left, n, numbers, result, 1)
			if distLeft != 0 { // удалось посчитать
				if result[left] == 0 { // ещё не было посчитано
					result[left] = distLeft
				} else { // уже было посчитано
					if distLeft < result[left] {
						result[left] = distLeft
					}
				}
			}
		}
		if numbers[right] != 0 {
			distRight = distance(right, n, numbers, result, -1)
			if distRight != 0 { // удалось посчитать
				if result[right] == 0 { // ещё не было посчитано
					result[right] = distRight
				} else { // уже было посчитано
					if distRight < result[right] {
						result[right] = distRight
					}
				}
			}
		}

		left++
		right--
	}

	// выводим результат
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n-1; i++ {
		writer.WriteString(strconv.Itoa(result[i]))
		writer.WriteString(" ")
	}
	writer.WriteString(strconv.Itoa(result[n-1]))
	writer.Flush()
}

func distance(index int, n int, numbers []int, distances []int, direction int) int {
	if direction == 1 { // движение слева-направо
		if index == 0 {
			return 0
		}
		if numbers[index-1] == 0 { // участок слева - пустой
			return 1
		} else { // участок слева - непустой
			if distances[index-1] != 0 { // для участка слева было посчитано расстояние
				return distances[index-1] + 1
			} else {
				return 0 // пока не можем посчитать
			}
		}
	} else if direction == -1 { // движение справа-налево
		if index == n-1 {
			return 0
		}
		if numbers[index+1] == 0 { // участок справа - пустой
			return 1
		} else { // участок справа - непустой
			if distances[index+1] != 0 {
				return distances[index+1] + 1
			} else {
				return 0 // пока не можем посчитать
			}
		}
	}
	return 0
}
