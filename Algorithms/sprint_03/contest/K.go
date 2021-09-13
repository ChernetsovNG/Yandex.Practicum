package main

/*
Функция merge_sort принимает некоторый подмассив, который нужно отсортировать.
Подмассив задаётся полуинтервалом — его началом и концом. Функция должна отсортировать
передаваемый в неё подмассив, она ничего не возвращает.

Функция merge_sort разбивает полуинтервал на две половинки и рекурсивно вызывает сортировку отдельно
для каждой. Затем два отсортированных массива сливаются в один с помощью merge.
*/
func merge_sort(arr []int, lf int, rg int) {
	if rg-lf == 1 {
		return
	}
	mid := (lf + rg) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)
	result := merge(arr, lf, mid, rg)
	for i := lf; i < rg; i++ {
		arr[i] = result[i-lf]
	}
}

/*
Функция merge принимает два отсортированных массива, сливает их в один отсортированный массив и возвращает его.
Первый массив задаётся полуинтервалом [left, mid) массива array, а второй - полуинтервалом [mid, right) массива
array
*/
func merge(arr []int, left int, mid int, right int) (result []int) {
	leftArray := arr[left:mid]
	rightArray := arr[mid:right]

	result = make([]int, right-left)
	l, r, k := 0, 0, 0
	for true {
		if l >= len(leftArray) || r >= len(rightArray) {
			break
		}
		if leftArray[l] <= rightArray[r] {
			result[k] = leftArray[l]
			l += 1
		} else {
			result[k] = rightArray[r]
			r += 1
		}
		k += 1
	}

	for true {
		if l >= len(leftArray) {
			break
		}
		result[k] = leftArray[l]
		l += 1
		k += 1
	}

	for true {
		if r >= len(rightArray) {
			break
		}
		result[k] = rightArray[r]
		r += 1
		k += 1
	}

	return result
}

/*func main() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}*/
