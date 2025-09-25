def twoSum(nums: list[int], target: int) -> list[int]:
    num_pos = {n: i for i, n in enumerate(nums)}

    for i in range(len(nums)):
        remain = target - nums[i]
        if remain in num_pos and num_pos[remain] != i:
            return [i, num_pos[remain]]
    return []
