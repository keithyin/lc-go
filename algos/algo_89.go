package algos

func grayCode(n int) []int {
	res := make([]int, 0)
	res = append(res, []int{0, 1}...)

	for i := 2; i <= n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			v := res[j]
			v = v | (1 << (i - 1))
			res = append(res, v)
		}
	}
	return res
}
