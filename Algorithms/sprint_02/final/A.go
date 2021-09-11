package main

/*
https://contest.yandex.ru/contest/22781/run-report/52647578/

-- ПРИНЦИП РАБОТЫ --
Очередь реализована с использованием кольцевого буфера, представленного массивом.
Используются два указателя, head и tail, при создании равные нулю.
При работе с хвостом очереди (back) используется указатель tail, при работе с головой очереди (front) - указатель head.
Указатель tail увеличивается "по часовой стрелке", а указатель head - уменьшается "против часовой стрелки".
При изменении указателей всегда используется деление по модулю на размер массива, чтобы поддерживать их
в диапазоне 0...(n - 1).
Очередь состоит из элементов массива d[head ... tail−1] или d[0 ... tail−1] и d[head ... n−1].
При выполнении операции мы поддерживаем переменную size, указывающую на текущее количество элементов в очереди.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
При добавлении элементов с разных сторон очереди мы перемещаем указатели по циклическому массиву навстречу друг другу.
Когда они сойдутся, то значит очередь полностью заполнена. Указатели всегда перемещаются в противоположные стороны.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Посколько мы работаем с массивом элементов в памяти с доступом по индексу, и при операциях просто перемещаем
соответствующий указатель на единицу, то все операции добавления и удаления элементов с обеих сторон очереди
выполняеются за константное время O(1).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Для хранения элементов испольются два указателя и массив постоянной длины n. Требуемый объём памяти составляет O(n).
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	head int
	tail int
	n    int
	size int
	d    []int
}

func newDeque(n int) Deque {
	return Deque{0, 0, n, 0, make([]int, n)}
}

func (deque *Deque) pushBack(x int) error {
	if deque.isFull() {
		return errors.New("overflow")
	}
	deque.d[deque.tail] = x
	deque.tail = (deque.tail + 1) % deque.n
	deque.size++
	return nil
}

func (deque *Deque) popBack() (int, error) {
	if deque.isEmpty() {
		return 0, errors.New("underflow")
	}
	deque.tail = (deque.tail - 1 + deque.n) % deque.n
	deque.size--
	return deque.d[deque.tail], nil
}

func (deque *Deque) pushFront(x int) error {
	if deque.isFull() {
		return errors.New("overflow")
	}
	deque.head = (deque.head - 1 + deque.n) % deque.n
	deque.d[deque.head] = x
	deque.size++
	return nil
}

func (deque *Deque) popFront() (int, error) {
	if deque.isEmpty() {
		return 0, errors.New("underflow")
	}
	elementToReturn := deque.d[deque.head]
	deque.head = (deque.head + 1) % deque.n
	deque.size--
	return elementToReturn, nil
}

func (deque *Deque) isEmpty() bool {
	return deque.size == 0
}

func (deque *Deque) isFull() bool {
	return deque.size == deque.n
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	var n, m int

	// читаем количество команд
	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем максимальный размер дека
	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	deque := newDeque(m)

	// читаем и выполняем команды
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "push_back") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			err := deque.pushBack(x)
			if err != nil {
				fmt.Println("error")
			}
		} else if strings.Contains(line, "push_front") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			err := deque.pushFront(x)
			if err != nil {
				fmt.Println("error")
			}
		} else if strings.Contains(line, "pop_back") {
			x, err := deque.popBack()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(x)
			}
		} else if strings.Contains(line, "pop_front") {
			x, err := deque.popFront()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(x)
			}
		}
	}
}
