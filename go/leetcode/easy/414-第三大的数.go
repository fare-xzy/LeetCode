package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	fmt.Print(nums)
	b := nums[0]
	m := nums[0]
	s := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > b {
			s = m
			m = b
			b = nums[i]
			continue
		}
		if nums[i] == b {
			continue
		}

	}
	if m == s {
		return b
	}
	return s
}

func main() {

}
