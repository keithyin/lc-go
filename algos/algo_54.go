package algos

func spiralOrder(matrix [][]int) []int {
	results := make([]int, 0)
	numRow := len(matrix)
	numCol := len(matrix[0])
	for i := 0; i <= ((numRow-1)/2) && i <= ((numCol-1)/2); i++ {
		for col := i; col < (numCol - i); col++ {
			results = append(results, matrix[i][col])
		}
		for row := i + 1; row < (numRow - i); row++ {
			results = append(results, matrix[row][numCol-i-1])
		}
		for col := numCol - i - 2; col >= i && (numRow-i-1) != i; col-- {
			results = append(results, matrix[numRow-i-1][col])
		}
		for row := numRow - i - 2; row > i && i != (numCol-i-1); row-- {
			results = append(results, matrix[row][i])
		}
	}
	return results
}
