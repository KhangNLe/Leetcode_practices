package solution

import (
	"strconv"
	"strings"
)

func FractionToDecimal(numerator int, denominator int) string {
	var res strings.Builder

	if numerator == 0 {
		return "0"
	}

	if (numerator < 0 || denominator < 0) &&
		!(numerator < 0 && denominator < 0) {
		res.WriteByte('-')
	}

	n := abs(numerator)
	d := abs(denominator)

	res.WriteString(strconv.Itoa(n / d))
	remainder := n % d

	if remainder == 0 {
		return res.String()
	}

	res.WriteByte('.')
	remainderMap := make(map[int]int)

	for remainder != 0 {
		if idx, ok := remainderMap[remainder]; ok {
			str := res.String()
			return str[:idx] + "(" + str[idx:] + ")"
		}

		remainderMap[remainder] = res.Len()
		remainder *= 10
		res.WriteString(strconv.Itoa(remainder / d))
		remainder = remainder % d
	}
	return res.String()
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
