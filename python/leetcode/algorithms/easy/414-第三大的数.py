from typing import List


class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        newNums = sorted(set(nums))
        if len(newNums) < 3:
            return newNums[-1]
        else:
            return newNums[-3]


s = Solution()
print(s.thirdMax([1, 1, 2]))
