# 166. Fraction to Recurring Decimal
**Medium**

## Question
Given two integers representing the `numerator` and `denominator` of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, **return any of them**.

It is **guaranteed** that the length of the answer string is less than `10**4` for all the given inputs.

## Example
### Example 1
```yaml
Input: numerator = 1, denominator = 2
Output: "0.5"
```

### Example 2
```yaml
Input: numerator = 2, denominator = 1
Output: "2"
```

### Example 3:
```
Input: numerator = 4, denominator = 333
Output: "0.(012)"
```

## Constraints:
- `-2**31 <= numerator, denominator <= 2**31 - 1`
- `denominator != 0`

## Solution
The first solution that I thought of was to use a hashset to keep track of each digit of the remainder that I "walked" through. That seems to works for most case. However, if the fractional part of the division start of with repeating the first few numbers but actually not an actual repeating fraction, the algorithm would fail. The example would be something like `113 / 1000 = "0.113"`; however, the algorthim would return a result of `"0.(1)"`.

To solve this problem, I decided to keep track of the fraction by using modulus to keep track of the remainder. I also switched out the hashset for a hashmap that keep track of the last index inside the string of the beginning of the first time the remainder first appear. This will help with determine where the parentheses start. We will then just have to keep mutiple the remainder to the power of 10 and get the new remainder, or moving down a step. We will just keep doing this until `remainder == 0` or we find the repeating parent.

The time complexity for this algorithm would be `O(n)` and the space complexity of the would also be `O(n)`, where n is the length of the fraction part.
