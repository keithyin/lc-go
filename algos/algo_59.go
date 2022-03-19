package algos

func generateMatrix(n int) [][]int {
	results := make([][]int, n)
	for i := 0; i < n; i++ {
		results[i] = make([]int, n)
	}

	numRow := n
	numCol := n

	curVal := 1

	for i := 0; i <= ((numRow - 1) / 2); i++ {
		for col := i; col < (numCol - i); col++ {
			results[i][col] = curVal
			curVal++
		}
		for row := i + 1; row < (numRow - i); row++ {
			results[row][numCol-i-1] = curVal
			curVal++
		}
		for col := numCol - i - 2; col >= i; col-- {
			results[numRow-i-1][col] = curVal
			curVal++
		}
		for row := numRow - i - 2; row > i; row-- {
			results[row][i] = curVal
			curVal++
		}
	}
	return results
}
