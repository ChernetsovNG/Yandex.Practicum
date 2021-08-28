package main

/*func getNodeByIndex(node *ListNode, index int) *ListNode {
	for true {
		if index <= 0 {
			break
		}
		node = node.next
		index--
	}
	return node
}

func Solution_C(head *ListNode, idx int) *ListNode {
	if idx == 0 { // удаляемый элемент - первый
		return head.next
	} else {
		// предыдущий элемент
		prevElement := getNodeByIndex(head, idx-1)

		// удаляемый элемент
		currElement := prevElement.next

		// следующий элемент
		nextElement := currElement.next

		if nextElement == nil { // удаляемый элемент - последний
			prevElement.next = nil
		} else { // удаляемый элемент - в середине списка
			prevElement.next = nextElement
		}

		return head
	}
}*/
