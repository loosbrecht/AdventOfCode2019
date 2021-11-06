package day4

import (
	"strconv"
	"strings"
)

type Day4 struct{}

func (Day4) SolveProblemPart1(input []string) (string, error) {
	inp := input[0]
	split := strings.Split(inp, "-")
	min, err := strconv.Atoi(split[0])
	if err != nil {
		return "", err
	}
	max, err := strconv.Atoi(split[1])
	if err != nil {
		return "", err
	}
	counted := generateValidPasswords(min, max, validationFuncs)
	return strconv.Itoa(counted), nil

}
func (Day4) SolveProblemPart2(input []string) (string, error) {
	inp := input[0]
	split := strings.Split(inp, "-")
	min, err := strconv.Atoi(split[0])
	if err != nil {
		return "", err
	}
	max, err := strconv.Atoi(split[1])
	if err != nil {
		return "", err
	}
	extraValidationFuncs := append(validationFuncs, haveGroupOfTwoAndNoLargerGroups)
	counted := generateValidPasswords(min, max, extraValidationFuncs)
	return strconv.Itoa(counted), nil
}
func (Day4) Day() string {
	return "4"
}

func generateValidPasswords(min, max int, validateFuncs []func(string) bool) int {
	count := 0
	for i := min; i < max; i++ {
		if validValue(i, validateFuncs) {
			count++
		}
	}
	return count
}

func validValue(val int, validateFuncs []func(string) bool) bool {
	strFormat := strconv.Itoa(val)
	for _, vFunc := range validateFuncs {
		if result := vFunc(strFormat); !result {
			return result
		}
	}
	return true
}

var validationFuncs = []func(string) bool{
	lengthSix,
	hasTwoAdjacent,
	valuesIncreaseOrStayTheSame,
}

func lengthSix(strFormat string) bool {
	return len(strFormat) == 6
}

func hasTwoAdjacent(strFormat string) bool {
	current := strFormat[0]
	for i := 1; i < len(strFormat); i++ {
		if current == strFormat[i] {
			return true
		}
		current = strFormat[i]
	}
	return false
}

func valuesIncreaseOrStayTheSame(strFormat string) bool {
	initial := '0'
	for _, c := range strFormat {
		if c < initial {
			return false
		}
		initial = c
	}
	return true
}

func haveGroupOfTwoAndNoLargerGroups(strFormat string) bool {
	counter := map[rune]int{}
	for _, c := range strFormat {
		counter[c]++
	}

	adjTwo := false
	biggerGroup := false
	for _, v := range counter {
		if v == 2 {
			adjTwo = true
		}
		if v > 2 {
			biggerGroup = true
		}
	}
	if !adjTwo && !biggerGroup {
		return false
	} else if adjTwo && biggerGroup {
		return true
	} else if adjTwo {
		return true
	} else {
		return false
	}

}
