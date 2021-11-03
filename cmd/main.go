package main

import (
	"log"
	"os"

	"github.com/AdventOfCode2019/day1"
	"github.com/AdventOfCode2019/day2"
	"github.com/AdventOfCode2019/day3"
	"github.com/AdventOfCode2019/util"
)

var solvers = []util.Solver{
	day1.Day1{},
	day2.Day2{},
	day3.Day3{},
}

func main() {
	solveOne := false
	toBeSolved := ""
	var err error
	if len(os.Args) > 1 {
		solveOne = true
		toBeSolved = os.Args[1]
		if err != nil {
			log.Fatal(err)
		}
	}
	if solveOne {
		solver := getSolver(toBeSolved)
		if solver == nil {
			log.Fatal("Solver not present")
		}
		util.HandleSolver(solver)

	} else {
		for _, s := range solvers {
			util.HandleSolver(s)
		}
	}
}

func getSolver(d string) util.Solver {
	for _, s := range solvers {
		if s.Day() == d {
			return s
		}
	}
	return nil
}
