package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	temp := 1
	tempNum := 0
	for _, num := range nums {
		if num <= 0 {
			continue
		} else {
			if temp != num {
				if tempNum > 0 && tempNum != num {
					return temp
				}
			} else {
				temp++
			}
			tempNum = num
		}
	}
	return temp
}

func main() {
	nums := []int{0, 2, 2, 1, 1}
	fmt.Println(firstMissingPositive(nums))
}
