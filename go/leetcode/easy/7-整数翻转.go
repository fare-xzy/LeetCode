package main

import (
	"bytes"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	str := strconv.Itoa(int(math.Abs(float64(x))))
	strArray := strings.Split(str, "")
	var buffer bytes.Buffer
	for i := len(strArray) - 1; i >= 0; i-- {
		buffer.WriteString(strArray[i])
	}
	strNew := buffer.String()
	intNew, _ := strconv.Atoi(strNew)
	if float64(intNew) < -math.Pow(2, 31) || float64(intNew) > math.Pow(2, 31)-1 {
		return 0
	}
	if x < 0 {
		return -intNew
	} else {
		return intNew
	}
}

func reverse1(x int) int {
	nums := make([]int, 0)
	leave := 0
	z := 0
	if x < 0 {
		z = 0 - x
	} else {
		z = x
	}

	for z != 0 {
		leave = z % 10
		z = z / 10
		nums = append(nums, leave)
	}

	returnNums := 0
	y := 1
	len := len(nums)
	for i := 0; i < len; i++ {
		y = y * 10
	}

	for _, v := range nums {
		y = y / 10
		returnNums = returnNums + v*y
	}

	if returnNums > (2147483648) || (returnNums < (2>>31)+1) {
		return 0
	}
	if x < 0 {
		returnNums = 0 - returnNums
	}
	return returnNums
}

func main() {

}
