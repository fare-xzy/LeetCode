package main

func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[k]
			nums[k] = temp
			k++
		}
	}
}

func main() {

}
