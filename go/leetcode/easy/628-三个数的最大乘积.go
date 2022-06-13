package main

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	num1 := nums[length-1] * nums[0] * nums[1]
	num2 := nums[length-1] * nums[length-2] * nums[length-3]
	//return num1 > num2 ? num1 : num2  无三目运算符 这样不符合Golang“一种事情有且只有一种方法完成”的思想。
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {

}
