package algos

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				pre := 0
				if i >= 2 {
					pre = dp[i-2]
				}
				dp[i] = pre + 2
			} else {
				if (i-dp[i-1]) > 0 && s[i-dp[i-1]-1] == '(' {
					pre := 0
					if (i - dp[i-1] - 1) > 0 {
						pre = dp[i-dp[i-1]-2]
					}
					dp[i] = 2 + pre + dp[i-1]
				}
			}
		}
	}
	if len(dp) == 0 {
		return 0
	}
	return SliceMaxInt(dp)
}

func ternaryConditionalOperator(flag bool, a, b int) int {
	if flag {
		return a
	}
	return b
}
