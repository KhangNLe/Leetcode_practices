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

	maxSize := max(v1Size, v2Size)

	for i := range maxSize {
		var ver1, ver2 int

		if i < v1Size {
			ver1, _ = strconv.Atoi(v1[i])
		}

		if i < v2Size {
			ver2, _ = strconv.Atoi(v2[i])
		}

		if ver1 > ver2 {
			return 1
		} else if ver2 > ver1 {
			return -1
		}
	}

	return 0
}
