package algos

func diameterOfBinaryTree(root *TreeNode) int {
	_, curTotLen := diameterOfBinaryTreeCore(root)
	return curTotLen
}

func diameterOfBinaryTreeCore(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftSideLen, leftTotLen := diameterOfBinaryTreeCore(root.Left)
	rightSideLen, rightTotLen := diameterOfBinaryTreeCore(root.Right)

	curTotLen := leftSideLen + rightSideLen + 1
	curSideLen := SliceMaxInt([]int{leftSideLen, rightSideLen}) + 1

	if leftTotLen > curTotLen {
		curTotLen = leftTotLen
	}
	if rightTotLen > curTotLen {
		curTotLen = rightTotLen
	}

	return curSideLen, curTotLen
}
