package algos

import "math"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for j := 0; j < len(grid[0]); j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = grid[0][j] + dp[0][j-1]
		}
	}
	for i := 0; i < len(grid); i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
