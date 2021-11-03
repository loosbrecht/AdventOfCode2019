package day3

import (
	"math"
	"strconv"
	"strings"
)

type Day3 struct {
}

func (Day3) SolveProblemPart1(input []string) (string, error) {
	grid := make(map[int]map[int]int)
	grid = addWire(strings.Split(input[0], ","), grid)
	grid = addWire(strings.Split(input[1], ","), grid)
	points := findCrossPoints(grid)
	dist := getSmallestDistance(points)
	distS := strconv.Itoa(dist)
	return distS, nil
}
func (Day3) SolveProblemPart2(input []string) (string, error) { return "", nil }
func (Day3) Day() string                                      { return "3" }

type point struct {
	x int
	y int
}

func (p point) Distance() float64 {
	return math.Abs(float64(p.x)) + math.Abs(float64(p.y))
}

func getSmallestDistance(points []point) int {
	shortest := math.MaxFloat64
	for _, p := range points {
		dist := p.Distance()
		if dist < shortest {
			shortest = dist
		}
	}
	return int(shortest)
}

func findCrossPoints(grid map[int]map[int]int) []point {
	results := make([]point, 0)
	for x, m := range grid {
		for y, c := range m {
			if x == 0 && y == 0 {
				continue
			}
			if c >= 2 {
				results = append(results, point{x: x, y: y})
			}
		}
	}
	return results
}

func addWire(wirePath []string, grid map[int]map[int]int) map[int]map[int]int {
	currX := 0
	currY := 0
	for _, path := range wirePath {
		dir := path[0]
		length, _ := strconv.Atoi(path[1:])

		switch dir {
		case 'R':
			for i := 0; i < length; i++ {
				if grid[currX] == nil {
					grid[currX] = make(map[int]int)
				}
				grid[currX][currY] += 1
				currX++
			}
		case 'L':
			for i := 0; i < length; i++ {
				if grid[currX] == nil {
					grid[currX] = make(map[int]int)
				}
				grid[currX][currY] += 1
				currX--
			}
		case 'U':
			for i := 0; i < length; i++ {
				if grid[currX] == nil {
					grid[currX] = make(map[int]int)
				}
				grid[currX][currY] += 1
				currY++
			}
		case 'D':
			for i := 0; i < length; i++ {
				if grid[currX] == nil {
					grid[currX] = make(map[int]int)
				}
				grid[currX][currY] += 1
				currY--
			}
		}
	}
	return grid
}
