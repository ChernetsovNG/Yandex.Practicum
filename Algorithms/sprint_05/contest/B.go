package main

/*
type Node struct {
	value int
	left  *Node
	right *Node
}
*/

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func isBalanced(root *Node, depth int) (bool, int) {
	if root == nil {
		return true, -1
	}

	leftBalanced, leftHeight := isBalanced(root.left, depth+1)
	rightBalanced, rightHeight := isBalanced(root.right, depth+1)

	balanced := Abs(leftHeight-rightHeight) <= 1
	subtreesAreBalanced := leftBalanced && rightBalanced
	height := Max(leftHeight, rightHeight) + 1

	return balanced && subtreesAreBalanced, height
}

func Solution(root *Node) bool {
	balanced, _ := isBalanced(root, 0)
	return balanced
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}
}
