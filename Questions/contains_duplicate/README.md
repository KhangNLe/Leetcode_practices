# Contains Duplicate

**Easy**

## Question:
Given an integer array `nums`, return `true` if any value appears more than once in the array, otherwise return `false`.

## Example
### Example 1
```yaml
Input: nums = [1, 2, 3, 3]
Output: true
```

### Example 2
```yaml
Input: nums = [1, 2, 3, 4]
Output: false
```

## Solution
The first approach would be to figure out a brute force solution to find the duplicate with a nested for loop and figure out if a number appeared again inside the array. However, that would take a `O(n**2)` running time and `O(1)` for memory space.

To reduce the running time, we can try to create a hashmap or hashset to figure out if the length of the original array is the same as the length of the hashset/hashmap. The would take `O(n)` running time and `O(n)` memory space.
