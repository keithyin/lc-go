package algos

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + SliceMaxInt([]int{leftDepth, rightDepth})
}
