package main

type Node struct {
	value int
	left  *Node
	right *Node
	size  int
}

func safeSize(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.size
	}
}

/*
Функция принимает корень дерева и число k, а возвращает два BST — в первом k
наименьших элементов из исходного дерева, а во втором — оставшиеся вершины BST
*/
func split(root *Node, k int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}
	if k == 0 {
		return nil, root
	}

	leftSize := safeSize(root.left)
	var splitLeft, splitRight *Node

	if leftSize+1 <= k {
		splitLeft, splitRight = split(root.right, k-leftSize-1)
		root.right = splitLeft
		root.size = safeSize(root.left) + safeSize(root.right) + 1
		return root, splitRight
	} else {
		splitLeft, splitRight = split(root.left, k)
		root.left = splitRight
		root.size = safeSize(root.left) + safeSize(root.right) + 1
		return splitLeft, root
	}
}

func test() {
	node1 := &Node{3, nil, nil, 1}
	node2 := &Node{2, nil, node1, 2}
	node3 := &Node{8, nil, nil, 1}
	node4 := &Node{11, nil, nil, 1}
	node5 := &Node{10, node3, node4, 3}
	node6 := &Node{5, node2, node5, 6}
	left, right := split(node6, 1)
	if left.size != 1 {
		panic("WA")
	}
	if right.size != 5 {
		panic("WA")
	}
}
