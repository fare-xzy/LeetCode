from typing import List


class Solution:
    def findShortestSubArray(self, nums: List[int]) -> int:
        start, end, count = {}, {}, {}
        for index, num in enumerate(nums):
            if num not in start:
                start[num] = index
            else:
                end[num] = index
                count[num] = count.get(num, 0) + 1
        if len(count) == 0:
            return 1
        maxCount = max(count.values())
        span = len(nums)
        for i in count:
            if count[i] == maxCount:
                temp = end[i] - start[i] + 1
                if span > temp:
                    span = temp
        return span


nums = [1, 2, 2, 3, 1, 4, 2]
s = Solution()
print(s.findShortestSubArray(nums))
