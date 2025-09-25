# 120. Triangle
**Medium**

## Question
Given a `triangle` array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index `i` on the current row, you may move to either index `i` or index `i + 1` on the next row.

## Example
### Example 1
```yaml
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
```
**Explanation** The triangle looks like:
```text
   2
  3 4
 6 5 7
4 1 8 3
```
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11.

### Example 2
```yaml
Input: triangle = [[-10]]
Output: -10
```

## Constraints
- `1 <= triangle.length <= 200`
- `triangle[0].length == 1`
- `triangle[i].length == triangle[i-1].length + 1`
- `-10**4 <= triangle[i][j] <= 10**4`

## Solution
When I looked at this problem, the first thing that came into my mind was implementing some form of priority queue. We can break every elements in the 2-D arrays into a nodes and each node with has up to two edges of `[row + 1][col]` and `[row + 1][col + 1]`. The graph would only be a directed graph with one direction of either top-down or bottom-up. After created the graph, we can implement a shortest path algorithm with Dijkstra's with priority queue to find the path with the minimum amount of value to each the bottom. This approach would give us a time complexity of `O((E + V)logV)` for `E` is the total connection edges between elements and `V` is the total elements in the array. The amount of `V` and `E` would be `n**2`, so the time complexity would be `O(n**2 log n)`. The space complexity of this would be `O(V)` or `O(n**2)`.

To implement Dijkstra's, we would need to first build up the graph by traversing through every elements and find its next edges. Then we can perform the algorithm on it. It feels that we may over complicated the problem trying to solve it this way.

I then looked at using Dynamic Programming to solve this problem. Since we can see that at each 'step' or row of the triangle, we only need to know the updated value from the previous row. We can approach this two way from using either top-down or bottom-up solution.

In top-down, we will just need to keep track of the value of the previous row and update the current row elements with the sum of the element with the minimum between the previous row `col_index` and `col_index - 1`. This will give us all the possible minimum value at the bottom row. This algorithm would give us a time complexity of `O(n**2)` and space complexity of `O(n**2)`. You can see the implementation for it at [top down](./Go/solution/ans_top_down.go).

With the bottom-up solution, we can approach it similar to the top down solution. Although the end result will only give us the most optimal path in the triangle, which is what we wanted. With bottom up, we only need to store value in one single array with the max amount of the base row of the triangle. To update the total element as we move up row, we will just need to replace the index of the current element with the sum of the `triangle[current_row][current_index]` and the minimum between `data[current_index]` and `data[current_index + 1]`. If we approach the data replacement in the row starting at `index 0`, we will update the total correctly and get the most optimal path at the end in `data[0]`. This solution would also give us a time complexity of `O(n**2)`, but now our space complexity is `O(n)`. You can see this algorithm implementation at [bottom up](./Go/solution/ans_bottom_up.go).

