package main

import (
	"bytes"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strArray := strings.Split(str, "")
	var buffer bytes.Buffer
	for i := len(strArray) - 1; i >= 0; i-- {
		buffer.WriteString(strArray[i])
	}
	strNew := buffer.String()
	if str == strNew {
		return true
	} else {
		return false
	}
}
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i != j; i, j = i+1, j-1 {
		if i >= j {
			break
		}
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {

}
