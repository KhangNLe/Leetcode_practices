# 611. Valid Triangle Number
**Medium**
## Question
Given an integer array `nums`, return the number of triplets chosen from the array that can make triangles if we take them as side lengths of a triangle.

## Example
### Example 1
```yaml
Input: nums = [2,2,3,4]
Output: 3

Explanation: Valid combinations are:
2,3,4 (using the first 2)
2,3,4 (using the second 2)
2,2,3
```

### Example 2
```yaml
Input: nums = [4,2,3,4]
Output: 4
```

## Constraints
- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`

## Solution
I don't think Leetcode do a good job with providing all the necessary information for the problem and somewhat just assume that everyone would know what make a triangle. When first trying to solve this, I was mistaken the problem to as for triangle triplet of `a**2 + b**2 = c**2` for the given `integer` array. However, that was not the case, the question just asked if any 3 numbers from different indexes in the given array would form a triangle. A triangle is valid if the sum of any two side is bigger than the third side.

With this information, I decided to use the pointers method find out all the possible number of triangles. In order to perform the pointer method, we will first need to sort our arrays. After than, I will start a for loop going back to 0 from the last index of the array. Inside each loop, there will a pointer at index 0 called `l`, and a pointer at index 1 step before the current index in the loop called `r`. Then we will do a while loop inside the loop of `l < r`. We will then take the sum of the element at index `l` and `r` and compare it with the element at the current index in the for loop. If the sum if bigger, we can conclude that every summation from `[l, l+1, ..., r-1] + r` is bigger than the current number in the index. We can add that to the total count by find out the steps with `r-l` and decrement `r` by 1. If the sum is equal or smaller than the current number, then we will increment `l` by 1. After the for loop finish, we will have the solution of the total possible triangle that would be form with our algorithm. For the time complexity, using the sort function will cost us `O(nlogn)` and the for loop function will cost us `O(n**2)`. So our total time complexity is `O(n**2)`. And for our space complexity, it will be `O(1)`.
