package algos

func exist(board [][]byte, word string) bool {
	visit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existCore(board, visit, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func existCore(board [][]byte, visit [][]bool, word string, i, j, pos int) bool {
	if pos == len(word) {
		return true
	}

	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		if visit[i][j] {
			return false
		}
		if board[i][j] == word[pos] {
			visit[i][j] = true
			res := (existCore(board, visit, word, i, j-1, pos+1)) ||
				(existCore(board, visit, word, i, j+1, pos+1)) ||
				(existCore(board, visit, word, i-1, j, pos+1)) ||
				(existCore(board, visit, word, i+1, j, pos+1))
			visit[i][j] = false
			return res
		} else {
			return false
		}
	}
	return false

}
