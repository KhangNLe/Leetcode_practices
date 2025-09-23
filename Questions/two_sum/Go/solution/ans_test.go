package solution_test

import (
	"sort"
	"testing"
	"twoSum/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			[]int{3, 4, 5, 6},
			7,
			[]int{0, 1},
		},
		{
			[]int{4, 5, 6},
			10,
			[]int{0, 2},
		},
		{
			[]int{5, 5},
			10,
			[]int{0, 1},
		},
	}

	for _, input := range inputs {
		ans := solution.TwoSum(input.nums, input.target)
		if !compareArrs(input.output, ans) {
			t.Fatalf(`
				Mismatching input. expected %v, got %v\n
				`, input.output, ans)
		}
	}
}

func compareArrs(expected []int, input []int) bool {
	if len(expected) != len(input) {
		return false
	}
	sort.Ints(expected)
	sort.Ints(input)

	for i := range len(expected) {
		if expected[i] != input[i] {
			return false
		}
	}
	return true
}
