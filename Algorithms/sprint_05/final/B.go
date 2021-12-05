package main

/*
https://contest.yandex.ru/contest/24810/run-report/60088795/

-- ПРИНЦИП РАБОТЫ --
Алгоритм работает рекурсивно. В зависимости от значения ключа возможны 3 случая:
- удаление нужно выполнить в левом поддереве (1)
- удаление нужно выполнить в правом поддереве (2)
- нужно удалить корень дерева (3)
Случаи (1) и (2) в итоге сводятся к (3).

Если при удалении корня у нас есть только или левое, или правое поддерево, то просто возвращаем соответствующее
поддерево в качестве решения на шаге рекурсии.

Если же есть оба поддерева, то нужно удалить корень так, чтобы дерево не распалось на 2 и не нарушилась
его целостность. Для этого мы находим минимальный узел в правом поддереве (самый левый узел), и перемещаем
его на место корня.

При этом возможны 2 варианта:
- минимальный узел является листом дерева.
В этом случае просто перемещаем его на место корня

- минимальный узел имеет поддерево справа.
Тогда мы переподвешиваем это поддерево к родителю минимального узла,
а затем перемещаем минимальный узел на место корня.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Мы пользуемся свойством двоичного дерева поиска, что в нём для любого узла все узлы в левом поддереве
меньше, а все узлы в правом поддереве - больше этого узла. В этом случае, если на место удаляемого
корня дерева переставить минимальный узел в правом поддереве, то свойство дерева не нарушится.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Вначале мы за какое количество шагов спускаемся до удаляемого узла, и рассматриваем поддерево
с корнем в этом узле. Затем рекурсивно находим в правом поддереве этого узла минимальный узел.

Затем за константное время выполняем переподвешивание узлов.

Количество "спусков" в дереве не превышает высоты дерева в худшем случае. А она, в свою очередь, не превышает
количества узлов в дереве (n) в худшем случае (для несбалансированного дерева). Для "нормального" дерева она в среднем
будет составлять O(log n).

Таким образом, в худшем случае алгоритм работает за O(n), в среднем случает - за O(log n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Каждый узел представлен хранимым в узле значением и двумя ссылками - на левое и правое поддерево.
Так что нужно хранить n значений и 2*n ссылок. Таким образом, сложность составляет O(n).
*/

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node *Node) isLeaf() bool {
	return node.left == nil && node.right == nil
}

// Находит самую левую вершину и её родителя в правом поддереве
// Вход - корень правого поддерева и его родитель
func findLeftestNodeAndParentInRightSubtree(root *Node, parent *Node) (*Node, *Node) {
	if root.left == nil {
		return root, parent
	} else {
		return findLeftestNodeAndParentInRightSubtree(root.left, root)
	}
}

func remove(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	if key < root.value {
		// удаляем узел в левом поддереве
		root.left = remove(root.left, key)
	} else if key > root.value {
		// удаляем узел в правом поддереве
		root.right = remove(root.right, key)
	} else {
		// удаляем корень

		// если есть только один ребёнок (или нет детей)
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		// если есть оба ребёнка
		// выбираем минимальный (самый левый) узел в правом поддереве
		leftestNode, leftestNodeParent := findLeftestNodeAndParentInRightSubtree(root.right, root)

		if leftestNode.isLeaf() {
			// самый левый узел в правом поддереве - лист. Переставляем его на место корня
			if leftestNodeParent != root {
				leftestNodeParent.left = nil
			}

			leftestNode.left = root.left
			if root.right != leftestNode {
				leftestNode.right = root.right
			}

			root = leftestNode
		} else {
			// самый левый узел в правом поддереве - имеет правого ребёнка. Переподвешиваем этого ребёнка
			leftestNodeParent.left = leftestNode.right

			// переставляем самый левый узел в правом поддереве на место корня
			leftestNode.left = root.left
			leftestNode.right = root.right

			root = leftestNode
		}
	}
	return root
}

func test() {
	node6 := Node{5, nil, nil}
	node7 := Node{7, nil, nil}

	node4 := Node{1, nil, nil}
	node5 := Node{3, nil, nil}

	node2 := Node{2, &node4, &node5}
	node3 := Node{6, &node6, &node7}

	node1 := Node{4, &node2, &node3}

	newHead := remove(&node1, 2)

	fmt.Println(newHead.value)
}