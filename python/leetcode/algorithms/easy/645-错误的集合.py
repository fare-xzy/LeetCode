from typing import List


class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        newNums = sorted(nums)
        temp = 0
        miss = 0
        redundant = 0
        for i in newNums:
            if i - temp == 0:
                redundant = i
            elif i - temp > 1:
                miss = i - 1
            temp = i
        if miss == 0:
            if newNums[redundant - 1] == redundant:
                return [redundant, newNums[len(newNums) - 1] + 1]
            return [redundant, redundant + 1]
        return [redundant, miss]

    def findErrorNums1(self, nums: List[int]) -> List[int]:
        return [sum(nums) - sum(set(nums)), sum(range(1, len(nums) + 1)) - sum(set(nums))]


nums = [1, 5, 3, 2, 2, 7, 6, 4, 8, 9]
s = Solution()
print(s.findErrorNums1(nums))
