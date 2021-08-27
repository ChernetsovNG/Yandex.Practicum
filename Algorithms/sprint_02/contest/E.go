package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	current := head

	for true {
		prev := current.prev
		next := current.next

		if next == nil {
			return current
		} else {
			if prev != nil {
				prev.prev = current
			}
			current.prev = next

			next.next = current
			next.prev = nil

			current.next = prev

			current = current.prev
		}
	}
	return nil
}

func test() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0

	newHead := Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
	fmt.Print(newHead)
}

func main() {
	test()
}
