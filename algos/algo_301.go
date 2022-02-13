package algos

func removeInvalidParentheses(s string) []string {

	candidates := make(map[string]struct{})
	removeInvalidParenthesesCore(s, 0, "", 0, candidates)
	maxLen := -1
	for item := range candidates {
		if maxLen < len(item) {
			maxLen = len(item)
		}
	}
	result := make([]string, 0)
	for item := range candidates {
		if len(item) == maxLen {
			result = append(result, item)

		}
	}
	return result
}

func removeInvalidParenthesesCore(s string, curPos int, prefix string, score int, candidates map[string]struct{}) {
	if curPos >= len(s) {
		if score == 0 {
			candidates[prefix] = struct{}{}

		}
		return
	}

	if score == -1 {
		return
	}

	if s[curPos] == '(' {
		removeInvalidParenthesesCore(s, curPos+1, prefix+"(", score+1, candidates)
		removeInvalidParenthesesCore(s, curPos+1, prefix, score, candidates)

	} else if s[curPos] == ')' {
		removeInvalidParenthesesCore(s, curPos+1, prefix+")", score-1, candidates)
		removeInvalidParenthesesCore(s, curPos+1, prefix, score, candidates)

	} else {
		removeInvalidParenthesesCore(s, curPos+1, prefix+s[curPos:curPos+1], score, candidates)
	}

}
