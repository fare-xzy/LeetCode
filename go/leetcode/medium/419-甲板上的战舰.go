package main

func countBattleships(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				board[i][j] = '.'
				count++
				a := i + 1
				b := j
				for z := a; z < len(board); z++ {
					if board[z][b] == 'X' {
						board[z][b] = '.'
					} else {
						break
					}
				}
				a = i
				b = j + 1
				for x := b; x < len(board[0]); x++ {
					if board[a][x] == 'X' {
						board[a][x] = '.'
					} else {
						break
					}
				}
			}
		}
	}
	return count
}

func countBattleships1(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			a := i - 1
			b := j - 1
			if (i == 0 || board[a][j] == '.') && (j == 0 || board[i][b] == '.') && board[i][j] == 'X' {
				count++
			}
		}
	}
	return count
}

func main() {

}
