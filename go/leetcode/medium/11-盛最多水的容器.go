package main

func maxArea(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := 1; j < len(height); j++ {
			temp := 0
			if height[i] > height[j] {
				temp = height[j] * (j - i)
			} else {
				temp = height[i] * (j - i)
			}
			if area < temp {
				area = temp
			}
		}
	}
	return area
}

func maxArea1(height []int) int {
	area := 0
	start, end := 0, len(height)-1
	for start < end {
		temp := 0
		if height[start] > height[end] {
			temp = height[end] * (end - start)
			end--
		} else {
			temp = height[start] * (end - start)
			start++
		}
		if temp > area {
			area = temp
		}
	}
	return area
}

func main() {

}
