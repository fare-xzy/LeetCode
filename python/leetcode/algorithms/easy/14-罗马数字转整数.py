class Solution:

    def romanToInt(self, s: str) -> int:
        RomeMap = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        strList = list(s)
        numResult = 0
        numTemp = 0
        for i, str in enumerate(strList):
            num = RomeMap.get(str)
            numResult += num
            if numTemp != 0:
                if numTemp < num:
                    numResult -= numTemp*2
            numTemp = num
        return numResult


solution = Solution()
print(solution.romanToInt("MCMXCIV"))
