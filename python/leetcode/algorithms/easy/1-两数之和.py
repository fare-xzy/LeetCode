from typing import List


# 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
# 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if j > i:
                    if (nums[i] + nums[j]) == target:
                        return {i, j}

    def twoSum1(self, nums: List[int], target: int) -> List[int]:
        sumMap = {}
        for i in range(len(nums)):
            sumMap[nums[i]] = i
        for i in range(len(nums)):
            result = target - nums[i]
            if sumMap.get(result) is not None and sumMap.get(result) > 0 and (sumMap.get(result) != i):
                return {i, sumMap.get(result)}

    def twoSum2(self, nums: List[int], target: int) -> List[int]:
        sumMap = {}
        for i in range(len(nums)):
            result = target - nums[i]
            if sumMap.get(result) is not None and sumMap.get(result) >= 0 and (sumMap.get(result) != i):
                return {i, sumMap.get(result)}
            else:
                sumMap[nums[i]] = i

    def twoSum3(self, nums: List[int], target: int) -> List[int]:
        sumMap = {}
        for i, num in enumerate(nums):
            if target - num in sumMap:
                return {i, sumMap.get(target - num)}
            else:
                sumMap[nums[i]] = i


nums = [3, 3]
target = 6
solution = Solution()
returnNums = solution.twoSum(nums, target)
print(returnNums)
