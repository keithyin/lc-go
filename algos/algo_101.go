package algos

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	leftSym := isSymmetricCore(left.Left, right.Right)
	rightSym := isSymmetricCore(left.Right, right.Left)

	return leftSym && rightSym
}
