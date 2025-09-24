package solution_test

import (
	"math/rand"
	"sequence/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		nums []int
		res  int
	}{
		{
			[]int{2, 20, 4, 10, 3, 4, 5},
			4,
		},
		{
			[]int{0, 3, 2, 5, 4, 6, 1, 1},
			7,
		},
	}

	for _, input := range inputs {
		ans := solution.LongestConsecutive(input.nums)
		if ans != input.res {
			t.Fatalf(`
				Mismatching result. Expect %d, got %d
				`, input.res, ans)
		}
	}

	arr := generateBigArr()

	ans := solution.LongestConsecutive(*arr)
	if ans == 0 {
		t.Fatalf(`
			Wrong result for big array of 10,000
			`)
	}

	arr = generateBigArrInOrder()
	ans = solution.LongestConsecutive(*arr)
	if ans == 0 {
		t.Fatal("Wrong result for big array in order")
	}
}

func generateBigArr() *[]int {
	arr := make([]int, 1000000)

	for i := range len(arr) {
		arr[i] = rand.Intn(100000)
	}
	return &arr
}

func generateBigArrInOrder() *[]int {
	arr := make([]int, 1000000)

	for i := range len(arr) {
		arr[i] = i
	}
	return &arr
}
