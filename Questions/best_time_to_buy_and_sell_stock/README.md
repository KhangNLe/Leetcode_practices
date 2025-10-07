# 121. Best Time to Buy and Sell Stock
**Easy**

## Question
You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return `0`.

## Example
### Example 1
```yaml
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

### Example 2
```yaml
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

## Constraints
- `1 <= prices.length <= 10e5`
- `0 <= prices[i] <= 10e4`

## Solution
For this problem, our ultimate goal is to **buy low sell high**, in other word, we will track is with 2 prices points and a `current_max`. We will move one of the price point toward the end of the `prices` array one step at a time starting at index 1. The other point, we will start at `prices[0]`, this will be our buying price. For every step of the first pointer, we will check if the price is lower than our buying price. If it is lower, the will now point the buying price to this. If is higher, we will then calculate our profit and assign the higher amount of that and the `current_max` to our `current_max`. Once we reach the end of `prices`, we will have our most optimal buying and selling profit.

This algorithm will take `O(n)` for time complexity and `O(1)` for space complexity.

