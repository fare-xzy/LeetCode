from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        temp = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[temp]:
                temp += 1
                nums[temp] = nums[i]
        return temp + 1


nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
s = Solution()
print(s.removeDuplicates(nums))
