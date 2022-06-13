class Solution:
    def reverse(self, x: int) -> int:
        maxInt = 2 ** 31
        strX = str(abs(x))
        strList = list(strX)
        newStr = ''.join(strList[-1::-1])
        newInt = int(newStr)
        if newInt < -maxInt or newInt > maxInt - 1:
            return 0
        if x < 0:
            return -newInt
        else:
            return newInt


soloution = Solution()
print(soloution.reverse(1534236469))
