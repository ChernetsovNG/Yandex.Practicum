package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Узел списка
type ListNode struct {
	key, value int
	next       *ListNode
}

// Связный список
type LinkedList struct {
	head *ListNode
}

func newLinkedList() LinkedList {
	return LinkedList{}
}

// добавление пары "ключ - значение" в связный список
func (list *LinkedList) add(key, value int) {
	node := ListNode{key, value, nil}
	if list.head == nil { // новый список
		list.head = &node
	} else {
		prevHead := list.head
		list.head = &node
		node.next = prevHead
	}
}

// поиск в связном списке значения по ключу
func (list *LinkedList) get(key int) (int, error) {
	node := list.head
	if node == nil {
		return 0, errors.New("list is empty")
	}
	for true {
		if node.key == key {
			return node.value, nil
		}
		node = node.next
		if node == nil { // дошли до конца списка
			return 0, errors.New("key not found")
		}
		if node.next == nil { // дошли до конца списка
			return 0, errors.New("key not found")
		}
	}
	return 0, errors.New("key not found")
}

func (list *LinkedList) getNode(key int) *ListNode {
	node := list.head
	if node == nil {
		return nil
	}
	for true {
		if node.key == key {
			return node
		}
		node = node.next
		if node == nil { // дошли до конца списка
			return nil
		}
		if node.next == nil { // дошли до конца списка
			return nil
		}
	}
	return nil
}

func (list *LinkedList) getNodeValueAndRemoveNode(key int) (int, error) {
	node := list.head
	if node == nil {
		return 0, errors.New("key not found")
	}
	if node.key == key { // нужно удалить голову списка
		value := list.head.value
		list.head = node.next
		return value, nil
	} else {
		for true {
			if node.next == nil { // дошли до конца списка
				return 0, errors.New("key not found")
			} else if node.next.key == key { // нужно удалить следующий узел
				value := node.next.value
				node.next = node.next.next
				return value, nil
			}
		}
	}
	return 0, errors.New("key not found")
}

type Bucket struct {
	values *LinkedList
}

func newBucket() Bucket {
	return Bucket{}
}

type HashTable struct {
	buckets []*Bucket
	m       int
}

func newHashTable(m int) HashTable {
	return HashTable{make([]*Bucket, m), m}
}

/*
Получение значения по ключу. Если ключа нет в таблице, то вывести «None».
Иначе вывести найденное значение
*/
func (table *HashTable) get(key int) (int, error) {
	h := hash(key)
	bucketNumber := h % table.m
	bucket := table.buckets[bucketNumber]
	if bucket == nil {
		return 0, errors.New("key not found")
	}
	return bucket.values.get(key)
}

/*
Добавление пары ключ-значение. Если заданный ключ уже есть в таблице,
то соответствующее ему значение обновляется
*/
func (table *HashTable) put(key, value int) {
	bucketNumber := hash(key) % table.m
	if table.buckets[bucketNumber] == nil {
		linkedList := newLinkedList()
		bucket := newBucket()
		bucket.values = &linkedList
		table.buckets[bucketNumber] = &bucket
	}
	// перед добавлением проверим, не было ли уже в списке такой пары "ключ-значение"
	bucket := table.buckets[bucketNumber]
	node := bucket.values.getNode(key)
	if node == nil { // нет узла с таким ключом
		bucket.values.add(key, value)
	} else {
		node.value = value // по ключу уже было значение => обновляем его
	}
}

/*
Удаление ключа из таблицы. Если такого ключа нет, то вывести «None», иначе вывести хранимое
по данному ключу значение и удалить ключ.
*/
func (table *HashTable) delete(key int) (int, error) {
	bucket := table.buckets[hash(key) % table.m]
	if bucket == nil {
		return 0, errors.New("key not found")
	}
	return bucket.values.getNodeValueAndRemoveNode(key)
}

// простой хеш-код для целого числа, когда код равен самому числу
func hash(key int) int {
	return key
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество запросов к таблице
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// создаём новую хеш таблицу
	table := newHashTable(10_000)

	// читаем и обрабатываем запросы
	// читаем и выполняем команды
	var results []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()

		if strings.Contains(line, "get") {
			split := strings.Split(line, " ")
			key, _ := strconv.Atoi(split[1])
			value, err := table.get(key)
			if err != nil {
				results = append(results, "None")
			} else {
				results = append(results, strconv.Itoa(value))
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
				results = append(results, "None")
			} else {
				results = append(results, strconv.Itoa(value))
			}
		}
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}
