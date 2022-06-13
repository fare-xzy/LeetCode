package main

func removeDuplicates(nums []int) int {
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[temp] {
			temp++
			nums[temp] = nums[i]
		}
	}
	return temp + 1
}

func main() {

}
