package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	sc, tc := [26]rune{}, [26]rune{}
	for _, si := range s {
		sc[si-'a']++
	}

	for _, ti := range t {
		tc[ti-'a']++
	}
	return sc == tc
}

func isAnagram1(s string, t string) bool {
	ss := strings.Split(s, "")
	st := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(st)
	return strings.Join(ss, "") == strings.Join(st, "")
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
