package main

func findShortestSubArray(nums []int) int {
	start := make(map[int]int)
	end := make(map[int]int)
	count := make(map[int]int)
	for index, num := range nums {
		_, startOk := start[num]
		if startOk {
			end[num] = index
			_, countOk := count[num]
			if countOk {
				count[num] = count[num] + 1
			} else {
				count[num] = 1
			}
		} else {
			start[num] = index
		}
	}
	if len(count) == 0 {
		return 1
	}
	bigCount := 0
	interval := len(nums)
	for i := range count {
		if count[i] > bigCount {
			bigCount = count[i]
		}
	}
	for i := range count {
		if count[i] == bigCount {
			bigCount = count[i]
			temp := end[i] - start[i] + 1
			if temp < interval {
				interval = temp
			}
		}
	}
	return interval
}

func main() {

}
