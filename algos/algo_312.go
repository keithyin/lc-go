package algos

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}

	for span := 2; span < len(nums); span++ {
		for i := 0; (i + span) < len(nums); i++ {
			end := i + span
			for j := 1; j < span; j++ {
				mid := i + j
				val := nums[i]*nums[end]*nums[mid] + dp[i][mid] + dp[mid][end]
				if val > dp[i][end] {
					dp[i][end] = val
				}

			}
		}
	}
	return dp[0][len(dp)-1]
}
