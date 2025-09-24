# 128. Longest Consecutive Sequence
**Medium**

## Question
Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element. The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in `O(n)` time.

## Example
### Example 1
```yaml
Input: nums = [100,4,200,1,3,2]
Output: 4
```
**Explanation**: The longest consecutive elements sequence is `[1, 2, 3, 4]`. Therefore, its length is 4.

### Example 2
```yaml
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

### Example 3
```yaml
Input: nums = [1,0,1,2]
Output: 3
```

## Constraints:
- `0 <= nums.length <= 10**5`
- `-10**9 <= nums[i] <= 10**9`

## Solution
To be honest, I think that LeetCode did a terrible job with this problem. In this, even with the requirement of running in the time complexity of `O(n)`, you can still use the built-in sort function. That would mathematically took `O(nlogn)` time and still be able to get accepted.

If your goal was to just get this problem submitted, you can just take the easy way out and use the sort function. However, we are here to learn how to create an algorithm that would take an actual `O(n)` run time. In order to do that, I decided to keep track of every number inside the function in a hashset. Since we have the hashset, we can look up the number with a constant time. We can now use that to skip out number that is not a start of the sequence, by checing `current_number - 1` inside the hashset to check if it exist. If the number exist, we can skip that number inside the for loop. If the previous number does not exist, we can start counting from the current number and increase by 1 inside a while loop until the number does not exist anymore. Although the algorithm looks like we are doing a `while loop` inside a `for loop`, since we are skipping out on numbers that are not the beginning of the sequence. This would make the time complexity of our solution to be `O(n)` with a `O(n)` space complexity.

