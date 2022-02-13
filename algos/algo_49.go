package algos

import "sort"

func groupAnagrams(strs []string) [][]string {
	results := make(map[string][]string)
	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		newStr := string(strBytes)
		if _, ok := results[newStr]; !ok {
			results[newStr] = make([]string, 0)
		}
		results[newStr] = append(results[newStr], str)
	}

	finalResult := make([][]string, 0)

	for _, v := range results {
		tmp := make([]string, 0)
		for _, oriStr := range v {
			tmp = append(tmp, oriStr)
		}
		finalResult = append(finalResult, tmp)
	}
	return finalResult

}
