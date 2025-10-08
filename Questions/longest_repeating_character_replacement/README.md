# 424. Longest Repeating Charater Replacement
**Medium**

## Question
You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

## Example
### Example 1
```yaml
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

### Example 2
```yaml
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
```

## Constraints
- `1 <= s.length <= 10e5`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`

## Solution
For this question, we can create two pointer of `l` and `r`, along with `mostFreq` to keep track to the most frequency character inside the string and `res` to keep track of the longest substring length after replacement. We then start moving `r` pointer forward starting at index `0`. At the beginning of each iteration, we will increase the character counter. I used an `int` array of sized `26` since `s` only consisted of 26 upper case letter, and calculated the index by using `ASCII value of s[i] - 65`. After updating the character frequency, afterward we can update the `mostFreq` if the current character frequency is bigger. We then check if the length of the current substring subtract the current highest frequency is smaller or equal to `k`, or `(r - l + 1) - mostFreq <= k`. If so, we can update our answer with `r - l + 1` if it is bigger. If not, we will then need to decrease the character frequency at `s[l]` and shift our `l` pointer forward by 1.

With this algorithm, we will have a time complexity of `O(s.length)` or `O(n)` and a space complexity of `O(26)` or `O(1)`
