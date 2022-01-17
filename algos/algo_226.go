package algos

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmpLeft := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmpLeft)
	return root

}
