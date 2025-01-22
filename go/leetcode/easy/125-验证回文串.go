package main

import (
	"fmt"
	"math"
	"strings"
)

func isPalindromeStr(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for {
		if left == right {
			break
		}
		if !((97 <= int32(s[left]) && int32(s[left]) <= 122) || (48 <= int32(s[left]) && int32(s[left]) <= 57)) {
			left++
			continue
		}
		if !((97 <= int32(s[right]) && int32(s[right]) <= 122) || (48 <= int32(s[right]) && int32(s[right]) <= 57)) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		if left == right || math.Abs(float64(right-left)) == 1 {
			break
		}
		left++
		right--
	}
	return true
}

func isPalindromeStr1(s string) bool {
	str := ""
	for _, c := range s {
		if (65 <= int32(c) && int32(c) <= 90) || (97 <= int32(c) && int32(c) <= 122) || 48 <= int32(c) && int32(c) <= 57 {
			str += strings.ToLower(string(c))
		}
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//s := "A man, a plan, a canal: Panama"
	//s := "aa"
	//s := "0P"
	//s := " "
	s := "race a car"
	fmt.Println(isPalindromeStr(s))
}
