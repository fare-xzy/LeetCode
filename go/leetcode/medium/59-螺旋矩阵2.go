package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	i := 1
	x := 0
	y := 0
	xn := n
	fxn := 0
	yn := n
	fyn := 0
	t := true
	r := false
	b := false
	l := false
	for {
		result[x][y] = i
		if t {
			y++
			if y == yn-1 {
				t = false
				r = true
				fxn++
			}
		} else if r {
			x++
			if x == xn-1 {
				r = false
				b = true
				yn--
			}
		} else if b {
			y--
			if y == fyn {
				b = false
				l = true
				xn--
			}
		} else if l {
			x--
			if x == fxn {
				l = false
				t = true
				fyn++
			}
		}

		i++
		if i > n*n {
			break
		}
	}
	return result
}

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}
