package main

import (
	"fmt"
	"strings"
)

func titleToNumber(columnTitle string) int {
	split := strings.Split(columnTitle, "")
	result := 0
	for i, s := range split {
		num := convertExcel(s)
		mv := num * power(26, len(split)-i-1)
		result += mv
	}
	return result
}

func power(x, y int) int {
	result := 1
	for i := 1; i <= y; i++ {
		result *= x
	}
	return result
}

func convertExcel(remainder string) int {
	switch remainder {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	case "D":
		return 4
	case "E":
		return 5
	case "F":
		return 6
	case "G":
		return 7
	case "H":
		return 8
	case "I":
		return 9
	case "J":
		return 10
	case "K":
		return 11
	case "L":
		return 12
	case "M":
		return 13
	case "N":
		return 14
	case "O":
		return 15
	case "P":
		return 16
	case "Q":
		return 17
	case "R":
		return 18
	case "S":
		return 19
	case "T":
		return 20
	case "U":
		return 21
	case "V":
		return 22
	case "W":
		return 23
	case "X":
		return 24
	case "Y":
		return 25
	case "Z":
		return 26
	}
	return 0

}

func main() {
	//columnTitle := "A"
	//columnTitle := "AB"
	columnTitle := "ZY"
	fmt.Println(titleToNumber(columnTitle))
}
