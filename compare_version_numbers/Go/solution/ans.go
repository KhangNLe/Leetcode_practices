package solution

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	v1Size := len(v1)
	v2Size := len(v2)

	if v1Size < v2Size {
		diff := v2Size - v1Size
		for diff > 0 {
			v1 = append(v1, "0")
			diff--
		}
	} else if v2Size < v1Size {
		diff := v1Size - v2Size
		for diff > 0 {
			v2 = append(v2, "0")
			diff--
		}
	}

	for i := 0; i < len(v1); i++ {
		if ok, res := compareArr(v1[i], v2[i]); ok {
			return res
		}
	}
	return 0
}

func compareArr(v1, v2 string) (bool, int) {
	num1, _ := strconv.Atoi(v1)
	num2, _ := strconv.Atoi(v2)

	if num1 == num2 {
		return false, 0
	} else if num1 > num2 {
		return true, 1
	} else {
		return true, -1
	}
}
