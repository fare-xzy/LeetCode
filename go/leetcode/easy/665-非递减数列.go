package main

func checkPossibility(nums []int) bool {
	length := len(nums)
	if length == 1 || length == 2 {
		return true
	}
	temp := -1
	for i := 1; i < length; i++ {
		if nums[i] < nums[i-1] {
			if i >= 2 && i < length-1 && nums[i] < nums[i-2] && nums[i+1] < nums[i-1] {
				return false
			}
			if temp != -1 {
				return false
			} else {
				temp = i
			}
		}
	}
	return true
}
