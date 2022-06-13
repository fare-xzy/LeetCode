package main

func findMaxConsecutiveOnes(nums []int) int {
	num := 0
	bigNum := 0
	for i := range nums {
		if nums[i] == 1 {
			num++
		} else {
			if bigNum < num {
				bigNum = num
			}
			num = 0
		}
	}
	if bigNum > num {
		return bigNum
	} else {
		return num
	}
}

func main() {

}
