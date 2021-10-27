package day1

import (
	"math"
	"strconv"
)

type Day1 struct {
}

func (d Day1) SolveProblemPart1(input []string) (string, error) {
	total := 0
	for _, s := range input {
		m, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		total += calculateFuel(m)
	}
	return strconv.Itoa(total), nil
}
func (d Day1) SolveProblemPart2(input []string) (string, error) {
	total := 0
	for _, s := range input {
		m, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		total += calculateFuelWithFuelMassIncluded(m)
	}
	return strconv.Itoa(total), nil
}
func (d Day1) Day() string {
	return "1"
}

func calculateFuel(mass int) int {
	m := float64(mass) / 3
	m = math.Floor(m)
	return int(m - 2)
}

func calculateFuelWithFuelMassIncluded(mass int) int {
	mass = calculateFuel(mass)
	if mass <= 0 {
		return 0
	} else {
		mass += calculateFuelWithFuelMassIncluded(mass)
	}
	return mass
}
