package solution

func AvoidFlood(rains []int) []int {
	ans := make([]int, len(rains))
	dryDays := []int{}
	waterMonitor := make(map[int]int)

	for i := range len(rains) {
		if rains[i] == 0 {
			dryDays = append(dryDays, i)
			ans[i] = 1
		} else {
			if lastRain, ok := waterMonitor[rains[i]]; ok {
				if len(dryDays) > 0 {
					waterMonitor[rains[i]] = i
					dryIdx := findDryDay(dryDays, lastRain)
					if dryDays[dryIdx] < lastRain {
						return []int{}
					}
					ans[dryDays[dryIdx]] = rains[i]
					dryDays = append(dryDays[:dryIdx], dryDays[dryIdx+1:]...)
					ans[i] = -1
				} else {
					return []int{}
				}
			} else if _, ok := waterMonitor[rains[i]]; !ok {
				waterMonitor[rains[i]] = i
				ans[i] = -1
			} else {
				return []int{}
			}
		}
	}
	return ans
}

func findDryDay(dryDays []int, lastRainDay int) int {
	l := 0
	r := len(dryDays) - 1
	ans := 0

	for l <= r {
		m := (l + r) / 2
		if dryDays[m] > lastRainDay {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
