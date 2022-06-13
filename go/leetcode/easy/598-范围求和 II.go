package main

import "math"

func maxCount(m int, n int, ops [][]int) int {
	for i := range ops {
		m = int(math.Min(float64(m), float64(ops[i][0])))
		n = int(math.Min(float64(n), float64(ops[i][1])))
	}
	return m * n
}

func main() {
	m := 3
	n := 3
	operations := [][]int{{2, 2}, {3, 3}}
	print(maxCount(m, n, operations))
}
