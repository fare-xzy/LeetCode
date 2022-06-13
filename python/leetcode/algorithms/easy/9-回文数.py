class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        xStr = str(x)
        xLen = len(xStr)
        if xLen == 1:
            return True
        if xLen % 2 == 0:
            aHalfx = int("".join(xStr[: xLen // 2 - 1: -1]))
            otherHafx = int("".join(xStr[0:xLen // 2:1]))
            return aHalfx == otherHafx
        else:
            aHalfx = int("".join(xStr[: xLen // 2: -1]))
            otherHafx = int("".join(xStr[0:xLen // 2:1]))
            return aHalfx == otherHafx
        # 为啥不直接判断反转后是否相等，上面的想法真脑残
    def isPalindrome1(self, x: int) -> bool:
        return str(x) == str(x)[::-1]

solution = Solution()
print(solution.isPalindrome(1000021))
