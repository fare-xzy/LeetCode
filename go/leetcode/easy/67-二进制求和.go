package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	As := strings.Split(a, "")
	Bs := strings.Split(b, "")
	more := 0
	num := len(As)
	if len(As) < len(Bs) {
		num = len(Bs)
	}
	returnStr := ""
	for i := 0; i < num; i++ {
		if len(As)-1-i < 0 && len(Bs)-1-i < 0 {
			break
		}
		A := 0
		if len(As)-1-i >= 0 {
			A, _ = strconv.Atoi(As[len(As)-1-i])
		}
		B := 0
		if len(Bs)-1-i >= 0 {
			B, _ = strconv.Atoi(Bs[len(Bs)-1-i])
		}
		if (A + B + more) == 3 {
			returnStr = "1" + returnStr
			more = 1
		} else if (A + B + more) == 2 {
			returnStr = "0" + returnStr
			more = 1
		} else if (A + B + more) == 1 {
			returnStr = "1" + returnStr
			more = 0
		} else {
			returnStr = "0" + returnStr
			more = 0
		}
	}

	returnStr = strconv.Itoa(more) + returnStr
	returnStr = strings.TrimLeft(returnStr, "0")
	if len(returnStr) == 0 {
		returnStr = "0"
	}
	return returnStr
}

//func addBinary(a string, b string) string {
//	AA, _ := new(big.Int).SetString(a, 2)
//	BB, _ := new(big.Int).SetString(b, 2)
//	AA.Add(AA, BB)
//	return AA.Text(2)
//}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
