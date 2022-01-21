package algos

func IsValidBST(root *TreeNode) bool {
	_, _, res := isValidBSTCore(root)
	return res
}

func isValidBSTCore(root *TreeNode) (max, min int, valid bool) {
	if root == nil {
		return int(^uint(0) >> 1), -int(^uint(0)>>1) - 1, true
	}
	leftMax, leftMin, leftValid := isValidBSTCore(root.Left)
	if !leftValid {
		return 0, 0, false
	}
	rightMax, rightMin, rightValid := isValidBSTCore(root.Right)
	if !rightValid {
		return 0, 0, false
	}
	if root.Val > leftMax && root.Val < rightMin {
		if root.Left == nil {
			leftMin = root.Val
		}
		if root.Right == nil {
			rightMax = root.Val
		}
		return rightMax, leftMin, true
	}
	return 0, 0, false
}
