# Two Sum
**Easy**

## Question
Given an array of integers `nums` and an integer `target`, return the indices `i` and `j` such that `nums[i] + nums[j] == target` and `i != j`.

You may assume that every input has exactly one pair of indices `i` and `j` that satisfy the condition.

Return the answer with the smaller index first

## Example
### Example 1
```yaml
Input: nums = [3,4,5,6], target = 7
Output: [0,1]
```
Explanation: `nums[0] + nums[1] == 7`, so we return `[0,1]`

### Example 2
```yaml
Input: nums = [4,5,6], target = 10
Output: [0,2]
```
### Example 3:
```yaml
Input: nums = [5,5], target = 10
Output: [0,1]
```

### Constraints:
- `2 <= nums.length <= 1000`
- `-10,000,000 <= nums[i] <= 10,000,000`
- `-10,000,000 <= target <= 10,000,000`

## Solution
The first approach I would think of to do a double nested loop. We can get the remain of `target - currentPosition` and compare it to the rest of the array using a second for loop. This would take us `O(n**2)` for time complexity and `O(1)` for space complexity.

The better solution that I came up it was to use a hashmap to save the location of each number inside the array. Since we use the number as the key index, the value for the key index would be the last time the number appear inside the array. This would make program running at `O(n)` for time complexity and `O(n)` for space complexity.
