# 1488. Avoid Flood in The City
**Medium**

## Question
Your country has an infinite number of lakes. Initially, all the lakes are empty, but when it rains over the `nth` lake, the `nth` lake becomes full of water. If it rains over a lake that is **full of water**, there will be a **flood**. Your goal is to avoid floods in any lake.

Given an integer array `rains` where:

- `rains[i] > 0` means there will be rains over the `rains[i]` lake.
- `rains[i] == 0` means there are no rains this day and you can choose one lake this day and dry it.

Return an array `ans` where:

- `ans.length == rains.length`
- `ans[i] == -1` if `rains[i] > 0`.
- `ans[i]` is the lake you choose to dry in the ith day if `rains[i] == 0`.

If there are multiple valid answers return **any** of them. If it is impossible to avoid flood return **an empty array**.

Notice that if you chose to dry a full lake, it becomes empty, but if you chose to dry an empty lake, nothing changes.

## Example
### Example 1
```yaml
Input: rains = [1,2,3,4]
Output: [-1,-1,-1,-1]
Explanation: After the first day full lakes are [1]
After the second day full lakes are [1,2]
After the third day full lakes are [1,2,3]
After the fourth day full lakes are [1,2,3,4]
There's no day to dry any lake and there is no flood in any lake.
```

### Example 2
```yaml
Input: rains = [1,2,0,0,2,1]
Output: [-1,-1,2,1,-1,-1]
Explanation: After the first day full lakes are [1]
After the second day full lakes are [1,2]
After the third day, we dry lake 2. Full lakes are [1]
After the fourth day, we dry lake 1. There is no full lakes.
After the fifth day, full lakes are [2].
After the sixth day, full lakes are [1,2].
It is easy that this scenario is flood-free. [-1,-1,1,2,-1,-1] is another acceptable scenario.
```

### Example 3
```yaml
Input: rains = [1,2,0,1,2]
Output: []
Explanation: After the second day, full lakes are  [1,2]. We have to dry one lake in the third day.
After that, it will rain over lakes [1,2]. It's easy to prove that no matter which lake you choose to dry in the 3rd day, the other one will flood.
```

## Constraints
- `1 <= rain.length <= 10e5`
- `0 <= rains[i] <= 10e9`

## Solution
The first thing that came to my mind water reading this question and understand the process of how the answer got there was that I needed a way to keep track of the "lake" that contain water. In order to do so, I decided to use a hash map with the lake number as the key index that points to the last day that the last day that the lake got rained on. Also we cannot just blindly drain out one of the lake every time there is a dry day or we will get a wrong answer. For example, if we have `rains = [1,2,0,3,2]`, we cannot just dry out a random lake or having to look up ahead to see which lake will be full. Instead, I decided to keep an array that will keep track of the day that we have a no-rain day. As we move through the rains array, if we encounter a lake that would potentially be flooded and we have a dry day sometime before that, we can use the no-rain day to dry said lake. We will also need to keep an answer array to keep track of it.

The time complexity of this implementation would be `O(n**2)`, although if we use a binary tree or heap to keep track of the dry days, it would go down to `O(nlogn)`. The space complexity would be `O(n)`.
