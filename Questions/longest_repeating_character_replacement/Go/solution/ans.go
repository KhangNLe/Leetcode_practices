package solution

func CharacterReplacements(s string, k int) int {
	charFreq := make([]int, 26)
	ans := 0
	start := 0
	mostFreq := 0

	for i := range len(s) {
		charFreq[s[i]-65]++

		if charFreq[s[i]-65] > mostFreq {
			mostFreq = charFreq[s[i]-65]
		}

		if (i-start+1)-mostFreq > k {
			charFreq[s[start]-65]--
			start++
		}

		ans = max(ans, (i - start + 1))
	}

	return ans
}
