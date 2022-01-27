package algos

func maximalRectangle(matrix [][]byte) int {
	leftConsectiveOnes := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		leftConsectiveOnes[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		count := 0
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				count++
			} else {
				count = 0
			}
			leftConsectiveOnes[i][j] = count
		}
	}

	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				minWidth := leftConsectiveOnes[i][j]
				h := 0

				for (i-h) >= 0 && leftConsectiveOnes[i-h][j] > 0 {
					if minWidth > leftConsectiveOnes[i-h][j] {
						minWidth = leftConsectiveOnes[i-h][j]
					}
					curArea := (h + 1) * minWidth
					if curArea > maxArea {
						maxArea = curArea
					}
					h++
				}
			}
		}
	}

	return maxArea
}
