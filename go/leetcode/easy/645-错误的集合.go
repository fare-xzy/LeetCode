package main

import (
	"math"
	"sort"
)

func findErrorNums(nums []int) []int {

	newSums := []int{}
	sort.Ints(nums)
	tempMap := make(map[int]bool)
	for _, i := range nums {
		tempMap[i] = true
	}
	trueSum := (len(nums) + 1) * len(nums) / 2
	setSum := 0
	sum := 0
	for m := range tempMap {
		newSums = append(newSums, m)
	}
	for _, ns := range newSums {
		setSum += ns
	}
	for _, num := range nums {
		sum += num
	}
	return []int{sum - setSum, trueSum - setSum}
}

// 方法很巧妙
func findErrorNums1(nums []int) []int {
	n := len(nums)
	res := make([]int, 2)
	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] < 0 {
			res[0] = int(math.Abs(float64(nums[i])))
		} else {
			nums[index] *= -1
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res[1] = i + 1
		}
	}
	return res
}

func main() {

}
