package main

import "strings"

func romanToInt(s string) int {
	romaMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sArray := strings.Split(s, "")
	num := 0
	temp := 0
	for i := 0; i < len(sArray); i++ {
		nowNum, _ := romaMap[sArray[i]]
		if temp != 0 && temp < nowNum {
			num -= temp * 2
		}
		temp = nowNum
		num += temp
	}
	return num
}

func main() {

}
