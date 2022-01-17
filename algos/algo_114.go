package algos

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

}

func flattenCore(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftNode := root.Left
	rightNode := root.Right

	root.Left = nil
	root.Right = leftNode

	flattenCore(leftNode)
	flattenCore(rightNode)

	return root
}
