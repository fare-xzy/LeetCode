from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        numsLen = len(nums)
        for i in range(k):
            temp = nums[numsLen - 1]
            for j in range(numsLen - 1, -1, -1):
                if j == 0:
                    nums[0] = temp
                else:
                    nums[j] = nums[j - 1]


nums = [1, 2, 3, 4, 5, 6, 7]
k = 3
s = Solution()
print(s.rotate(nums, k))
