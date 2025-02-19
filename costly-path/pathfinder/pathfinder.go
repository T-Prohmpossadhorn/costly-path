package pathfinder

import (
	"errors"
)

func FindCostlyPath(val [][]int) (int, error) {
	if len(val) == 0 {
		return 0, errors.New("input error")
	}
	for i, v := range val {
		if len(v) != i+1 {
			return 0, errors.New("input error")
		}
	}
	return recur(val[0][0], val), nil
}

func recur(total int, val [][]int) int {
	if len(val) >= 2 {
		value1 := make([][]int, len(val)-1)
		copy(value1, val[1:])
		for i, _ := range value1 {
			value1[i] = value1[i][:len(value1[i])-1]
		}

		total1 := recur(total+value1[0][0], value1)

		value2 := make([][]int, len(val)-1)
		copy(value2, val[1:])
		for i, _ := range value2 {
			value2[i] = value2[i][1:]
		}

		total2 := recur(total+value2[0][0], value2)

		if total1 > total2 {
			total = total1
		} else {
			total = total2
		}
	}

	return total
}
