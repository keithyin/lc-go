package algos

func numberIsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	stack := make([]int, 0)

	for x > 0 {
		stack = append(stack, x%10)
		x = x / 10
	}
	base := 1
	reversedRes := 0
	for _, v := range stack {
		reversedRes += v * base
		base *= 10
	}
	return x == reversedRes
}
