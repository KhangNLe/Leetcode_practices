package solution

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		key := createCharFrequenceKey(word)
		if _, ok := groups[key]; !ok {
			groups[key] = []string{word}
		} else {
			groups[key] = append(groups[key], word)
		}
	}

	var ans [][]string

	for _, val := range groups {
		ans = append(ans, val)
	}
	return ans
}

func createCharFrequenceKey(word string) string {
	freq := make([]byte, 26)
	for _, char := range []byte(word) {
		freq[char%97] += char
	}
	return string(freq)
}
