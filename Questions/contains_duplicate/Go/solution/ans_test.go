package solution

import (
	"duplicate/solution"
	"math/rand"
	"testing"
)

func TestDuplicate(t *testing.T) {
	for i := range 100 {
		arr, duplicate := generateArr(i)
		if isDuplicate := solution.HasDuplicate(arr); isDuplicate != duplicate {
			t.Fatalf("Mismatching solution. Expected: %t, got %t\n",
				duplicate, isDuplicate)
		}
	}
}

func generateArr(iter int) ([]int, bool) {
	amount := rand.Intn(100)
	arr := make([]int, amount)
	if iter^1 < iter {
		generateArrWithDuplicate(&arr)
		return arr, true
	} else {
		generateArrWithoutDuplicate(&arr)
		return arr, false
	}
}

func generateArrWithDuplicate(arr *[]int) {
	size := len(*arr)
	for i := range size {
		(*arr)[i] = rand.Intn(size - 1)
	}
	(*arr)[size-1] = (*arr)[0]
}

func generateArrWithoutDuplicate(arr *[]int) {
	for i := range len(*arr) {
		(*arr)[i] = i
	}
}
