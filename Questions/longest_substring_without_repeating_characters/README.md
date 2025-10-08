# 3. Longest Substring Without Repeating Characters
**Medium**

## Question
Given a string `s`, find the length of the **longest substring** without duplicate characters.

A **substring** is a contiguous sequence of characters within a string.

## Example
### Example 1
```yaml
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
```

### Example 2
```yaml
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

### Example 3
```yaml
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Constraints
- `0 <= s.length <= 5 * 10e4`
- `s` consists of English letters, digits, symbols, and spaces

## Solution
For this problem, since we are only working with English letters, digits, symbols, and space, we can just use the 127 ASCII character to represent our substring. Let's has an array `substring` with the size of `127`. Since there only thing we care about is if a character present inside the substring, we can use a `boolean` array. Now, to keep track of our longest substring, we will need two pointer and an available to keep track of the current longest. We will then set the `start` pointer at the beginning of at index `0`, and the second pointer will be our iteration from `1` to `s.length`. Every step, we will check if the ASCII value of `s[i]` is presented inside the `substring` array. If so, we can calculate the length of that substring using `i - start`. After than we can shift a `start` index one by one and until `s[start] == s[i]` and set the ASCII index of each `s[start]` inside the array back to `false`. We will then need to shift the `start` up one more point to represent the beginning of the new substring. Then we need to keep track of the longest length substring for our answer.

This algorithm will give us a time complexity of `O(n)` and a space complexity of `O(1)`.
