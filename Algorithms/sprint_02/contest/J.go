package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	data int
	prev *ListNode
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (list *LinkedList) popTail() *ListNode {
	tail := list.tail
	prev := tail.prev
	if prev != nil {
		prev.next = nil
	}
	list.tail = prev
	if list.tail == nil { // если удалили из списка все элементы
		list.head = nil
	}
	return tail
}

func (list *LinkedList) add(x int) {
	node := ListNode{x, nil, nil}
	if list.head == nil { // новый список
		list.head = &node
		list.tail = &node
	} else {
		prevHead := list.head
		list.head = &node
		node.next = prevHead
		prevHead.prev = &node
	}
}

type LinkedListQueue struct {
	list LinkedList
	size int
}

func newLinkedListQueue() LinkedListQueue {
	return LinkedListQueue{LinkedList{nil, nil}, 0}
}

// Вывести элемент, находящийся в голове очереди, и удалить его. Если очередь пуста, то вывести «error»
func (queue *LinkedListQueue) get() (int, error) {
	if queue.size == 0 {
		return 0, errors.New("queue is empty")
	} else {
		tail := queue.list.popTail()
		queue.size--
		return tail.data, nil
	}
}

// добавить число x в очередь
func (queue *LinkedListQueue) put(x int) {
	queue.list.add(x)
	queue.size++
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	var n int

	// читаем количество команд
	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	queue := newLinkedListQueue()
	// читаем и выполняем команды
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "put") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			queue.put(x)
		} else if strings.Contains(line, "get") {
			x, err := queue.get()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(x)
			}
		} else if strings.Contains(line, "size") {
			fmt.Println(queue.size)
		}
	}
}
