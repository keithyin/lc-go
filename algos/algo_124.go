package algos

func maxPathSum(root *TreeNode) int {
	tracer := make([]int, 0)
	singleSideMax, passRootMax := maxPathSumCore(root, &tracer)
	tracer = append(tracer, []int{singleSideMax, passRootMax}...)
	return SliceMaxInt(tracer)

}

// tracer 可以仅使用一个值来追踪最大值！！！！
func maxPathSumCore(root *TreeNode, tracer *[]int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftSingleSideMax, leftPassRootMax := maxPathSumCore(root.Left, tracer)
	rightSingleSideMax, rightPassRootMax := maxPathSumCore(root.Right, tracer)
	if root.Left != nil {
		*tracer = append(*tracer, []int{leftSingleSideMax, leftPassRootMax}...)

	}
	if root.Right != nil {
		*tracer = append(*tracer, []int{rightSingleSideMax, rightPassRootMax}...)

	}
	val := root.Val
	singleSideMax := SliceMaxInt([]int{leftSingleSideMax, rightSingleSideMax, 0}) + val
	passRootMax := val + SliceMaxInt([]int{leftSingleSideMax, 0}) + SliceMaxInt([]int{rightSingleSideMax, 0})
	return singleSideMax, passRootMax
}
