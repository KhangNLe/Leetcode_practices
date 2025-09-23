def hasDuplicate(arr: list[int]) -> bool:
    return len(arr) != len(set(arr))
