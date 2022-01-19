package algos

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) {
		curVal := matrix[i][j]
		if curVal == target {
			return true
		} else if curVal < target {
			j++
		} else if curVal > target {
			i--
		}
	}
	return false

}
