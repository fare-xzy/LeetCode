package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	W, H := len(M), len(M[0])
	newM := [][]int{}
	for w := 0; w < W; w++ {
		newM = append(newM, make([]int, H))
		for h := 0; h < H; h++ {
			count := 0
			for sw := w - 1; sw <= w+1; sw++ {
				for sh := h - 1; sh <= h+1; sh++ {
					if sw < 0 || sh < 0 || sw >= W || sh >= H {
						continue
					}
					newM[w][h] += M[sw][sh]
					count++
				}
			}
			newM[w][h] /= count
		}
	}
	return newM
}

func main() {
	M := [][]int{{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}
	fmt.Print(imageSmoother(M))
}
