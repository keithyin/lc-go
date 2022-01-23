package algos

func CanPartition(nums []int) bool {
	return canPartition(nums)
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}

	half := sum / 2

	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, half+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	for j := 0; j < (half + 1); j++ {
		if j == nums[0] {
			dp[0][j] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j < (half + 1); j++ {
			if dp[i-1][j] || ((j-nums[i]) > 0 && dp[i-1][j-nums[i]]) {
				dp[i][j] = true
			}
		}
	}
	return dp[len(nums)-1][half]

}
