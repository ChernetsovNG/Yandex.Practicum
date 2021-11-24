package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func printLMR(root *Node) {
	if root.left != nil {
		printLMR(root.left)
	}
	fmt.Printf("%d ", root.value)
	if root.right != nil {
		printLMR(root.right)
	}
}

func doInsert(root *Node, key int) {
	if key < root.value {
		if root.left == nil {
			root.left = &Node{key, nil, nil}
			return
		} else {
			doInsert(root.left, key)
		}
	} else {
		if root.right == nil {
			root.right = &Node{key, nil, nil}
			return
		} else {
			doInsert(root.right, key)
		}
	}
}

func insert(root *Node, key int) *Node {
	doInsert(root, key)
	return root
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	printLMR(&node3)
	fmt.Println()
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		panic("WA")
	}
	if newHead.left.value != 6 {
		panic("WA")
	}
	printLMR(&node3)
}

func main() {
	test()
}
