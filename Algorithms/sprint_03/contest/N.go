package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	left  int
	right int
}

func cmp(interval1 interval, interval2 interval) int {
	if interval1.right < interval2.left { // interval1 левее
		return -1
	} else if interval1.left > interval2.right { // interval1 правее
		return 1
	} else { // интервалы пересекаются
		return 0
	}
}

func addInterval(intervals []interval, intervalToAdd interval) []interval {
	if len(intervals) == 0 {
		return append(intervals, intervalToAdd)
	} else {
		lastInterval := intervals[len(intervals)-1]
		if cmp(lastInterval, intervalToAdd) == 0 { // интервалы пересекаются
			union := unionIntervals(lastInterval, intervalToAdd)
			intervals[len(intervals)-1] = union
		} else {
			intervals = append(intervals, intervalToAdd)
		}
		return intervals
	}
}

func mergeSort(array []interval) []interval {
	length := len(array)
	if length == 1 {
		return array
	}

	left := mergeSort(array[0 : length/2])
	right := mergeSort(array[length/2 : length])

	var result []interval

	l, r := 0, 0

	for {
		if l >= len(left) || r >= len(right) {
			break
		}
		if cmp(left[l], right[r]) == -1 {
			result = addInterval(result, left[l])
			l += 1
		} else if cmp(left[l], right[r]) == 1 {
			result = addInterval(result, right[r])
			r += 1
		} else { // сливаем интервалы в один
			union := unionIntervals(left[l], right[r])
			result = addInterval(result, union)
			l += 1
			r += 1
		}
	}

	for {
		if l >= len(left) {
			break
		}
		result = addInterval(result, left[l])
		l += 1
	}

	for {
		if r >= len(right) {
			break
		}
		result = addInterval(result, right[r])
		r += 1
	}

	return result
}

func unionIntervals(interval1, interval2 interval) interval {
	left1 := interval1.left
	left2 := interval2.left

	right1 := interval1.right
	right2 := interval2.right

	return interval{min(left1, left2), max(right1, right2)}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество садовников
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем координаты клумб
	intervals := make([]interval, n)

	var start, end int
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		split := strings.Split(line, " ")
		start, _ = strconv.Atoi(split[0])
		end, _ = strconv.Atoi(split[1])
		intervals[i] = interval{start, end}
	}

	result := mergeSort(intervals)

	for _, i := range result {
		fmt.Println(strconv.Itoa(i.left) + " " + strconv.Itoa(i.right))
	}
}
