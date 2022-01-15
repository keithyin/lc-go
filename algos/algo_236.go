package algos

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := lowestCommonAncestorCore(root, p, q)
	return node
}

func lowestCommonAncestorCore(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	if root == p || root == q {
		return root, p == q
	}

	leftNode, leftFound := lowestCommonAncestorCore(root.Left, p, q)
	rightNode, rightFound := lowestCommonAncestorCore(root.Right, p, q)
	if leftFound {
		return leftNode, true
	}
	if rightFound {
		return rightNode, true
	}

	if leftNode != nil && rightNode != nil {
		return root, true
	}

	if leftNode != nil {
		return leftNode, false
	}
	if rightNode != nil {
		return rightNode, false
	}

	return nil, false

}
