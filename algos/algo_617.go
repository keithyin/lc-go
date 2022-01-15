package algos

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var curNode *TreeNode

	if root1 == nil && root2 == nil {
		return curNode
	}

	curNode = new(TreeNode)
	if root1 != nil {
		curNode.Val += root1.Val
	}
	if root2 != nil {
		curNode.Val += root2.Val
	}
	if root1 == nil {
		curNode.Left = mergeTrees(nil, root2.Left)
		curNode.Right = mergeTrees(nil, root2.Right)
	} else if root2 == nil {
		curNode.Left = mergeTrees(root1.Left, nil)
		curNode.Right = mergeTrees(root1.Right, nil)
	} else {
		curNode.Left = mergeTrees(root1.Left, root2.Left)
		curNode.Right = mergeTrees(root1.Right, root2.Right)
	}

	return curNode

}
