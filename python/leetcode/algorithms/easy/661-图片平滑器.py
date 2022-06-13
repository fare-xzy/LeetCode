from typing import List


class Solution:
    def imageSmoother(self, M: List[List[int]]) -> List[List[int]]:
        W, H = len(M), len(M[0])
        newList = [[0] * H for _ in range(W)]
        for w in range(W):
            for h in range(H):
                count = 0
                for sw in range(w - 1, w + 2):
                    for sh in range(h - 1, h + 2):
                        if sw < 0 or sh < 0 or sw >= W or sh >= H:
                            continue
                        newList[w][h] += M[sw][sh]
                        count += 1
                newList[w][h] //= count
        return newList


s = Solution()
l = [[2, 3, 4], [5, 6, 7], [8, 9, 10], [11, 12, 13], [14, 15, 16]]
print(s.imageSmoother(l))
