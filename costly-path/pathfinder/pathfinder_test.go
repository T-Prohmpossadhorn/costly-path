package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCostlyPath(t *testing.T) {
	testcase := []struct {
		name  string
		input [][]int

		expect struct {
			result  int
			isError bool
		}
	}{
		{
			name:  "Happy case",
			input: [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}},

			expect: struct {
				result  int
				isError bool
			}{
				result:  237,
				isError: false,
			},
		}, {
			name:  "Not symmetry",
			input: [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6}},

			expect: struct {
				result  int
				isError bool
			}{
				result:  0,
				isError: true,
			},
		}, {
			name:  "Blank",
			input: [][]int{},

			expect: struct {
				result  int
				isError bool
			}{
				result:  0,
				isError: true,
			},
		},
	}

	for _, v := range testcase {
		t.Run(v.name, func(t *testing.T) {
			ret, err := FindCostlyPath(v.input)
			if !v.expect.isError {
				assert.Nil(t, err)
				assert.Equal(t, ret, v.expect.result)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
