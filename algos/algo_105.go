package algos

func buildTreeFromPreOrderAndInOrder(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	val := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if val == inorder[i] {
			break
		}
	}

	inOrderLeftSide := inorder[0:i]
	inOrderRightSide := inorder[i+1:]

	preOrderLeftSide := preorder[1:(1 + len(inOrderLeftSide))]
	preOrderRightSide := preorder[1+len(inOrderLeftSide):]
	root.Val = val
	root.Left = buildTreeFromPreOrderAndInOrder(preOrderLeftSide, inOrderLeftSide)
	root.Right = buildTreeFromPreOrderAndInOrder(preOrderRightSide, inOrderRightSide)
	return root
}
