# 125. Valid Palindrome
**Easy**

## Question
A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a **palindrome**, or `false` otherwise.

## Example
### Example 1
```yaml
Input: s = "A man, a plan, a canal: Panama"
Output: true
```
**Explanation**: `amanaplanacanalpanama` is a palindrome.

### Example 2
```yaml
Input: s = "race a car"
Output: false
```
**Explanation**: `raceacar` is not a palindrome.

### Example 3
```yaml
Input: s = " "
Output: true
```
**Explanation**: `s` is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

## Constraints:
- `1 <= s.length <= 2 * 10**5`
- `s` consists only of printable ASCII characters

## Solution
There are many way to solve this problem. I decided to solve with the two pointer method. Since both upper and lower case letter would be considered the same in the problem, I used the `lower()` function make everything in the lower-case. Since the question does not mention anything about keep the input string the same, we can just reassign the lower-case string back into the input string. This would help us keep the space complexity to `O(1)`.

Using the two pointer method to keep track of the beginning and end of the string, we can check to make sure that they are the same for the palindrome. We can check if each character, or byte, is within the alphabetical or numerical by checking its ASCII value. We then keep shifting the two pointer closer into each other until the left side is bigger than the right side. After running this algorithm, we would produce a time complexity of `O(n)`.
