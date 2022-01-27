package algos

func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	cursor := head
	for cursor != nil {
		stack = append(stack, cursor.Val)
		cursor = cursor.Next
	}

	cursor = head
	for cursor != nil {
		if cursor.Val != stack[len(stack)-1] {
			return false
		}
		cursor = cursor.Next
		stack = stack[:len(stack)-1]
	}
	return true
}
