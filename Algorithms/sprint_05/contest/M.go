package main

func swap(heap []int, from int, to int) {
	tmp := heap[from]
	heap[from] = heap[to]
	heap[to] = tmp
}

/*
Функция принимает в качестве аргументов массив, в котором хранятся элементы кучи, и индекс элемента,
от которого надо сделать просеивание вверх. Функция должна вернуть индекс, на котором элемент оказался
после просеивания. Также необходимо изменить порядок элементов в переданном в функцию массиве.
*/
func siftUp(heap []int, idx int) int {
	if idx == 1 {
		return 1
	}
	parentIndex := idx / 2
	if heap[parentIndex] < heap[idx] {
		swap(heap, parentIndex, idx)
		return siftUp(heap, parentIndex)
	} else {
		return idx
	}
}

func test() {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	if siftUp(sample, 5) != 1 {
		panic("WA")
	}
}
