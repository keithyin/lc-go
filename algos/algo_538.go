package algos

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	rightNode := convertBST(root.Right)
	if rightNode != nil {
		root.Val += rightNode.Val
	}
	convertBST(root.Left)
	return root
}

func inorderTraverse(root *TreeNode) {
	visit := make(map[*TreeNode]bool)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	val := 0
	for len(stack) > 0 {
		_, ok := visit[stack[len(stack)-1].Right]
		if stack[len(stack)-1].Right != nil && !ok {

			stack = append(stack, stack[len(stack)-1].Right)

		} else {
			top := stack[len(stack)-1]
			visit[top] = true
			val += top.Val
			top.Val = val
			stack = stack[0 : len(stack)-1]
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		}
	}
}
