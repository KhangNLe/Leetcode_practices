package solution

import "strings"

func IsPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)

	for l < r {
		if !('a' <= s[l] && s[l] <= 'z' || '0' <= s[l] && s[l] <= '9') {
			l++
		} else if !('a' <= s[r] && s[r] <= 'z' || '0' <= s[r] && s[r] <= '9') {
			r--
		} else {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}
	return s[l] == s[r]
}
