package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	high := len(matrix[0])
	width := len(matrix)
	allEle := width * high
	top := false
	left := false
	right := true
	bottom := false
	x := 0
	y := 0
	result := make([]int, allEle)
	index := 0
	for {

		if matrix[x][y] != 999 {
			result[index] = matrix[x][y]
			matrix[x][y] = 999
			index++
		}
		if index == allEle {
			break
		}
		if right && y < high {
			y++
			if y >= high || matrix[x][y] == 999 {
				y--
				right = false
				bottom = true
			}
		} else if bottom {
			x++
			if x >= width || matrix[x][y] == 999 {
				x--
				bottom = false
				left = true
			}
		} else if left {
			y--
			if y < 0 || matrix[x][y] == 999 {
				y++
				left = false
				top = true
			}
		} else if top {
			x--
			if x < 0 || matrix[x][y] == 999 {
				x++
				top = false
				right = true
			}
		}
	}
	return result
}

func main() {
	//operations := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	//operations := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}
	operations := [][]int{{3}, {2}}
	fmt.Print(spiralOrder(operations))
}
