package algos

func lowestCommonAncestorInBST(root, p, q *TreeNode) *TreeNode {
	common := new(*TreeNode)
	lowestCommonAncestorInBSTCore(root, p, q, common)
	return *common
}

func lowestCommonAncestorInBSTCore(root, p, q *TreeNode, common **TreeNode) bool {
	if root == nil {
		return false
	}
	leftFound := lowestCommonAncestorInBSTCore(root.Left, p, q, common)
	rightFound := lowestCommonAncestorInBSTCore(root.Right, p, q, common)
	if (leftFound && rightFound) || ((leftFound || rightFound) && (root == p || root == q)) {
		*common = root
		return true
	}
	if root == p || root == q || leftFound || rightFound {
		return true
	}
	return false
}
