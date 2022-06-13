from typing import List


class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        res = []
        for i, num in enumerate(nums):
            newNum = abs(num) - 1
            if nums[newNum] < 0:
                res.append(abs(num))
            else:
                nums[newNum] = -nums[newNum]
        return res


nums = [4, 3, 2, 7, 8, 2, 3, 1]
s = Solution()
print(s.findDuplicates(nums))
