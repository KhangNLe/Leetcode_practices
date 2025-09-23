package solution

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	product := 1
	for i := range len(nums) {
		ans[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(ans) - 1; i > -1; i-- {
		ans[i] *= product
		product *= nums[i]
	}

	return ans
}
