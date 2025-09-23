import random
from solution import hasDuplicate


def test_contain_duplicate():
    for i in range(100):
        (arr, ok) = generateRandomArr(i)
        assert hasDuplicate(arr) == ok


def generateRandomArr(iter: int) -> (list[int], bool):
    arr = [0] * random.randint(2, 100)
    if iter ^ 1 < iter:
        generate_arr_no_duplicate(arr)
        return (arr, False)
    else:
        generate_arr_with_duplicate(arr)
        return (arr, True)


def generate_arr_no_duplicate(arr: list[int]):
    for i in range(len(arr)):
        arr[i] = i


def generate_arr_with_duplicate(arr: list[int]):
    size = len(arr)
    for i in range(size):
        arr[i] = random.randint(0, size-1)

    arr[-1] = arr[0]
