package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	result := make([]int, len(mat)*len(mat[0]))
	x, y := 0, 0
	tmpX, tmpY := 0, 0
	direction := false
	for i := range result {
		result[i] = mat[x][y]
		if x == 0 && y < len(mat[0])-1 && !direction {
			tmpY = y + 1
			tmpX = 0
			x, y = tmpX, tmpY
			direction = true
			continue
		} else if y == 0 && x < len(mat)-1 && direction {
			tmpY = 0
			tmpX = x + 1
			x, y = tmpX, tmpY
			direction = false
			continue
		} else if y == len(mat[0])-1 && x < len(mat)-1 && !direction {
			tmpY = len(mat[0]) - 1
			tmpX = x + 1
			x, y = tmpX, tmpY
			direction = true
			continue
		} else if x == len(mat)-1 && y < len(mat[0])-1 && direction {
			tmpY = y + 1
			tmpX = len(mat) - 1
			x, y = tmpX, tmpY
			direction = false
			continue
		}

		if direction {
			x++
			y--
		} else {
			y++
			x--
		}
	}
	return result
}

func main() {
	//mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//mat := [][]int{{1, 2, 3}, {4, 5, 6}}
	mat := [][]int{{1, 2}, {3, 4}}
	fmt.Println(findDiagonalOrder(mat))
}
