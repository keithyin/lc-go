package algos

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode, buffer map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if val, ok := buffer[root]; ok {
		return val
	}

	robThisVal := root.Val
	if root.Left != nil {
		robThisVal = robThisVal + rob(root.Left.Left, buffer) + rob(root.Left.Right, buffer)
	}
	if root.Right != nil {
		robThisVal = robThisVal + rob(root.Right.Left, buffer) + rob(root.Right.Right, buffer)
	}

	result := int(math.Max(math.Max(float64(robThisVal), float64(rob(root.Left, buffer))), float64(rob(root.Right, buffer))))
	buffer[root] = result

	return result

}
