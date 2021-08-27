package main

func FindIndex(head *ListNode, elem string, idx int) int {
	if head == nil {
		return -1
	} else if head.data == elem {
		return idx
	} else {
		return FindIndex(head.next, elem, idx+1)
	}
}

func Solution_D(head *ListNode, elem string) int {
	return FindIndex(head, elem, 0)
}
