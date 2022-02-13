package algos

func rotate(matrix [][]int) {
	rotateCore(matrix, len(matrix), 0)
}

func rotateCore(matrix [][]int, width int, beginPos int) {
	if width <= 1 {
		return
	}
	endPos := width - 1
	for i := 0; i < endPos; i++ {
		tmp := matrix[beginPos+i][endPos+beginPos]
		matrix[beginPos+i][beginPos+endPos] = matrix[beginPos+0][beginPos+i]
		tmp2 := matrix[endPos+beginPos][endPos-i+beginPos]
		matrix[endPos+beginPos][endPos-i+beginPos] = tmp
		tmp = matrix[endPos-i+beginPos][0+beginPos]
		matrix[endPos-i+beginPos][0+beginPos] = tmp2
		matrix[0+beginPos][i+beginPos] = tmp
	}

	rotateCore(matrix, width-2, beginPos+1)

}
