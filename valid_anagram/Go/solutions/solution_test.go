package solutions

import (
	"solution/solutions"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		t       string
		s       string
		anagram bool
	}{
		{
			"shrub",
			"brush",
			true,
		},
		{
			"skill",
			"kills",
			true,
		},
		{
			"apple",
			"pleaps",
			false,
		},
		{
			"aaa",
			"aa",
			false,
		},
		{
			"looted",
			"toledo",
			true,
		},
		{
			"leetcode",
			"clodet",
			false,
		},
	}

	for _, input := range inputs {
		anagram := solutions.IsAnagram(input.s, input.t)
		if input.anagram != anagram {
			t.Errorf(
				`Error when comparing %s and %s
				Expect: %t, got %t`,
				input.t, input.s, input.anagram, anagram,
			)
		}
	}
}
