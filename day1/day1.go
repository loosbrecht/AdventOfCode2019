package day1

import (
	"fmt"
	"log"

	"github.com/AdventOfCode2019/util"
)

type Day1 struct {
}

func (d Day1) SolveProblemPart1() (string, error) {
	lines, err := util.ReadInput(d.Day())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	return lines[0], nil
}
func (d Day1) SolveProblemPart2() (string, error) {
	lines, err := util.ReadInput(d.Day())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	return lines[0], nil
}
func (d Day1) Day() string {
	return "1"
}
