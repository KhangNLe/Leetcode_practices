package solution

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	substring := make([]bool, 127)
	longest := 1
	start := 0
	substring[s[0]] = true

	for i := 1; i < len(s); i++ {
		if substring[s[i]] {
			longest = max(longest, (i - start))
			for s[start] != s[i] {
				substring[s[start]] = false
				start++
			}
			start++
		}
		substring[s[i]] = true
	}
	return max(longest, (len(s) - start))
}
