package algos

func NumTrees(n int) int {
	return numTrees(n)
}

func numTrees(n int) int {
	gn := make([]int, n+1)
	fin := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		fin[i] = make([]int, n+1)
	}
	gn[0] = 1
	gn[1] = 1
	fin[1][1] = 1
	for i := 2; i <= n; i++ {
		tmp := 0
		for j := 1; j <= i; j++ {
			fin[j][i] = gn[j-1] * gn[i-j]
			tmp += fin[j][i]
		}
		gn[i] = tmp
	}

	return gn[n]

}
