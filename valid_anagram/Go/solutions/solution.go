package solutions

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	//Create a counting map of 127 standard ASCII range 0-126
	var characters [127]int

	for i := range len(s) {
		characters[byte(s[i])]++
		characters[byte(t[i])]--
	}

	for _, num := range characters {
		if num != 0 {
			return false
		}
	}
	return true
}
