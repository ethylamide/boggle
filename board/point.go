package board

import "fmt"

type Point struct {
	x int
	y int
	n int
}

func (point Point) inside() bool {
	x := point.x
	y := point.y
	n := point.n
	return (x >= 0 && x < n) && (y >= 0 && y < n)
}

func (point Point) Vertex() int {
	return point.x*point.n + point.y
}

func (point Point) Neighbors() []Point {
	x := point.x
	y := point.y
	n := point.n

	points := []Point{
		Point{x: x - 1, y: y, n: n},
		Point{x: x - 1, y: y - 1, n: n},
		Point{x: x - 1, y: y + 1, n: n},
		Point{x: x, y: y - 1, n: n},
		Point{x: x, y: y + 1, n: n},
		Point{x: x + 1, y: y, n: n},
		Point{x: x + 1, y: y - 1, n: n},
		Point{x: x + 1, y: y + 1, n: n},
	}

	neighbors := make([]Point, 0)
	for _, point := range points {
		if point.inside() {
			neighbors = append(neighbors, point)
		}
	}

	return neighbors
}

func (point *Point) String() string {
	return fmt.Sprintf("Point{x:%d, y:%d, n:%d}", point.x, point.y, point.n)
}
