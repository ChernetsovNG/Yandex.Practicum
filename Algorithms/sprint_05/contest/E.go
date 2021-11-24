package main

import "math"

/*type Node struct {
	value int
	left  *Node
	right *Node
}*/

func CheckTreeInRange(root *Node, lo, hi int) bool {
	if lo > hi {
		panic("lo > hi")
	}
	if root == nil {
		return true
	}
	if root.value <= lo || root.value >= hi {
		return false
	}
	return CheckTreeInRange(root.left, lo, root.value) && CheckTreeInRange(root.right, root.value, hi)
}

func Solution(root *Node) bool {
	return CheckTreeInRange(root, math.MinInt32, math.MaxInt32)
}
