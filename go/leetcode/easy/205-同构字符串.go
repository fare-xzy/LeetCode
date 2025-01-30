package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	sm := make(map[uint8]uint8)
	tm := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		if sm[s[i]] == 0 {
			sm[s[i]] = t[i]
		} else if sm[s[i]] != t[i] {
			return false
		}
		if tm[t[i]] == 0 {
			tm[t[i]] = s[i]
		} else if tm[t[i]] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	//s := "egg"
	//t := "add"
	//s := "foo"
	//t := "bar"
	//s := "paper"
	//t := "title"
	s := "badc"
	t := "baba"
	fmt.Println(isIsomorphic(s, t))
}
