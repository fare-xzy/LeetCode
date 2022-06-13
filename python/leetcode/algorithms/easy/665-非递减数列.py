from typing import List


class Solution:
    def checkPossibility(self, nums: List[int]) -> bool:
        temp = nums[0]
        a = 0
        for i in range(1, len(nums)):
            if a == 0 and nums[i] < temp:
                if a != 0:
                    return False
                else:
                    if i != len(nums) - 1 and i != 2 and nums[i + 1] < nums[i - 1]:
                        return False
                    a = nums[i]
            temp = nums[i]
        if a == 0:
            return False
        return True

nums = [4,2,3]
s = Solution()
print(s.checkPossibility(nums))