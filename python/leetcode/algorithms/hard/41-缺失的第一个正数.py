from typing import List


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        nums = sorted(set(nums))
        temp = 1
        for i, num in enumerate(nums):
            if num <= 0:
                continue
            else:
                if temp != num:
                    return temp
                if i == len(nums) - 1:
                    return num + 1
                temp += 1
        return temp


nums = [0, 2, 2, 1, 1]
s = Solution()
print(s.firstMissingPositive(nums))
