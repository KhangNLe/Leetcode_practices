import pytest
from ans import twoSum


def test_solution():
    inputs = [
        {
            "nums": [2, 7, 11, 15],
            "target": 9,
            "res": [0, 1],
        },
        {
            "nums": [3, 2, 4],
            "target": 6,
            "res": [1, 2],
        },
        {
            "nums": [3, 3],
            "target": 6,
            "res": [0, 1],
        }
    ]

    for input in inputs:
        assert twoSum(input["nums"], input['target']) == input['res']
