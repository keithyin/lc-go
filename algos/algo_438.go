package algos

func findAnagrams(s string, p string) []int {

	result := make([]int, 0)

	target := make(map[byte]int)
	for _, b := range []byte(p) {
		if _, ok := target[b]; !ok {
			target[b] = 0
		}
		target[b] += 1
	}

	source := make(map[byte]int)
	begin := 0

	for end, b := range []byte(s) {
		if _, ok := target[b]; !ok {
			source = make(map[byte]int)
			begin = end + 1
			end++
			continue
		}

		if _, ok := source[b]; !ok {
			source[b] = 0
		}
		source[b]++

		matchFlag := dictMatch(source, target)
		if matchFlag == 0 {
			result = append(result, begin)
			source[s[begin]]--
			begin++
		} else if matchFlag == 1 {
			for dictMatch(source, target) == 1 {
				source[s[begin]]--
				begin++
			}
			if dictMatch(source, target) == 0 {
				result = append(result, begin)
			}
		} // else, move left
	}
	return result

}

// 0: match, 1 all big or equal, -1 got less
func dictMatch(source, target map[byte]int) int {
	if len(source) != len(target) {
		return -1
	}
	equalNum := 0
	lessNum := 0
	biggerNum := 0
	for tK, tV := range target {
		if sV, ok := source[tK]; !ok {
			return -1
		} else {
			if sV == tV {
				equalNum++
			} else if sV < tV {
				lessNum++
			} else {
				biggerNum++
			}
		}
	}

	if equalNum == len(target) {
		return 0
	}
	if lessNum > 0 {
		return -1
	}
	return 1
}
