package main

import "fmt"

/*type Node struct {
	value int
	left  *Node
	right *Node
}*/

func getElementsInRange(root *Node, left int, right int, array *[]int) {
	if root == nil {
		return
	}
	if left <= root.value {
		getElementsInRange(root.left, left, right, array)
	}
	if left <= root.value && root.value <= right {
		*array = append(*array, root.value)
	}
	getElementsInRange(root.right, left, right, array)
}

func printRange(root *Node, left int, right int) {
	var array []int
	getElementsInRange(root, left, right, &array)

	if len(array) == 0 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("%d", array[i])
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", array[len(array)-1])
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}
