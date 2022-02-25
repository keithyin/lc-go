package algos

func myAtoi(s string) int {
	beginFlag := false
	stack := make([]byte, 0)
	isNeg := false
	for i := 0; i < len(s); i++ {
		if !beginFlag {
			if s[i] == ' ' {
				continue
			}
			if s[i] == '-' {
				isNeg = true
				beginFlag = true
				continue
			}
			if s[i] == '+' {
				beginFlag = true
				continue
			}
			if s[i] >= '1' && s[i] <= '9' {
				beginFlag = true
				stack = append(stack, s[i])
				continue
			}
			break

		}
		if beginFlag {
			if s[i] >= '0' && s[i] <= '9' {
				stack = append(stack, s[i])
			} else {
				break
			}
		}
	}
	base := 1
	result := 0
	for i := len(stack) - 1; i >= 0; i-- {
		val := int(stack[i] - '0')
		result += val * base
		if isNeg && (-result) < -0x80000000 {
			return -0x80000000
		}
		if !isNeg && (result > 0x7fffffff) {
			return 0x7fffffff
		}
		if base > 0x80000000 {
			if isNeg {
				return -0x80000000
			} else {
				return 0x7fffffff
			}
		}
		base *= 10
	}

	if isNeg {
		result = -result
	}

	return result

}
