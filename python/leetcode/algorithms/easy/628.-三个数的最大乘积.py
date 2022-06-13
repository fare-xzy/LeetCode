from typing import List


class Solution:
    def maximumProduct(self, nums: List[int]) -> int:
        newNums = sorted(nums)
        length = len(newNums)
        if length < 5 and newNums[0] < 0 and newNums[1] < 0 and newNums[length - 1] > 0:
            return newNums[length - 1] * newNums[0] * newNums[1]
        elif length >= 5 and newNums[0] * newNums[1] > newNums[length - 2] * newNums[length - 3]:
            return newNums[length - 1] * newNums[0] * newNums[1]
        return newNums[length - 1] * newNums[length - 2] * newNums[length - 3]

    def maximumProduct1(self, nums: List[int]) -> int:
        newNums = sorted(nums)
        length = len(nums)
        num1 = newNums[length - 1] * newNums[0] * newNums[1]
        num2 = newNums[length - 1] * newNums[length - 2] * newNums[length - 3]
        return max(num1, num2)


s = Solution()
print(s.maximumProduct1(
    [-710, -107, -851, 657, -14, -859, 278, -182, -749, 718, -640, 127, -930, -462, 694, 969, 143, 309, 904, -651, 160,
     451, -159, -316, 844, -60, 611, -169, -73, 721, -902, 338, -20, -890, -819, -644, 107, 404, 150, -219, 459, -324,
     -385, -118, -307, 993, 202, -147, 62, -94, -976, -329, 689, 870, 532, -686, 371, -850, -186, 87, 878, 989, -822,
     -350, -948, -412, 161, -88, -509, 836, -207, -60, 771, 516, -287, -366, -512, 509, 904, -459, 683, -563, -766,
     -837, -333, 93, 893, 303, 908, 532, -206, 990, 280, 826, -13, 115, -732, 525, -939, -787]))
