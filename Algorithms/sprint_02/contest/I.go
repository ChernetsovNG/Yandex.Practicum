package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QueueSized struct {
	head  int
	tail  int
	size  int
	maxN  int
	queue []int
}

func newQueueSized(n int) QueueSized {
	return QueueSized{0, 0, 0, n, make([]int, n)}
}

func (queueSized *QueueSized) push(x int) error {
	if queueSized.size != queueSized.maxN {
		queueSized.queue[queueSized.tail] = x
		queueSized.tail = (queueSized.tail + 1) % queueSized.maxN
		queueSized.size += 1
		return nil
	} else {
		return errors.New("queue is full")
	}
}

func (queueSized *QueueSized) pop() (int, error) {
	if queueSized.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	x := queueSized.queue[queueSized.head]
	queueSized.head = (queueSized.head + 1) % queueSized.maxN
	queueSized.size -= 1
	return x, nil
}

func (queueSized *QueueSized) peek() (int, error) {
	if queueSized.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	return queueSized.queue[queueSized.head], nil
}

func (queueSized *QueueSized) isEmpty() bool {
	return queueSized.size == 0
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	var n, maxN int

	// читаем количество команд
	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем максимальный размер очереди
	scanner.Scan()
	line = scanner.Text()
	maxN, _ = strconv.Atoi(line)

	queue := newQueueSized(maxN)

	// читаем и выполняем команды
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "push") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			err := queue.push(x)
			if err != nil {
				fmt.Println("error")
			}
		} else if strings.Contains(line, "pop") {
			x, err := queue.pop()
			if err != nil {
				fmt.Println("None")
			} else {
				fmt.Println(x)
			}
		} else if strings.Contains(line, "peek") {
			x, err := queue.peek()
			if err != nil {
				fmt.Println("None")
			} else {
				fmt.Println(x)
			}
		} else if strings.Contains(line, "size") {
			fmt.Println(queue.size)
		}
	}
}
