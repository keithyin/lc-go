package algos

func reverseInt(x int) int {
	isNeg := (x < 0)
	if isNeg {
		x = -x
	}
	stack := make([]int, 0)
	for x > 0 {
		mod := x % 10
		stack = append(stack, mod)
		x = x / 10
	}

	base := 1
	result := 0
	for i := len(stack) - 1; i >= 0; i-- {
		result += base * stack[i]
		base *= 10
	}
	if isNeg {
		result = -result
	}
	if result > (0x7fffffff) || result < (-0x80000000) {
		result = 0
	}
	return result

}
