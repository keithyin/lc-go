package algos

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}

	prePrev := 1
	prev := 2

	for i := 3; i <= n; i++ {
		tmp := prev
		prev = prePrev + prev
		prePrev = tmp
	}

	return prev

}
