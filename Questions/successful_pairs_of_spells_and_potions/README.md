# 2300. Successful Pairs of Spells and Potions
**Medium**

## Question
You are given two positive integer arrays `spells` and `potions`, of length `n` and `m` respectively, where `spells[i]` represents the strength of the `ith` spell and `potions[j]` represents the strength of the `jth` potion.

You are also given an integer `success`. A spell and potion pair is considered **successful** if the **product** of their strengths is **at least** `success`.

Return an integer array **pairs** of length **n** where `pairs[i]` is the number of **potions** that will form a successful pair with the `ith` spell.

## Example
### Example 1
```yaml
Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
Explanation:
- 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
- 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
- 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
Thus, [4,0,3] is returned.
```

### Example 2
```yaml
Input: spells = [3,1,2], potions = [8,5,8], success = 16
Output: [2,0,2]
Explanation:
- 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
- 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful.
- 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful.
Thus, [2,0,2] is returned.
```

## Constraints
- `n == spells.length`
- `m == potions.length`
- `1 <= n, m <= 10e5`
- `1 <= spells[i], potions[i] <= 10e5`
- `1 <= success <= 10e10`

## Solution
In order to check if the spell success, we will have to check by iterate through every `spells` and `potions` such as `spells[i] * potions[j] >= success`. However, to iterate through all of that would take a long time. So I decided to use binary search to perform the comparison to check for how many potions would be success for each `spells[i]`. In order to perform the binary search, we will first need to sort the `potions` array in either ascending or descending order. For every `spells[i]` we iterate through, we perform a binary search to check what is the smallest `potions[j]` that would be able to be success.

This algorithm would produce us with a time complexity of `O(nlog)` and a space complexity `O(1)`.
