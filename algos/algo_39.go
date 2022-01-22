package algos

func combinationSum(candidates []int, target int) [][]int {
	dp := make([][][]int, len(candidates))
	for i := 0; i < len(candidates); i++ {
		dp[i] = make([][]int, target+1)
		for j := 0; j < target+1; j++ {
			dp[i][j] = make([]int, 0)
		}
	}
	for i := 0; i < len(candidates); i++ {
		dp[i][0] = append(dp[i][0], 0)
	}

	for j := 0; j < target+1; j++ {
		if j%candidates[0] == 0 {
			dp[0][j] = append(dp[0][j], j/candidates[0])
		}
	}

	for i := 1; i < len(candidates); i++ {
		for j := 1; j < target+1; j++ {
			num := 0
			for (j - num*candidates[i]) >= 0 {
				if len(dp[i-1][j-num*candidates[i]]) > 0 {
					dp[i][j] = append(dp[i][j], num)
				}
				num++
			}
		}
	}

	results := make([][]int, 0)
	tracer := make([]int, 0)
	backtracer(candidates, dp, len(candidates)-1, target, &tracer, &results)
	return results

}

func backtracer(candidates []int, dp [][][]int, row, col int, tracer *[]int, results *[][]int) {
	if col == 0 {
		tmp := make([]int, len(*tracer))
		copy(tmp, *tracer)
		*results = append(*results, tmp)
		return
	}

	if col < 0 {
		return
	}

	choices := dp[row][col]
	for i := 0; i < len(choices); i++ {
		*tracer = append(*tracer, sliceInit(candidates[row], choices[i])...)
		backtracer(candidates, dp, row-1, col-candidates[row]*choices[i], tracer, results)
		*tracer = (*tracer)[:len(*tracer)-choices[i]]
	}
}

func sliceInit(val, copies int) (results []int) {
	results = make([]int, copies)
	for i := 0; i < copies; i++ {
		results[i] = val
	}
	return
}
