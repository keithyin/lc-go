package algos

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorderTraversalCore(root, &res)
	return res
}

func inorderTraversalCore(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalCore(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalCore(root.Right, res)
}
