package algos

func kthSmallestInBinarySearchTree(root *TreeNode, k int) int {
	result := -1
	kthSmallestInBinarySearchTreeCore(root, &k, &result)
	return result

}

func kthSmallestInBinarySearchTreeCore(root *TreeNode, remaining *int, result *int) bool {
	if root == nil {
		return false
	}
	if kthSmallestInBinarySearchTreeCore(root.Left, remaining, result) {
		return true
	}
	*remaining -= 1
	if *remaining == 0 {
		*result = root.Val
		return true
	}
	if kthSmallestInBinarySearchTreeCore(root.Right, remaining, result) {
		return true
	}
	return false

}
