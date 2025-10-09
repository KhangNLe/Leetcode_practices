package solution

func MinTime(skill []int, mana []int) int64 {
	n := len(skill)
	freeTime := make([]int64, n)

	for _, m := range mana {
		time := freeTime[0]

		for i := 1; i < n; i++ {
			time = max(time+int64(skill[i-1]*m), freeTime[i])
		}

		freeTime[n-1] = time + int64(skill[n-1]*m)

		for i := n - 2; i >= 0; i-- {
			freeTime[i] = freeTime[i+1] - int64(skill[i+1]*m)

		}
	}

	return freeTime[n-1]
}
