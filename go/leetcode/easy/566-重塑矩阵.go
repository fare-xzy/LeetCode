package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}
	x, y := 0, 0
	for i := range result {
		for j := range result[i] {
			result[i][j] = mat[x][y]
			if y == n-1 {
				y = 0
				x++
			} else {
				y++
			}
		}
	}
	return result
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {7, 8, 9}}
	//mat := [][]int{{1, 2, 3}, {4, 5, 6}}
	//mat := [][]int{{1, 2}, {3, 4}}
	r, c := 2, 6
	fmt.Println(matrixReshape(mat, r, c))
}
