package algos

func pathSum(root *TreeNode, targetSum int) int {
	prefixSum2CountMap := make(map[int]int)
	count := 0
	prefixSum2CountMap[0] = 1
	pathSumCore(root, 0, prefixSum2CountMap, &count, targetSum)
	return count
}

func pathSumCore(root *TreeNode, preCumSum int,
	prefixSum2CountMap map[int]int, count *int, target int) {

	if root != nil {
		curCumSum := preCumSum + root.Val

		if _, ok := prefixSum2CountMap[curCumSum-target]; ok {
			*count = *count + prefixSum2CountMap[curCumSum-target]
		}

		if _, ok := prefixSum2CountMap[curCumSum]; !ok {
			prefixSum2CountMap[curCumSum] = 0
		}
		prefixSum2CountMap[curCumSum]++

		pathSumCore(root.Left, curCumSum, prefixSum2CountMap, count, target)
		pathSumCore(root.Right, curCumSum, prefixSum2CountMap, count, target)
		prefixSum2CountMap[curCumSum]--
	}

}
