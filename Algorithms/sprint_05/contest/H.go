package main

import "strconv"

/*type Node struct {
	value int
	left  *Node
	right *Node
}*/

func Path(root *Node, acc string, sum *int) {
	valueString := strconv.Itoa(root.value)
	// дошли до листа
	if root.left == nil && root.right == nil {
		addToSum, _ := strconv.Atoi(acc + valueString)
		*sum += addToSum
		return
	} else {
		// добавляем к аккумулятору значение и спускаемся ниже
		if root.left != nil {
			Path(root.left, acc+valueString, sum)
		}
		if root.right != nil {
			Path(root.right, acc+valueString, sum)
		}
	}
}

func Solution(root *Node) int {
	var result = 0
	Path(root, "", &result)
	return result
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, nil, nil}
	node5 := Node{1, &node4, &node3}
	if Solution(&node5) != 275 {
		panic("WA")
	}
}
