from typing import List


class Solution:
    def findReplaceString(self, S: str, indexes: List[int], sources: List[str], targets: List[str]) -> str:
        newMap = {}
        for i, j in enumerate(indexes):
            newMap[j] = i
        newIndexes = sorted(indexes)
        temp = ""
        tempEnd = 0
        for j in newIndexes:
            if sources[newMap[j]] not in S[tempEnd:len(S)]:
                temp += S[tempEnd:len(S)]
                tempEnd = len(S)
                break
            si = S.index(sources[newMap[j]], j)
            temp += S[tempEnd:si]
            temp += targets[newMap[j]]
            tempEnd = si + len(sources[newMap[j]])
        if tempEnd < len(S):
            temp += S[tempEnd:len(S)]
        return temp


S = "jjievdtjfb"
indexes = [4, 6, 1]
sources = ["md", "tjgb", "jf"]
targets = ["foe", "oov", "e"]
s = Solution()
print(s.findReplaceString(S, indexes, sources, targets))
