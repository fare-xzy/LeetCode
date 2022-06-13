package main

import "math"

func findDuplicates(nums []int) []int {
	var result []int
	for _, num := range nums {
		newNum := int(math.Abs(float64(num)))
		if nums[newNum-1] < 0 {
			result = append(result, newNum)
		} else {
			nums[newNum-1] = -nums[newNum-1]
		}
	}
	return result
}

func main() {

}
