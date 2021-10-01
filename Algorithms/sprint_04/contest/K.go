package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	var line string

	// читаем количество выходов из метро
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	// читаем координаты выходов из метро
	xMetroCoords := make([]int, n)
	yMetroCoords := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		values := strings.Split(line, " ")
		xMetroCoords[i], _ = strconv.Atoi(values[0])
		yMetroCoords[i], _ = strconv.Atoi(values[1])
	}

	// читаем количество автобусных остановок
	var m int

	scanner.Scan()
	line = scanner.Text()
	m, _ = strconv.Atoi(line)

	// читаем координаты остановок и запоминаем для каждой точки (x,y) количество остановок в ней
	stopsCount := make(map[int]int)

	// также сохраняем множество координат точек остановок
	stopsPointsSet := makePointSet()

	var xStop, yStop, h int
	for i := 0; i < m; i++ {
		scanner.Scan()
		line = scanner.Text()
		values := strings.Split(line, " ")
		xStop, _ = strconv.Atoi(values[0])
		yStop, _ = strconv.Atoi(values[1])
		stopsPointsSet.Add(newPoint(xStop, yStop))
		h = hash(xStop, yStop)
		count, ok := stopsCount[h]
		if !ok {
			stopsCount[h] = 1
		} else {
			stopsCount[h] = count + 1
		}
	}

	var xMetro, yMetro int
	maxOutNumber := -1
	maxStopsCount := 0
	nearStopsCount := 0
	// для каждого выхода из метро определяем, сколько остановок находится на расстоянии <= 20 метров от него
	for i := 0; i < n; i++ {
		nearStopsCount = 0
		xMetro = xMetroCoords[i]
		yMetro = yMetroCoords[i]
		for xPoint := xMetro - 20; xPoint <= xMetro+20; xPoint++ {
			for yPoint := yMetro - 20; yPoint <= yMetro+20; yPoint++ {
				if xPoint == xMetro && yPoint == yMetro {
					continue
				}
				// если точка на расстоянии не более 20 от выхода из метро
				if distanceLessThan(xMetro, yMetro, xPoint, yPoint, 20) {
					// проверяем, сколько в этой точке остановок
					h = hash(xPoint, yPoint)
					count, ok := stopsCount[h]
					if ok {
						// если в этой точке действительно есть остановка (могут быть коллизии хеш-функции)
						if stopsPointsSet.Exists(newPoint(xPoint, yPoint)) {
							nearStopsCount += count
						}
					}
				}
			}
		}
		if nearStopsCount > maxStopsCount {
			maxStopsCount = nearStopsCount
			maxOutNumber = i + 1 // отсчёт - от единицы
		}
	}
	fmt.Println(maxOutNumber)
}

// Cantor pairing function
func hash(x, y int) int {
	return ((x + y) * (x + y + 1) / 2) + y
}

func distanceLessThan(x1, y1, x2, y2, delta int) bool {
	return (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= delta*delta
}

func distanceSquare(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

// set of pairs
func makePointSet() *PointSet {
	return &PointSet{
		container: make(map[Point]bool),
	}
}

type PointSet struct {
	container map[Point]bool
}

func (c *PointSet) Exists(key Point) bool {
	_, exists := c.container[key]
	return exists
}

func (c *PointSet) Add(key Point) {
	c.container[key] = true
}

type Point struct {
	x, y int
}

func newPoint(x, y int) Point {
	return Point{x, y}
}
