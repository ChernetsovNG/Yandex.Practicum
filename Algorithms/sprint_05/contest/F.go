package main

/*type Node struct {
	value int
	left  *Node
	right *Node
}*/

func Solution(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := Solution(root.left)
	rightDepth := Solution(root.right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	if Solution(&node5) != 3 {
		panic("WA")
	}
}
