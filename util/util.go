package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ReadInput(dayNum, partNum string) ([]string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	name := filepath.Join(path, "day"+dayNum, "input", "inputpart"+partNum)
	f, _ := os.Open(name)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)

	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result, nil
}

type Solver interface {
	SolveProblemPart1(nput []string) (string, error)
	SolveProblemPart2(nput []string) (string, error)

	Day() string
}

func HandleSolver(s Solver) {
	LogTitle(s.Day())
	lines, err := ReadInput(s.Day(), "1")
	if err != nil {
		log.Fatal(err)
	}
	result, err := s.SolveProblemPart1(lines)
	if err != nil {
		LogError(err.Error())
	} else {
		LogSolution("1", result)
	}
	lines, err = ReadInput(s.Day(), "2")
	if err != nil {
		log.Fatal(err)
	}
	result, err = s.SolveProblemPart2(lines)
	if err != nil {
		LogError(err.Error())
	} else {
		LogSolution("2", result)
	}

}

func LogSolution(part, msg string) {
	color.Green("-----------------------")
	color.Green(fmt.Sprintf("Solution part %s: %s", part, msg))
	color.Green("-----------------------")

}

func LogError(msg string) {
	color.Red("-----------------------")
	color.Red("Error: " + msg)
	color.Red("-----------------------")
}

func LogDebug(msg string) {
	color.Magenta(msg)
}

func LogTitle(dayNum string) {
	color.Blue("-----------------------")
	color.Blue(" Start solving day " + dayNum)
	color.Blue("-----------------------")
}
