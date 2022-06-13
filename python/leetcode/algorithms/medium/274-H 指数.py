from typing import List


class Solution:
    def hIndex(self, citations: List[int]) -> int:
        citations.sort()
        temp = 0
        for i, citation in enumerate(citations):
            if citation == len(citations) - i:
                temp = len(citations) - i
                break
            if citation > len(citations) - i:
                temp = len(citations) - i
                break
        return temp


citations = [2, 3, 4, 1, 7, 6]
s = Solution()
print(s.hIndex(citations))
