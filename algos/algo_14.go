package algos

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	commonPrefix := ""
	minLength := 300
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	for i := 0; i < minLength; i++ {
		char := strs[0][i]
		allMatch := true
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != char {
				allMatch = false
				break
			}
		}
		if allMatch {
			commonPrefix += string(char)
		} else {
			break
		}
	}
	return commonPrefix
}
