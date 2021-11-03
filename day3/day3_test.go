package day3

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddWire(t *testing.T) {
	grid := make(map[int]map[int]int)
	split := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	grid = addWire(split, grid)
	split = strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	grid = addWire(split, grid)
	points := findCrossPoints(grid)
	fmt.Println(points)
	fmt.Println(getSmallestDistance(points))
	fmt.Println("stop")
}
