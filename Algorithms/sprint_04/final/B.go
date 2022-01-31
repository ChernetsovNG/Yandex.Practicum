package main

/*
// https://contest.yandex.ru/contest/24414/run-report/53768518/

-- ПРИНЦИП РАБОТЫ --
Реализована хеш-таблица с разрешением коллизий методом цепочек. Таблица представляет собой массив бакетов, каждый
из которых представляет собой связный список из элементов, имеющих одинаковый хеш. При добавлении элемента в таблицу
вычисляет хеш от ключа, на основе которого определяется индекс бакета по формуле hash(key) mod m, где m - размер
таблицы. Далее элемент добавляется в связный список данного бакета. Если список пустой, то элемент добавляется в голову
списка, если непустой, то ищется элемент с таким же ключом. Если такой элемент найден, то его значение перезаписывается
новым значением. Если такой элемент не найден, то в голову списка добавляется новый узел.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Вычисление хеш-кода от ключа занимает O(1), т.к. в нашем случае хеш-код от целого числа равен самому этому числу.
Так что перед нами классическая хеш-таблица с амортизированным временем выполнения всех операций за O(1).

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Все операции выполняются за амортизированное время O(1). Возможны специальные случаи, когда все добавляемые элементы
имеют одинаковый хеш-код. Например, 1, 1+m, 1+2*m, ..., 1+n*m, где m - размер хеш-таблицы. Т.к. индекс бакета
определяется путём взятия остатка от деления на m, то все элементы с такими ключами попадут в один бакет с индексом 1,
в результате чего в нём образуется длинный список элементов. И в таком, худшем случае, время выполнения операций
составит O(n). В случае же более-менее равномерного распределения элементов сложность будет составлять O(1).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Мы храним массив бакетов - указателей на соответствующие связные списки. При добавлении n элементов каждый из них
попадёт в тот или иной бакет, и для него будет создан узел списка. Поэтому пространственная сложность составит O(n).
*/

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Хеш-таблица
type HashTable struct {
	buckets []*List
	m       int
}

func newHashTable(m int) HashTable {
	return HashTable{make([]*List, m), m}
}

/*
Получение значения по ключу. Если ключа нет в таблице, то вывести «None».
Иначе вывести найденное значение
*/
func (table *HashTable) get(key int) (int, error) {
	bucket := table.buckets[hash(key)%table.m]
	if bucket == nil {
		return 0, errors.New("key not found")
	}
	return bucket.get(key)
}

/*
Добавление пары ключ-значение. Если заданный ключ уже есть в таблице,
то соответствующее ему значение обновляется
*/
func (table *HashTable) put(key, value int) {
	bucketIndex := hash(key) % table.m
	if table.buckets[bucketIndex] == nil {
		table.buckets[bucketIndex] = &List{nil}
	}
	table.buckets[bucketIndex].put(key, value)
}

/*
Удаление ключа из таблицы. Если такого ключа нет, то вывести «None», иначе вывести хранимое
по данному ключу значение и удалить ключ.
*/
func (table *HashTable) delete(key int) (int, error) {
	bucket := table.buckets[hash(key)%table.m]
	if bucket == nil {
		return 0, errors.New("key not found")
	}
	return bucket.delete(key)
}

// простой хеш-код для целого числа, когда код равен самому числу
func hash(key int) int {
	return key
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	writer := bufio.NewWriter(os.Stdout)

	var line string

	// читаем количество запросов к таблице
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// создаём новую хеш таблицу
	table := newHashTable(1_000)

	// читаем и обрабатываем запросы
	// читаем и выполняем команды
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "get") {
			split := strings.Split(line, " ")
			key, _ := strconv.Atoi(split[1])
			value, err := table.get(key)
			if err != nil {
				writer.WriteString("None\n")
			} else {
				writer.WriteString(strconv.Itoa(value))
				writer.WriteString("\n")
			}
		} else if strings.Contains(line, "put") {
			split := strings.Split(line, " ")
			key, _ := strconv.Atoi(split[1])
			value, _ := strconv.Atoi(split[2])
			table.put(key, value)
		} else if strings.Contains(line, "delete") {
			split := strings.Split(line, " ")
			key, _ := strconv.Atoi(split[1])
			value, err := table.delete(key)
			if err != nil {
				writer.WriteString("None\n")
			} else {
				writer.WriteString(strconv.Itoa(value))
				writer.WriteString("\n")
			}
		}
	}
	writer.Flush()
}

// Бакеты для хеш-таблицы на основе связных списков
type Node struct {
	key, value int
	next       *Node
}

type List struct {
	head *Node
}

func (list *List) put(key int, value int) {
	if list.head == nil {
		list.head = &Node{key, value, nil}
	} else {
		node := list.head
		for {
			if node == nil {
				break
			}
			if node.key == key {
				node.value = value // обновляем значение в узле
				return
			}
			node = node.next
		}
		// если ничего не нашли, то добавляем новый элемент в голову списка
		node = &Node{key, value, nil}
		node.next = list.head
		list.head = node
	}
}

func (list *List) get(key int) (int, error) {
	if list.head == nil {
		return 0, errors.New("key not found")
	}
	node := list.head
	for {
		if node == nil {
			return 0, errors.New("key not found")
		}
		if node.key == key {
			return node.value, nil
		}
		node = node.next
	}
	return 0, errors.New("key not found")
}

func (list *List) delete(key int) (int, error) {
	if list.head == nil {
		return 0, errors.New("key not found")
	}
	var prevNode *Node = nil
	node := list.head
	for {
		if node == nil {
			return 0, errors.New("key not found")
		}
		if node.key == key {
			if prevNode == nil { // удаляем голову списка
				list.head = node.next
			} else {
				prevNode.next = node.next
			}
			return node.value, nil
		}
		prevNode = node
		node = node.next
	}
	return 0, errors.New("key not found")
}
