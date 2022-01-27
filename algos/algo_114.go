package algos

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

}

func flattenCore(root *TreeNode) *TreeNode {
	fakeRoot := new(TreeNode)
	cursor := fakeRoot
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		cursor.Right = top
		top.Left = nil
		cursor = cursor.Right
	}

	return fakeRoot.Right
}
