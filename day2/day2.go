package day2

import (
	"strconv"
	"strings"
)

type Day2 struct {
}

func (Day2) SolveProblemPart1(input []string) (string, error) {
	program, err := transformInput(input[0])
	if err != nil {
		return "", err
	}
	program[1] = 12
	program[2] = 2
	solved := solveProgram(program)
	return strconv.Itoa(solved[0]), nil
}
func (Day2) SolveProblemPart2(input []string) (string, error) {
	program, err := transformInput(input[0])
	if err != nil {
		return "", err
	}
	solved := findCorrectNounAndVerb(program)

	return strconv.Itoa(solved), nil
}
func (Day2) Day() string {
	return "2"
}

func transformInput(str string) ([]int, error) {
	split := strings.Split(str, ",")
	prog := make([]int, len(split))
	var err error
	for i, s := range split {
		prog[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return prog, nil
}

func solveProgram(program []int) []int {
	ind := 0
	var op func(i, j int) int
	pos1 := func() int { return program[program[ind+1]] }
	pos2 := func() int { return program[program[ind+2]] }
FORLOOP:
	for {
		switch program[ind] {
		case 1:
			op = func(i, j int) int { return i + j }
		case 2:
			op = func(i, j int) int { return i * j }
		case 99:
			break FORLOOP
		}
		program[program[ind+3]] = op(pos1(), pos2())
		ind += 4
	}
	return program
}

func solveProgramGiveResult(program []int) int {
	resultProgram := solveProgram(program)
	return resultProgram[0]
}

func applyInput(noun, verb int, program []int) []int {
	program[1] = noun
	program[2] = verb
	return program
}

type result struct {
	noun   int
	verb   int
	result int
}

func findCorrectNounAndVerb(program []int) int {
	goal := 19690720
	resultChan := make(chan result)

	go func() {
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				programN := make([]int, len(program))
				copy(programN, program)
				res := solveProgramGiveResult(applyInput(noun, verb, programN))
				resultChan <- result{
					noun:   noun,
					verb:   verb,
					result: res,
				}
			}
		}
		close(resultChan)
	}()

	for res := range resultChan {
		if res.result == goal {
			return 100*res.noun + res.verb
		}
	}
	return 0
}
