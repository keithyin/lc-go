package algos

func maximalSquare(matrix [][]byte) int {
	area := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				curArea := search(matrix, i, j)
				if area < curArea {
					area = curArea
				}
			}
		}
	}
	return area
}

func search(matrix [][]byte, i, j int) int {

	if len(matrix) == i || len(matrix[0]) == j {
		return 0
	}

	iNext := i + 1
	jNext := j + 1

	for iNext < len(matrix) && jNext < len(matrix[0]) && matrix[iNext][jNext] == '1' {
		allOne := true
		for tmpI := i; tmpI <= iNext; tmpI++ {
			if matrix[tmpI][jNext] != '1' {
				allOne = false
				break
			}
		}

		if !allOne {
			break
		}

		for tmpJ := j; tmpJ <= jNext; tmpJ++ {
			if matrix[iNext][tmpJ] != '1' {
				allOne = false
				break
			}
		}
		if !allOne {
			break
		}

		iNext++
		jNext++
	}

	iNext -= 1
	jNext -= 1
	return (iNext - i + 1) * (jNext - j + 1)
}
