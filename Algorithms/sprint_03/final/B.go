package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type participant struct {
	login string
	p     int
	f     int
}

/*
при сравнении двух участников выше будет идти тот, у которого решено больше задач. При равенстве
числа решённых задач первым идёт участник с меньшим штрафом. Если же и штрафы совпадают, то первым
будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке
*/
func cmp(p1 participant, p2 participant) int {
	if p1.p > p2.p {
		return 1
	} else if p1.p < p2.p {
		return -1
	} else { // число решённых задач одинаково
		if p1.f < p2.f {
			return 1
		} else if p1.f > p2.f {
			return -1
		} else { // штрафы одинаковые
			if p1.login < p2.login {
				return 1
			} else if p1.login > p2.login {
				return -1
			} else { // равны все параметры участников
				return 0
			}
		}
	}
}

/*
   less = элементы array, меньшие pivot
   center = элементы array, равные pivot
   greater = элементы array, большие pivot
   return less, center, greater
*/
func partition(arr []participant, pivot participant) ([]participant, []participant, []participant) {
	var less, center, greater []participant

	var compareResult int
	for _, p := range arr {
		compareResult = cmp(p, pivot)
		if compareResult == -1 {
			less = append(less, p)
		} else if compareResult == 0 {
			center = append(center, p)
		} else if compareResult == 1 {
			greater = append(greater, p)
		}
	}

	return less, center, greater
}

func quicksort(arr []participant) []participant {
	if len(arr) < 2 {
		return arr
	} else {
		pivot, _ := medianOf3(arr)
		less, center, greater := partition(arr, pivot)
		return merge(quicksort(less), center, quicksort(greater))
	}
}

func merge(less, center, greater []participant) []participant {
	result := make([]participant, len(less)+len(center)+len(greater))
	k := 0
	for i := 0; i < len(less); i++ {
		result[k] = less[i]
		k += 1
	}
	for i := 0; i < len(center); i++ {
		result[k] = center[i]
		k += 1
	}
	for i := 0; i < len(greater); i++ {
		result[k] = greater[i]
		k += 1
	}
	return result
}

func medianOf3(arr []participant) (participant, error) {
	length := len(arr)
	if length == 0 {
		return participant{}, errors.New("empty array")
	} else if length == 1 {
		return arr[0], nil
	} else if length == 2 {
		return arr[0], nil
	}
	first := arr[0]
	middle := arr[length/2]
	last := arr[length-1]
	if cmp(first, middle) == -1 && cmp(middle, last) == -1 {
		return middle, nil
	} else if cmp(middle, last) == -1 && cmp(last, first) == -1 {
		return last, nil
	} else if cmp(middle, first) == -1 && cmp(first, last) == -1 {
		return first, nil
	}
	return first, nil
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 10_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество участников
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем данные по участникам
	participants := make([]participant, n)

	var login string
	var p, f int

	for i := 0; i < n; i++ {
		scanner.Scan()
		row := scanner.Text()
		values := strings.Split(row, " ")
		login = values[0]
		p, _ = strconv.Atoi(values[1])
		f, _ = strconv.Atoi(values[2])
		participants[i] = participant{login, p, f}
	}

	// отсортированы по возрастанию
	sorted := quicksort(participants)

	// выводим в обратном порядке: от лучшего к худшему
	for i := len(sorted) - 1; i >= 0; i-- {
		fmt.Println(sorted[i].login)
	}
}
