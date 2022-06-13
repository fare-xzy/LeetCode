from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        old = len(nums)
        i = 0
        temp = 0
        while i < len(nums) and i + temp != len(nums) - 1:
            if nums[i] == 0:
                nums.remove(0)
                nums.append(0)
                temp += 1
                i -= 1
            i += 1

    def moveZeroes1(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        pos = 0
        for i in range(len(nums)):
            if (nums[i] != 0):
                nums[pos], nums[i] = nums[i], nums[pos]
                pos += 1


nums = [0, 1, 0, 3, 12]
s = Solution()
s.moveZeroes1(nums)
