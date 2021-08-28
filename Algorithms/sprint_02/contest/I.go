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

func (queue *QueueSized) push(x int) error {
	if queue.size != queue.maxN {
		queue.queue[queue.tail] = x
		queue.tail = (queue.tail + 1) % queue.maxN
		queue.size += 1
		return nil
	} else {
		return errors.New("queue is full")
	}
}

func (queue *QueueSized) pop() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	x := queue.queue[queue.head]
	queue.head = (queue.head + 1) % queue.maxN
	queue.size -= 1
	return x, nil
}

func (queue *QueueSized) peek() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	return queue.queue[queue.head], nil
}

func (queue *QueueSized) isEmpty() bool {
	return queue.size == 0
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
