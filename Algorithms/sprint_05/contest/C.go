package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func isSymmetric(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil && node1.value == node2.value {
		return isSymmetric(node1.left, node2.right) && isSymmetric(node1.right, node2.left)
	}
	return false
}

func Solution(root *Node) bool {
	return isSymmetric(root, root)
}

func test() {
	node1 := Node{3, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{4, nil, nil}
	node4 := Node{3, nil, nil}
	node5 := Node{2, &node1, &node2}
	node6 := Node{2, &node3, &node4}
	node7 := Node{1, &node5, &node6}

	if !Solution(&node7) {
		panic("WA")
	}
}
