package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	temp := 0
	for i, citation := range citations {
		if citation >= len(citations)-i {
			temp = len(citations) - i
			break
		}
	}
	return temp
}

func main() {

}
