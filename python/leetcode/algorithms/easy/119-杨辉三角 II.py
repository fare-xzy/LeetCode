from typing import List


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        lists = []
        for i in range(rowIndex + 1):
            temp = [1] * (i + 1)
            if i > 1:
                for j in range(1, i):
                    temp[j] = lists[i - 1][j - 1] + lists[i - 1][j]
            lists.append(temp)
        return lists[rowIndex]


s = Solution()
print(s.getRow(5))
