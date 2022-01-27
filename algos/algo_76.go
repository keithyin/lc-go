package algos

func MinCoveredString(s, t string) string {
	return minCoveredString(s, t)
}

func minCoveredString(s string, t string) string {
	source := make(map[byte]int)
	target := make(map[byte]int)
	for _, b := range []byte(t) {
		if _, ok := target[b]; !ok {
			target[b] = 0
		}
		target[b] += 1
	}
	length := len(s) + 1
	result_begin := -1
	result_end := -1
	begin := 0
	for end, b := range []byte(s) {
		if _, ok := source[b]; !ok {
			source[b] = 0
		}
		source[b] += 1

		for satisfy(target, source) {
			if length > (end - begin + 1) {
				length = end - begin + 1
				result_begin = begin
				result_end = end
			}
			source[s[begin]] -= 1
			begin += 1

		}
	}
	if result_begin == -1 || result_begin == -1 {
		return ""
	}
	return s[result_begin:(result_end + 1)]

}

func satisfy(target map[byte]int, source map[byte]int) bool {
	result := true
	for k, v := range target {

		if count, ok := source[k]; ok {
			if count < v {
				result = false
				break
			}
		} else {
			result = false
			break
		}
	}
	return result
}
