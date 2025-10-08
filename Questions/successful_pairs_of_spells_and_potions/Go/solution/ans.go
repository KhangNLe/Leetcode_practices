package solution

import "sort"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))

	for i := range len(spells) {
		l := 0
		r := len(potions)

		for l < r {
			m := (l + r) / 2
			strength := int64(spells[i] * potions[m])
			if strength == success {
				if m != 0 && potions[m] != potions[m-1] {
					l = m
					break
				}
				r = m - 1
			} else if strength > success {
				r = m
			} else {
				l = m + 1
			}
		}
		ans[i] = len(potions) - l
	}
	return ans
}
