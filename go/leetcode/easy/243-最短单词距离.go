package main

import (
	"fmt"
	"math"
)

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	a := 0
	b := 0
	distance := 0
	for i, s := range wordsDict {
		if s == word1 {
			a = i + 1
		}
		if s == word2 {
			b = i + 1
		}
		if a != 0 && b != 0 && (distance == 0 || int(math.Abs(float64(a-b))) < distance) {
			distance = int(math.Abs(float64(a - b)))
		}
	}
	return distance
}

func main() {
	//wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	//word1 := "coding"
	//word2 := "practice"
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "makes"
	word2 := "coding"
	fmt.Println(shortestDistance(wordsDict, word1, word2))
}
