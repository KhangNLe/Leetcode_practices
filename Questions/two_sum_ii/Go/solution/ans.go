package solution

func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		remain := target - numbers[right]
		if remain > numbers[left] {
			left++
		} else if remain < numbers[left] {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{left + 1, right + 1}
}
