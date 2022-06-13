package main

import "sort"

func minMoves(nums []int) int {
	sort.Ints(nums)
	move := 0
	for i := 1; i < len(nums); i++ {
		temp := nums[i] + move - nums[i-1]
		nums[i] += move
		move += temp
	}
	return move
}

func main() {

}
