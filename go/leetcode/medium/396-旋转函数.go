package main

import "fmt"

func maxRotateFunction(nums []int) int {
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			temp += nums[j] * j
		}

		if temp < 0 && max == 0 {
			max = temp
		}
		if max < temp {
			max = temp
		}
		ktemp := nums[0]
		for k := 0; k < len(nums)-1; k++ {
			nums[k] = nums[k+1]
		}
		nums[len(nums)-1] = ktemp
		temp = 0
	}
	return max
}

func maxRotateFunction1(nums []int) int {
	max, total := 0, 0
	for i, num := range nums {
		total += num
		max += num * i
	}

	com, l := max, len(nums)
	for i := range nums {
		com += total - l*nums[l-1-i]
		if com > max {
			max = com
		}
	}
	return max
}

func main() {
	nums := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(nums))
	fmt.Println(maxRotateFunction1(nums))
}
