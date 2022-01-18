package algos

func countSubstrings(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for span := 1; span < len(s); span++ {
		for end := span; end < len(s); end++ {
			begin := end - span
			if s[begin] == s[end] {
				if span > 1 {
					dp[begin][end] = dp[begin+1][end-1]
				} else {
					dp[begin][end] = 1
				}
			} else {
				dp[begin][end] = 0
			}
		}
	}
	count := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			count += dp[i][j]
		}
	}
	return count
}
