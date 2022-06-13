package main

func getRow(rowIndex int) []int {
	lists := [][]int{}
	for i := 0; i < rowIndex+1; i++ {
		temp := make([]int, i+1)
		for j := 0; j < len(temp); j++ {
			if j == 0 || j == len(temp)-1 {
				temp[j] = 1
			} else {
				temp[j] = lists[i-1][j-1] + lists[i-1][j]
			}
		}
		lists = append(lists, temp)
	}
	return lists[rowIndex]
}

func main() {

}
