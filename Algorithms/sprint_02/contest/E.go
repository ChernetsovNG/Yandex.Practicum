package main

/*func Solution(head *ListNode) *ListNode {
	var temp *ListNode = nil
	current := head

	for true {
		if current == nil {
			break
		}
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		head = temp.prev
	}

	return head
}*/
