package solution

/*
array length up to 1000
starting from idx = 0 , get the sum of the current idx and its neighboor
mod 10 the sum as the new number, or we an get the abs of num - 10
repeat until array len is 1
*/
func TriangularSum(nums []int) int {
	size := len(nums)
	for size != 1 {
		for i := 1; i < size; i++ {
			nums[i-1] = (nums[i-1] + nums[i]) % 10
		}
		size -= 1
	}
	return nums[0]
}
