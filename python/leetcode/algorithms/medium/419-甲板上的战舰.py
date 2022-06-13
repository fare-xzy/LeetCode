from typing import List


class Solution:
    def countBattleships(self, board: List[List[str]]) -> int:
        count = 0
        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == 'X':
                    board[i][j] = '.'
                    count += 1
                    a = i + 1
                    b = j
                    while a < len(board) and board[a][b] == 'X':
                        board[a][b] = '.'
                        a += 1
                    a = i
                    b = j + 1
                    while b < len(board[0]) and board[a][b] == 'X':
                        board[a][b] = '.'
                        b += 1
        return count

    def countBattleships1(self, board: List[List[str]]) -> int:
        count = 0
        for i in range(len(board)):
            for j in range(len(board[0])):
                a = i - 1
                b = j - 1
                if (i == 0 or board[a][j] == '.') and (j == 0 or board[i][b] == '.') and board[i][j] == 'X':
                    count += 1
        return count


s = Solution()
operations = [["X", ".", ".", "X"], [".", ".", ".", "X"], [".", ".", ".", "X"]]
# operations = [[".", "."], ["X", "X"]]
print(s.countBattleships1(operations))
