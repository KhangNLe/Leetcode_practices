# 167. Two Sum II - Input Array is Sorted
**Medium**

## Question
Given a 1-indexed array of integers `numbers` that is already **sorted in non-decreasing order**, find two numbers such that they add up to a specific `target` number. Let these two numbers be `numbers[index1]` and `numbers[index2]` where `1 <= index1 < index2 <= numbers.length`.

Return the indices of the two numbers, `index1` and `index2`, **added by one** as an integer array `[index1, index2]` of length 2.

The tests are generated such that there is **exactly one solution**. You may not use the same element twice.

Your solution must use only constant extra space.

## Example
### Example 1
```yaml
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
```
**Explanation**: The sum of 2 and 7 is 9. Therefore, `index1 = 1`, `index2 = 2`. We return `[1, 2]`.

### Example 2
```yaml
Input: numbers = [2,3,4], target = 6
Output: [1,3]
```
**Explanation**: The sum of 2 and 4 is 6. Therefore, `index1 = 1`, `index2 = 3`. We return `[1, 3]`.

### Example 3
```yaml
Input: numbers = [-1,0], target = -1
Output: [1,2]
```
**Explanation**: The sum of -1 and 0 is -1. Therefore, `index1 = 1`, `index2 = 2`. We return `[1, 2]`.

## Constraints
- `2 <= numbers.length <= 3 * 10**4`
- `-1000 <= numbers[i] <= 1000`
- `numbers` is sorted in **non-decreasing order**.
- `-1000 <= target <= 1000`

## Solution
The first solution that came into my mind would be just using the brute force double nested loop. This would solve the problem of keeping the space complexity at `O(1)`. However, by doing this, we would increase the time complexity to `O(n**2)`.

We can use the fact the the input array is sorted in a non-decreasing order to use a two pointers method to figure out the solution. We can use two pointers of a left and right pointers that point at two end of the array. We can find the remainder between the target and the number at the right indexed. We then can compare if the remainder is smaller or bigger than the left indexed number. We can determine to whether shift the left index forward or the right index backward. The algorithm will run until the left index is equal to the right index. With this, we came up with a solution of `O(1)` for space complexity and `O(n)` for time complexity.
