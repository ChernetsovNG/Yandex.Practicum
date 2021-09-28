package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	i, j int
}

func newPair(i, j int) Pair {
	return Pair{i, j}
}

type Quad struct {
	a, b, c, d int
}

func newQuad(a, b, c, d int) Quad {
	return Quad{a, b, c, d}
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	const maxCapacity = 32 * 1_000
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем общее количество элементов в массиве
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем сумму S, для которой мы будем находить четвёрки
	var s int

	scanner.Scan()
	line = scanner.Text()
	s, _ = strconv.Atoi(line)

	// читаем массив чисел
	x := make([]int, n)

	var value int
	scanner.Scan()
	row := scanner.Text()
	values := strings.Split(row, " ")
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		x[i] = value
	}

	// отсортируем исходный массив
	sort.Ints(x)

	// ключ - сумма двух элементов, значение - позиции элементов
	sumPairsMap := make(map[int][]Pair)

	var sum int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum = x[i] + x[j]
			pairs, ok := sumPairsMap[sum]
			if !ok {
				sumPairsMap[sum] = []Pair{newPair(i, j)}
			} else {
				sumPairsMap[sum] = append(pairs, newPair(i, j))
			}
		}
	}

	sums := make([]int, len(sumPairsMap))
	i := 0
	for k := range sumPairsMap {
		sums[i] = k
		i++
	}

	quadSet := makeQuadSet()
	for i := 0; i < len(sums); i++ {
		sum := sums[i]
		addition := s - sum
		pairs2, ok := sumPairsMap[addition]
		if ok {
			pairs1 := sumPairsMap[sum]
			for i1 := 0; i1 < len(pairs1); i1++ {
				for i2 := 0; i2 < len(pairs2); i2++ {
					pair1 := pairs1[i1]
					pair2 := pairs2[i2]
					if checkAllNumbersAreDifferent(pair1.i, pair1.j, pair2.i, pair2.j) {
						a := x[pair1.i]
						b := x[pair1.j]
						c := x[pair2.i]
						d := x[pair2.j]
						a, b, c, d = sortFour(a, b, c, d)
						quadSet.Add(newQuad(a, b, c, d))
					}
				}
			}
		}
	}

	quads := quadSet.GetArray()

	// Выводим количество найденных четвёрок
	fmt.Println(len(quads))

	// Выводим сами найденные четвёрки, упорядоченные лексикографически
	sort.Slice(quads, func(i, j int) bool {
		quad1 := quads[i]
		quad2 := quads[j]
		if quad1.a < quad2.a {
			return true
		} else if quad1.a > quad2.a {
			return false
		} else {
			if quad1.b < quad2.b {
				return true
			} else if quad1.b > quad2.b {
				return false
			} else {
				if quad1.c < quad2.c {
					return true
				} else if quad1.c > quad2.c {
					return false
				} else {
					if quad1.d < quad2.d {
						return true
					} else if quad1.d > quad2.d {
						return false
					} else {
						return false
					}
				}
			}
		}
	})

	for i := 0; i < len(quads); i++ {
		quad := quads[i]
		fmt.Printf("%d %d %d %d\n", quad.a, quad.b, quad.c, quad.d)
	}
}

// set of quads
func makeQuadSet() *QuadSet {
	return &QuadSet{
		container: make(map[Quad]bool),
	}
}

type QuadSet struct {
	container map[Quad]bool
}

func (c *QuadSet) GetArray() []Quad {
	keys := make([]Quad, c.Size())
	i := 0
	for k := range c.container {
		keys[i] = k
		i++
	}
	return keys
}

func (c *QuadSet) Exists(key Quad) bool {
	_, exists := c.container[key]
	return exists
}

func (c *QuadSet) Add(key Quad) {
	c.container[key] = true
}

func (c *QuadSet) Remove(key Quad) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("remove error: item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *QuadSet) Size() int {
	return len(c.container)
}

// set of integers
func makeIntSet() *IntSet {
	return &IntSet{
		container: make(map[int]bool),
	}
}

type IntSet struct {
	container map[int]bool
}

func (c *IntSet) Exists(key int) bool {
	_, exists := c.container[key]
	return exists
}

func (c *IntSet) Add(key int) {
	c.container[key] = true
}

func (c *IntSet) Remove(key int) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("remove error: item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *IntSet) Size() int {
	return len(c.container)
}

// быстрая сортировка четырёх элементов
func sortFour(a, b, c, d int) (int, int, int, int) {
	var low1, high1, low2, high2, lowest, middle1, highest, middle2 int
	if a < b {
		low1 = a
		high1 = b
	} else {
		low1 = b
		high1 = a
	}

	if c < d {
		low2 = c
		high2 = d
	} else {
		low2 = d
		high2 = c
	}

	if low1 < low2 {
		lowest = low1
		middle1 = low2
	} else {
		lowest = low2
		middle1 = low1
	}

	if high1 > high2 {
		highest = high1
		middle2 = high2
	} else {
		highest = high2
		middle2 = high1
	}

	if middle1 < middle2 {
		return lowest, middle1, middle2, highest
	} else {
		return lowest, middle2, middle1, highest
	}
}

func checkAllNumbersAreDifferent(i1, j1, i2, j2 int) bool {
	if i1 == j1 || i1 == i2 || i1 == j2 {
		return false
	}
	if j1 == i1 || j1 == i2 || j1 == j2 {
		return false
	}
	if i2 == i1 || i2 == j1 || i2 == j2 {
		return false
	}
	if j2 == i1 || j2 == j1 || j2 == i2 {
		return false
	}
	return true
}
