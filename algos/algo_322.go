package algos

func coinChange(coins []int, amount int) int {
	dp := make([][]int, 0)

	for i := 0; i < len(coins); i++ {
		dp = append(dp, make([]int, amount+1))
	}

	for i := 0; i < len(coins); i++ {
		for j := 0; j < (amount + 1); j++ {
			if j == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 {
				if (j % coins[i]) == 0 {
					dp[i][j] = j / coins[0]
				} else {
					dp[i][j] = 999999
				}
				continue
			}

			dp[i][j] = 999999
		}
	}

	for i := 1; i < len(coins); i++ {
		for j := 1; j < (amount + 1); j++ {
			k := 0
			minCoinNum := 999999
			for ; (j - k*coins[i]) >= 0; k++ {
				curCoinNum := k + dp[i-1][j-k*coins[i]]
				if curCoinNum < minCoinNum {
					minCoinNum = curCoinNum
				}
			}
			dp[i][j] = minCoinNum
		}
	}

	return dp[len(coins)-1][amount]
}
