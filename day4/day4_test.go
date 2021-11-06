package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValideRules(t *testing.T) {
	var tests = []struct {
		name     string
		password int
		res      bool
	}{
		{
			name:     "basic",
			password: 111111,
			res:      true,
		},
		{
			name:     "basic2",
			password: 223450,
			res:      false,
		},
		{
			name:     "basic3",
			password: 123789,
			res:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validValue(tt.password)

			assert.Equal(t, result, tt.res)
		})
	}
}

func TesthaveGroupOfTwoAndNoLargerGroups(t *testing.T) {
	var tests = []struct {
		name     string
		password string
		res      bool
	}{
		{
			name:     "basic",
			password: "112233",
			res:      true,
		},
		{
			name:     "basic2",
			password: "123444",
			res:      false,
		},
		{
			name:     "basic3",
			password: "111122",
			res:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := haveGroupOfTwoAndNoLargerGroups(tt.password)
			assert.Equal(t, result, tt.res)
		})
	}
}
