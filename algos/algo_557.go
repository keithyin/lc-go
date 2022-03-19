package algos

func reverseWords(s string) string {
	sBytes := []byte(s)
	begin := 0
	end := 0
	for end < len(s) {
		if sBytes[end] == ' ' {
			reverseRange(sBytes, begin, end-1)
			end++
			begin = end
		}
		end++
	}
	reverseRange(sBytes, begin, end-1)
	return string(sBytes)
}

func reverseRange(s []byte, begin, end int) {
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}
