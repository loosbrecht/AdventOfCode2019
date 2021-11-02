package day2

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveProgram(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "basic",
			in:   "1,0,0,0,99",
			res:  "2,0,0,0,99",
		},
		{
			name: "basic2",
			in:   "2,3,0,3,99",
			res:  "2,3,0,6,99",
		},
		{
			name: "basic3",
			in:   "2,4,4,5,99,0",
			res:  "2,4,4,5,99,9801",
		},
		{
			name: "basic4",
			in:   "1,1,1,4,99,5,6,0,99",
			res:  "30,1,1,4,2,5,6,0,99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog, err := transformInput(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			result := solveProgram(prog)

			assert.Equal(t, toOneLine(result), tt.res)
		})
	}
}

func toOneLine(res []int) string {
	s := make([]string, len(res))
	for i, r := range res {
		s[i] = strconv.Itoa(r)
	}
	return strings.Join(s, ",")
}
