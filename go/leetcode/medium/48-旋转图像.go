package main

import "fmt"

/**
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rotate(matrix [][]int) {
	lens := len(matrix) - 1
	a := (lens + 1) / 2
	for i := 0; i < a; i++ {
		//start = (i, i)
		for j := 0; j < lens-2*i; j++ {
			ji := i + j
			temp := matrix[i][ji]
			matrix[i][i+j] = matrix[lens-ji][i]
			matrix[lens-ji][i] = matrix[lens-i][lens-ji]
			matrix[lens-i][lens-ji] = matrix[ji][lens-i]
			matrix[ji][lens-i] = temp
		}
	}
}
func main() {
	//matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Print(matrix)
}
