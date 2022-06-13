from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows == 0:
            return []
        lists = [[0] * numRows for i in range(numRows)]
        lists[0][0] = 1
        for i in range(numRows - 1, 0, -1):
            del lists[0][i]
        if numRows == 1:
            return lists
        lists[1][0] = 1
        lists[1][1] = 1
        for i in range(numRows - 1, 1, -1):
            del lists[1][i]
        if numRows == 2:
            return lists
        for i in range(2, numRows):
            lists[i][0] = 1
            for index in range(1, i):
                lists[i][index] = lists[i - 1][index - 1] + lists[i - 1][index]
            lists[i][i] = 1
            for j in range(numRows - 1, i, -1):
                del lists[i][j]
        return lists

    def generate1(self, numRows: int) -> List[List[int]]:
        lists = list()
        for i in range(numRows):
            temp = [1] * (i + 1)
            if i > 1:
                for j in range(1, i):
                    temp[j] = lists[i - 1][j - 1] + lists[i - 1][j]
            lists.append(temp)
        return lists



s = Solution()
print(s.generate1(5))
