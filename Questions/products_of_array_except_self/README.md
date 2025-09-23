# Product of Array Except Self
**Medium**

## Question
Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

## Example
### Example 1
```yaml
Input: nums = [1,2,4,6]
Output: [48,24,12,8]
```
### Example 2
```yaml
Input: nums = [-1,0,1,2,3]
Output: [0,-6,0,0,0]
```

### Constraints:
- `2 <= nums.length <= 1,000`
- `-20 <= nums[i] <= 20`

## Solution
The first thing that came to my mind was to find the total product of the array. After getting the number, we can simply devide that product to each number inside the array to get the product of the arrays. That would provide us a run time of `O(n)` and the space compexity of `O(n)` for the output array and `O(1)` for any extra space.
That would solve the problem of creating an `O(n)` run time but we still used the division to solve our problem.

The solution that I came up was to do similar to getting the total product of the array, however, it will increase by step.

For example, if we have an array of
```yaml
arr = [a, b, c, d, e]
```
We can start the total product with 1 and multiple it to the current array. Then we can store it into the answer array. The steps would look something like

|       |   0   |   1   |   2   |   3   |   4   |
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|  arr  |   a   |   b   |   c   |   d   |   e   |
|new_arr|   1   |   a   |   ab  |   abc |   abcd|
|product|   1   |   a   |   ab  |   abc |   abcd|

Now, if we repeat the step the other way by multiplying the `new_arr` index with the product, then once again multiply the product with the current array index. We will have something like:
We will set product start at 1 again.

|       |   4   |   3   |   2   |   1   |   0   |
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|  arr  |   e   |   d   |   c   |   b   |   a   |
|old_arr|   abcd|   abc |   ab  |   a   |   1   |
|product|   1   |   e   |   ed  |   edc |   edcb|
|res_arr|   abcd|   abce|   abde|   acde|   edcb|

Now if er talk a look at the `res_arr`, we can see that every index inside the array is the total product of the array expect itself. Of course we do not need to create a new `res_arr` to store that answer, we can just add that into `new_arr`. By doing so, our program will have a time complexity of `O(n)` and space complexity of `O(n)`


