package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFuel(t *testing.T) {
	var tests = []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "12",
			in:   12,
			res:  2,
		},
		{
			name: "14",
			in:   14,
			res:  2,
		},
		{
			name: "1969",
			in:   1969,
			res:  654,
		},
		{
			name: "100756",
			in:   100756,
			res:  33583,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateFuel(tt.in)
			assert.Equal(t, result, tt.res)
		})
	}
}

func TestCalculateFuelWithFuelMassIncluded(t *testing.T) {
	var tests = []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "12",
			in:   12,
			res:  2,
		},
		{
			name: "1969",
			in:   1969,
			res:  966,
		},
		{
			name: "100756",
			in:   100756,
			res:  50346,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateFuelWithFuelMassIncluded(tt.in)
			assert.Equal(t, result, tt.res)
		})
	}
}
