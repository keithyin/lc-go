package algos

/**

 */

func WordBreak(s string, wordDict []string) bool {
	result := wordBreak(s, wordDict)
	return result
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if !dp[i] && endWith(s, word, i) {
				if (i - len(word) + 1) > 0 {
					dp[i] = dp[i-len(word)]
				} else {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)-1]
}

func EndWith(source, target string, endPos int) bool {
	return endWith(source, target, endPos)
}

func endWith(source string, target string, endPos int) bool {
	if endPos >= (len(target) - 1) {
		beginPos := endPos - len(target) + 1
		for i, char := range []byte(target) {
			if source[beginPos+i] != char {
				return false
			}
		}
		return true

	}
	return false
}
