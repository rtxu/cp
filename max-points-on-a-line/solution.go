package main

import "fmt"

type Point struct {
	X int
	Y int
}

func onTheSameLine(p1, p2, p3 Point) bool {
	v1 := Point{
		X: p2.X - p1.X,
		Y: p2.Y - p1.Y,
	}
	v2 := Point{
		X: p3.X - p1.X,
		Y: p3.Y - p1.Y,
	}

	return (int64(v1.X)*int64(v2.Y) - int64(v2.X)*int64(v1.Y)) == 0
}

func maxPoints(points []Point) int {
	pointMap := make(map[Point]int)
	n := len(points)
	for i := 0; i < n; i++ {
		pointMap[points[i]]++
	}
	diffPointCount := len(pointMap)
	switch diffPointCount {
	case 0:
		return 0
	case 1:
		return pointMap[points[0]]
	default:
		max := 2
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if points[i] == points[j] {
					continue
				}
				count := 0
				for k := 0; k < n; k++ {
					if onTheSameLine(points[i], points[j], points[k]) {
						count++
					}
				}
				if count > max {
					max = count
				}
			}
		}
		return max
	}
}

func main() {
	var input []Point
	var expected int

	input = []Point{}
	expected = 0
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		Point{0, 0},
	}
	expected = 1
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		Point{0, 0},
		Point{5, 5},
	}
	expected = 2
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		Point{0, 0},
		Point{5, 5},
		Point{7, 9},
	}
	expected = 2
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		Point{0, 0},
		Point{5, 5},
		Point{7, 9},
		Point{10, 10},
	}
	expected = 3
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		{84, 250}, {0, 0}, {1, 0}, {0, -70}, {0, -70}, {1, -1}, {21, 10}, {42, 90}, {-42, -230},
	}
	expected = 6
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))

	input = []Point{
		{1, 1}, {1, 1}, {1, 1},
	}
	expected = 3
	fmt.Printf("input: %v, expected: %d, actual: %d\n", input, expected, maxPoints(input))
}
