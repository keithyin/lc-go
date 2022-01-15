package algos

import "math"

func robCore(root *TreeNode, buffer map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if val, ok := buffer[root]; ok {
		return val
	}

	robThisVal := root.Val
	if root.Left != nil {
		robThisVal = robThisVal + robCore(root.Left.Left, buffer) + robCore(root.Left.Right, buffer)
	}
	if root.Right != nil {
		robThisVal = robThisVal + robCore(root.Right.Left, buffer) + robCore(root.Right.Right, buffer)
	}

	result := int(math.Max(float64(robThisVal), float64(robCore(root.Left, buffer)+robCore(root.Right, buffer))))
	buffer[root] = result

	return result

}
