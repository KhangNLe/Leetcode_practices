package solution_test

import (
	"anagrams/solution"
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	inputs := []struct {
		input  []string
		output [][]string
	}{
		{
			[]string{"act", "pots", "tops", "cat", "stop", "hat"},
			[][]string{
				{"hat"},
				{"act", "cat"},
				{"stop", "pots", "tops"},
			},
		},
		{
			[]string{"x"},
			[][]string{
				{"x"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
	}

	for _, input := range inputs {
		ans := solution.GroupAnagrams(input.input)
		if !compareAns(ans, input.output) {
			t.Fatalf(`
				Mismatching answer.
				Expected %s, got %s\n
				`, input.input, ans)
		}
	}
}

func compareAns(a [][]string, b [][]string) bool {
	input := make(map[string]struct{})
	expect := make(map[string]struct{})

	mergeWords(a, &input)
	mergeWords(b, &expect)
	return reflect.DeepEqual(input, expect)
}

func mergeWords(a [][]string, group *map[string]struct{}) {
	for _, words := range a {
		for _, word := range words {
			(*group)[word] = struct{}{}
		}
	}
}
