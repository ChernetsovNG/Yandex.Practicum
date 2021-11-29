package main

func swap(heap []int, from int, to int) {
	tmp := heap[from]
	heap[from] = heap[to]
	heap[to] = tmp
}

/*
Функция принимает в качестве аргументов массив, в котором хранятся элементы кучи, и индекс элемента,
от которого надо сделать просеивание вниз. Функция должна вернуть индекс, на котором элемент оказался
после просеивания. Также необходимо изменить порядок элементов в переданном в функцию массиве.
*/
func siftDown(heap []int, idx int) int {
	heapSize := len(heap)

	left := 2 * idx
	right := 2*idx + 1

	// нет дочерних узлов
	if heapSize < left+1 {
		return idx
	}

	var indexLargest int
	// right <= heap.size проверяет, что есть оба дочерних узла
	if right <= heapSize-1 && heap[left] < heap[right] {
		indexLargest = right
	} else {
		indexLargest = left
	}

	if heap[idx] < heap[indexLargest] {
		swap(heap, idx, indexLargest)
		return siftDown(heap, indexLargest)
	} else {
		return idx
	}
}

func test() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		panic("WA")
	}
}
