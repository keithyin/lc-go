package algos

func reverseString(s []byte) {

	left := 0
	right := len(s) - 1

	for left < right {
		tmp := s[right]
		s[right] = s[left]
		s[left] = tmp
		left++
		right--
	}
}
