from typing import List


class Solution:
    def minMoves(self, nums: List[int]) -> int:
        nums.sort()
        move = 0
        for i in range(1, len(nums)):
            temp = nums[i] + move - nums[i - 1]
            nums[i] += move
            move += temp
        return move

    def minMoves1(self, nums: List[int]) -> int:
        nums.sort()
        move = 0
        for num in nums:
            move += num - nums[0]
        return move


nums = [1, 5, 3, 2, 2, 7, 6, 4, 8, 9]
s = Solution()
print(s.minMoves(nums))