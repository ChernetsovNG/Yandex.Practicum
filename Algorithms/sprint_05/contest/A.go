package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func GetMax(node *Node, max *int) int {
	if node.value > *max {
		*max = node.value
	}
	if node.left != nil {
		maxLeft := GetMax(node.left, max)
		if maxLeft > *max {
			*max = maxLeft
		}
	}
	if node.right != nil {
		maxRight := GetMax(node.right, max)
		if maxRight > *max {
			*max = maxRight
		}
	}
	return *max
}

func Solution(root *Node) int {
	max := root.value
	if root.left != nil {
		maxLeft := GetMax(root.left, &max)
		if maxLeft > max {
			max = maxLeft
		}
	}
	if root.right != nil {
		maxRight := GetMax(root.right, &max)
		if maxRight > max {
			max = maxRight
		}
	}
	return max
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	fmt.Print(Solution(&node4))
}

func main() {
	test()
}
